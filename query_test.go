package goqs

import (
	"fmt"
	"testing"
)

var verboseCombinators map[int]string = map[int]string{
	single:           "Single",
	descendant:       "Descandant",
	directDescendant: "Direct Descendant",
	sibling:          "Sibling",
	adjacentSibling:  "Adjacent Sibling",
}

func TestParsePatterns(t *testing.T) {
	patterns := []string{"main + div", "main+div"}
	for _, p := range patterns {
		pat, err := parsePattern(p)
		if err != nil {
			panic(err)
		}
		for {
			fmt.Println("Pattern:", p)
			fmt.Println("Tag:", pat.tag, "Inv:", pat.tagInv)
			fmt.Println("Class:", pat.class, "Inv:", pat.classInv)
			fmt.Println("ID:", pat.id, "Inv:", pat.idInv)
			fmt.Println("Attr:", pat.attr, "Inv:", pat.attrInv)
			fmt.Println("Combinator:", verboseCombinators[pat.combinator])
			fmt.Println()

			if pat.combined == nil {
				break
			}

			pat = *pat.combined
		}
	}
}