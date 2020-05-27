package goqs

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type textOrder struct {
	text  string
	index int
}

//HTMLElement holds the data for the html tags parsed from the file
type HTMLElement struct {
	Raw         html.Token
	Tag, ID     string
	text        []textOrder
	ClassList   ClassSet
	Attributes  map[string]string
	Children    []*HTMLElement
	IsSingleTag bool
	TokenType   html.TokenType

	pageIndex, depth int
	textIndex        int
	atomicTag        atom.Atom
	nextSibling      *HTMLElement
}

//Text formats the stored text objects into a proper string
func (elem *HTMLElement) Text() string {
	// fmt.Println(elem.String())
	str := ""
	for i, childIndex := 0, 0; i+childIndex < len(elem.text)+len(elem.Children); {
		// fmt.Println(elem.Tag, elem.ClassList.Slice(), elem.depth, len(elem.text)+len(elem.Children), elem.pageIndex)
		if i >= len(elem.text) {
			if formattingTags.has(elem.Children[childIndex].atomicTag) || textTags.has(elem.Children[childIndex].atomicTag) || sectioningTags.has(elem.Children[childIndex].atomicTag) {
				str += elem.Children[childIndex].Text()
			}
			childIndex++
		} else {
			if i+childIndex >= elem.text[i].index {
				str += strings.ReplaceAll(elem.text[i].text, "\n", "")
				i++
			} else {
				if formattingTags.has(elem.Children[childIndex].atomicTag) || textTags.has(elem.Children[childIndex].atomicTag) || sectioningTags.has(elem.Children[childIndex].atomicTag) {
					str += elem.Children[childIndex].Text()
				}
				childIndex++
			}
		}
	}

	numTab, numNL := strings.Count(str, "\t"), strings.Count(str, "\n")
	_, _ = numTab, numNL
	str = strings.Replace(str, "\t", "", numTab)
	// str = strings.Replace(str, "\n", "", numNL)

	numDSpace := strings.Count(str, "  ")
	for numDSpace > 0 {
		str = strings.Replace(str, "  ", " ", numDSpace)
		numDSpace = strings.Count(str, "  ")
	}
	return str
}

func (elem *HTMLElement) String() string {
	str := ""
	str += elem.Raw.String() + "\n"
	str += "PageIndex:" + strconv.Itoa(elem.pageIndex) + "\n"
	str += "Tag:" + elem.Tag + "\n"
	str += "Text:\n"
	for _, t := range elem.text {
		str += t.text + " " + strconv.Itoa(t.index) + "\n"
	}
	str += "ID:" + elem.ID + "\n"
	str += "Classes:" + strings.Join(elem.ClassList.Slice(), ",") + "\n"
	str += "Attributes:\n"
	for k, v := range elem.Attributes {
		str += k + ":" + v + "\n"
	}
	str += "Children:\n"
	for _, c := range elem.Children {
		str += c.Raw.String() + "\n"
	}
	return str
}

//QuerySelector ...
func (elem *HTMLElement) QuerySelector(pattern string) (*HTMLElement, error) {
	matches, err := makeQueryPatterns(pattern)
	if err != nil {
		return nil, err
	}

	var matchedElems []*HTMLElement

	for _, p := range matches {
		if child := elem.getChild(p); child != nil {
			matchedElems = append(matchedElems, child)
		}
	}

	var finalElem *HTMLElement
	if len(matchedElems) > 0 {
		finalElem = matchedElems[0]
		for _, e := range matchedElems {
			if e.pageIndex < finalElem.pageIndex {
				finalElem = e
			}
		}
		return finalElem, nil
	}

	return nil, nil
}

func (elem *HTMLElement) getChild(qp queryPattern) *HTMLElement {
	switch qp.combinator {
	case descendant:
		fallthrough
	case directDescendant:
		return elem.getFirstMatchingDescendant(qp, elem.depth+1, qp.combinator)
	case sibling:
		fallthrough
	case adjacentSibling:
		qpCombined := qp.combined
		qp.combined = nil
		topSibling := elem.getFirstMatchingDescendant(qp, elem.depth, qp.combinator)
		if topSibling == nil {
			return nil
		}
		return topSibling.getFirstMatchingSibling(*qpCombined, qp.combinator)
	}
	return nil
}

func (elem *HTMLElement) getFirstMatchingDescendant(qp queryPattern, depth int, combinator int) *HTMLElement {
	for _, child := range elem.Children {
		if child.checkMatch(qp) {
			if qp.combined == nil {
				return child
			}
			if match := child.getChild(*qp.combined); match != nil {
				return match
			}
		}
		if combinator != directDescendant {
			if match := child.getFirstMatchingDescendant(qp, depth, combinator); match != nil {
				return match
			}
		}
	}
	return nil
}

func (elem *HTMLElement) getFirstMatchingSibling(qp queryPattern, combinator int) *HTMLElement {
	fmt.Println(verboseCombinators[combinator])
	if elem.nextSibling == nil {
		return nil
	}

	elem = elem.nextSibling
	if elem.checkMatch(qp) {
		if qp.combined == nil {
			return elem
		}
		if match := elem.getChild(*qp.combined); match != nil {
			return match
		}
	}

	if combinator == adjacentSibling {
		return nil
	}

	elem = elem.nextSibling
	for elem != nil {
		if elem.checkMatch(qp) {
			if qp.combined == nil {
				return elem
			}
			return elem.getChild(*qp.combined)
		}
	}
	return nil
}

func (elem *HTMLElement) checkMatch(qp queryPattern) bool {
	if (qp.tag != "" && qp.tag != elem.Tag) || (qp.id != "" && qp.id != elem.ID) {
		return false
	}

	for _, class := range qp.class {
		if !elem.ClassList.Has(class) {
			return false
		}
	}

	for attr, val := range qp.attr {
		if val == "" {
			if elem.Attributes[attr] == "" {
				return false
			}
		} else {
			if elem.Attributes[attr] != val {
				return false
			}
		}
	}

	return true
}

var selfClosingTags atomSet = atomSet{
	atom.Area:   true,
	atom.Base:   true,
	atom.Br:     true,
	atom.Embed:  true,
	atom.Hr:     true,
	atom.Iframe: true,
	atom.Img:    true,
	atom.Input:  true,
	atom.Link:   true,
	atom.Meta:   true,
	atom.Param:  true,
	atom.Source: true,
	atom.Track:  true,
}

var formattingTags atomSet = atomSet{
	atom.B:      true,
	atom.Strong: true,
	atom.I:      true,
	atom.Em:     true,
	atom.Mark:   true,
	atom.Small:  true,
	atom.Del:    true,
	atom.Ins:    true,
	atom.Sub:    true,
	atom.Sup:    true,
	atom.Br:     true,
	atom.Hr:     true,
}
var textTags atomSet = atomSet{
	atom.P:    true,
	atom.H1:   true,
	atom.H2:   true,
	atom.H3:   true,
	atom.H4:   true,
	atom.H5:   true,
	atom.H6:   true,
	atom.A:    true,
	atom.Pre:  true,
	atom.Code: true,
}

var sectioningTags atomSet = atomSet{
	atom.Section: true,
	atom.Div:     true,
	atom.Header:  true,
	atom.Aside:   true,
	atom.Footer:  true,
	atom.Nav:     true,
	atom.Article: true,
	atom.Main:    true,
	atom.Span:    true,
}
