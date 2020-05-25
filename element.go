package goqs

//HTMLElement holds the data for the html tags parsed from the file
type HTMLElement struct {
	Tag, Text   string
	Attributes  map[string]string
	Children    []*HTMLElement
	IsSingleTag bool
}
