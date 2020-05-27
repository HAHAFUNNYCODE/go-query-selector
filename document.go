package goqs

//Document stores the contents of an html page in order, and all of the top level elements
type Document struct {
	page, topElements []*HTMLElement
}

//QuerySelector ...
//Do later
func (doc *Document) QuerySelector(pattern string) (*HTMLElement, error) {
	topDummyElement := HTMLElement{Children: doc.topElements}
	return topDummyElement.QuerySelector(pattern)
}
