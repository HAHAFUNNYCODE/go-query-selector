package goqs

//Document stores the contents of an html page in order, and all of the top level elements
type Document struct {
	page, topElements []*HTMLElement
}

//QuerySelector returns a pointer to the first element
//satisfying the criteria in the input string found in the html document.
//The function uses CSS selectors like in JavaScript. For reference, follow
//the link to https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors
func (doc *Document) QuerySelector(pattern string) (*HTMLElement, error) {
	topDummyElement := HTMLElement{Children: doc.topElements}
	return topDummyElement.QuerySelector(pattern)
}
