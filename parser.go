package goqs

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//Document stores the contents of an html page in order, and all of the top level elements
type Document struct {
	page, topElements []HTMLElement
}

type ParsingError struct{}

func (ParsingError) Error() string {
	return "HTML was parsed unsuccessfully."
}

//ParseHTML takes an html string and parses the html into elements
func ParseHTML(htmlStr string) (Document, error) {
	htmlReader := strings.NewReader(htmlStr)
	tokenizer := html.NewTokenizer(htmlReader)
	for tokenizer.Err() != io.EOF {
		fmt.Println(tokenizer.Token().Data)
		fmt.Println(tokenizer.Next())
	}
	return Document{}, nil
}

//ParseHTMLFile takes a filepath and parses the html inside
func ParseHTMLFile(filePath string) (Document, error) {
	htmlFile, err := os.Open(filePath)
	if err != nil {
		return Document{}, err
	}
	htmlBytes := make([]byte, 0)
	_, err = htmlFile.Read(htmlBytes)
	if err != nil {
		return Document{}, err
	}
	htmlStr := string(htmlBytes)
	return ParseHTML(htmlStr)
}
