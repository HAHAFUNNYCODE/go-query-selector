package goqs

import (
	"strings"
)

type queryPattern struct {
	tag, tagInv     string
	class, classInv []string
	id              string
	idInv           []string
	attr, attrInv   map[string]string
	combined        *queryPattern
	valid           bool
	combinator      int
}

func makeQueryPatterns(base string) ([]queryPattern, error) {
	patSlice := make([]queryPattern, 0)
	baseSplit := strings.Split(base, ",")
	for _, v := range baseSplit {
		pattern, err := parsePattern(v)
		if err != nil {
			return nil, err
		}
		patSlice = append(patSlice, pattern)
	}
	return patSlice, nil
}

const (
	single = iota + 1
	descendant
	directDescendant
	sibling
	adjacentSibling
)

var combinators map[string]int = map[string]int{
	">": directDescendant,
	"~": sibling,
	"+": adjacentSibling,
}

//PatternError is thrown when a parsed querySelector pattern's syntax is unrecognized
type PatternError struct{}

func (e PatternError) Error() string {
	return "Pattern syntax is unrecognized"
}

func parsePattern(base string) (queryPattern, error) {
	heirarchy := strings.Split(base, " ")
	patterns := make([]queryPattern, 0)
	for _, pat := range heirarchy {
		if combinators[pat] > 0 {
			if len(patterns) == 0 {
				return queryPattern{}, PatternError{}
			}
			patterns[len(patterns)-1].combinator = combinators[pat]
			continue
		}

		var curPattern queryPattern = queryPattern{combinator: single}
		if len(patterns) > 0 && patterns[len(patterns)-1].combinator == single {
			patterns[len(patterns)-1].combinator = descendant
		}
		var curString *string = &curPattern.tag
		for _, char := range pat {
			if combinators[string(char)] > 0 {
				curPattern.combinator = combinators[string(char)]
				patterns = append(patterns, curPattern)
				curPattern = queryPattern{combinator: single}
				continue
			}

			switch char {
			case '.':
				curPattern.class = append(curPattern.class, "")
				curString = &curPattern.class[len(curPattern.class)-1]
			case '#':
				curString = &curPattern.id
				// case '[':
				// 	curPattern.attr = append(curPattern.attr)
			default:
				*curString = *curString + string(char)
			}
		}
		patterns = append(patterns, curPattern)
	}
	for i := 0; i < len(patterns)-1; i++ {
		patterns[i].combined = &patterns[i+1]
	}
	if len(patterns) > 0 {
		return patterns[0], nil
	}
	return queryPattern{}, nil
}
