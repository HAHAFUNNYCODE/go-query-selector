package goqs

import (
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
	ClassList   []string
	Attributes  map[string]string
	Children    []*HTMLElement
	IsSingleTag bool
	TokenType   html.TokenType

	pageIndex, depth int
	textIndex        int
	atomicTag        atom.Atom
}

//Text formats the stored text objects into a proper string
func (elem *HTMLElement) Text() string {
	str := ""
	for i, childIndex := 0, 0; i+childIndex < len(elem.text)+len(elem.Children); {
		if i >= len(elem.text) {
			if formattingTags.has(elem.Children[childIndex].atomicTag) || textTags.has(elem.Children[childIndex].atomicTag) || sectioningTags.has(elem.Children[childIndex].atomicTag) {
				str += elem.Children[childIndex].Text()
			}
			childIndex++
		} else {
			if i+childIndex == elem.text[i].index {
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

//QuerySelector ...
func (elem *HTMLElement) QuerySelector(pattern string) {

}

func (elem *HTMLElement) getFirstMatchingChild() {

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
	atom.P:  true,
	atom.H1: true,
	atom.H2: true,
	atom.H3: true,
	atom.H4: true,
	atom.H5: true,
	atom.H6: true,
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
}
