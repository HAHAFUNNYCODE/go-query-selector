package goqs

import (
	"fmt"
	"strings"
	"testing"
)

func printTimes(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(s)
	}
}

// func digElements(elem *HTMLElement, tabs int) {
// 	printTimes("\t", tabs)
// 	fmt.Println(elem.Raw)
// 	printTimes("\t", tabs)
// 	fmt.Println(elem.pageIndex, elem.depth)
// 	printTimes("\t", tabs)
// 	fmt.Println(elem.TokenType, "Text:", elem.Text())
// 	printTimes("\t", tabs)
// 	fmt.Println("Tag:", elem.Tag, "ID:", elem.ID, "Classes:", strings.Join(elem.ClassList, ","))
// 	printTimes("\t", tabs)
// 	fmt.Println("Attributes:")
// 	for k, v := range elem.Attributes {
// 		printTimes("\t", tabs)
// 		fmt.Println(k, ":", v)
// 	}
// 	fmt.Println()

// 	for _, e := range elem.Children {
// 		digElements(e, tabs+1)
// 	}
// }

func printElem(elem *HTMLElement) {
	printTimes("\t", elem.depth)
	fmt.Println(elem.Raw)
	printTimes("\t", elem.depth)
	fmt.Println(elem.pageIndex, elem.depth)
	printTimes("\t", elem.depth)
	fmt.Println(elem.TokenType, "Text:", elem.Text())
	printTimes("\t", elem.depth)
	fmt.Println("Tag:", elem.Tag, "ID:", elem.ID, "Classes:", strings.Join(elem.ClassList, ","))
	printTimes("\t", elem.depth)
	fmt.Println("Attributes:")
	for k, v := range elem.Attributes {
		printTimes("\t", elem.depth)
		fmt.Println(k, ":", v)
	}
	fmt.Println()
}

func TestParseHTML(t *testing.T) {
	doc, _ := ParseHTMLFile("testing/test1.html")

	// for _, e := range doc.topElements {
	// 	digElements(e, 0)
	// 	_ = e
	// }

	for _, e := range doc.page {
		printElem(e)
	}
}
