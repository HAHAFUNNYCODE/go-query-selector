package goqs

import (
	"fmt"
	"testing"
)

func TestElementQS(t *testing.T) {
	doc, _ := ParseHTMLFile("testing/test1.html")
	tests := []string{"[use]"}
	for _, t := range tests {
		elem, err := doc.QuerySelector(t)
		if err != nil {
			fmt.Println(err)
		} else {
			printElem(elem)
		}
	}

	// e, err := doc.QuerySelector("header")
	// if err != nil {
	// 	panic(err)
	// }
	// for e != nil {
	// 	printElem(e)
	// 	e = e.nextSibling
	// }
}

func TestElementQS_SO(t *testing.T) {
	doc, _ := ParseHTMLFile("testing/so.html")
	tests := []string{} //".post-text"}
	for _, t := range tests {
		elem, err := doc.QuerySelector(t)
		if err != nil {
			fmt.Println(err)
		} else {
			printElem(elem)
		}
	}

}
