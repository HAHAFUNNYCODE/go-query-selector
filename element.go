package goqs

//HTMLElement holds the data for the html tags parsed from the file
type HTMLElement struct {
	Tag, Text             string
	ClassList, Attributes map[string]string
	Children              []*HTMLElement
	IsSingleTag           bool
}

func (elem *HTMLElement) QuerySelector(pattern string) {

}

func (elem *HTMLElement) getFirstMatchingChild() {

}
