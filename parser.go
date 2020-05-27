package goqs

import (
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

//ParsingError occurs when an parsing fails or an html.ErrorToken appears
type ParsingError struct{}

func (ParsingError) Error() string {
	return "HTML was parsed unsuccessfully."
}

//ParseHTML takes an html string and parses the html into elements
func ParseHTML(htmlStr string) (Document, error) {
	htmlReader := strings.NewReader(htmlStr)
	tokenizer := html.NewTokenizer(htmlReader)
	stack := HTMLStack{}
	doc := Document{}

	index := 0

	for {
		tokenizer.Next()
		if tokenizer.Err() == io.EOF {
			break
		}
		token := tokenizer.Token()

		element := HTMLElement{Raw: token,
			TokenType:  token.Type,
			Attributes: make(map[string]string),
			atomicTag:  token.DataAtom,
			Children:   make([]*HTMLElement, 0),
			text:       make([]textOrder, 0)}

		switch token.Type {
		case html.StartTagToken:
			fallthrough

		case html.SelfClosingTagToken:
			element.Tag = token.Data
			for _, attr := range token.Attr {
				if attr.Key == "class" {
					element.ClassList = MakeClassSet(strings.Split(attr.Val, " "))
				} else if attr.Key == "id" {
					element.ID = attr.Val
				}
				element.Attributes[attr.Key] = attr.Val
			}
			fallthrough

		case html.DoctypeToken:
			element.depth = stack.Size()
			element.pageIndex = index
			index++
			doc.page = append(doc.page, &element)

			if stack.Size() != 0 {
				top, _ := stack.Top()
				if l := len(top.Children); l > 0 {
					top.Children[l-1].nextSibling = &element
				}
				top.Children = append(top.Children, &element)
				if formattingTags.has(token.DataAtom) || textTags.has(token.DataAtom) || sectioningTags.has(token.DataAtom) {
					top.textIndex++
				}
			}
			if token.Type == html.StartTagToken && !selfClosingTags.has(token.DataAtom) {
				stack.Push(&element)
			}

		case html.EndTagToken:
			endElem, err := stack.Pop()
			if err != nil {
				return Document{}, err
			}
			if stack.Size() == 0 {
				doc.topElements = append(doc.topElements, endElem)
			}

		case html.CommentToken:
		case html.TextToken:
			if stack.Size() != 0 {
				top, _ := stack.Top()
				if token.Data != "\n" {
					for strings.HasPrefix(token.Data, "\n") {
						token.Data = strings.Replace(token.Data, "\n", "", 1)
					}
					top.text = append(top.text, textOrder{text: token.Data, index: top.textIndex})
					top.textIndex++
				}
			}

		case html.ErrorToken:
			return Document{}, ParsingError{}
		}
	}
	return doc, nil
}

//ParseHTMLFile takes a filepath and parses the html inside
func ParseHTMLFile(filePath string) (Document, error) {
	htmlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Document{}, err
	}

	htmlStr := string(htmlBytes)
	if err != nil {
		return Document{}, err
	}

	return ParseHTML(htmlStr)
}
