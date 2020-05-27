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
	descendant = iota + 1
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
type PatternError struct {
	msg string
}

func (e PatternError) Error() string {
	return e.msg
}

func parsePattern(base string) (queryPattern, error) {
	heirarchy := strings.Split(base, " ")
	patterns := make([]queryPattern, 0)
	for _, pat := range heirarchy {
		if combinators[pat] > 0 {
			if len(patterns) == 0 {
				return queryPattern{}, PatternError{"Pattern syntax is unrecognized"}
			}
			patterns[len(patterns)-1].combinator = combinators[pat]
			continue
		}

		var curPattern queryPattern = queryPattern{combinator: descendant, attr: make(map[string]string)}
		if len(patterns) > 0 && patterns[len(patterns)-1].combinator == descendant {
			patterns[len(patterns)-1].combinator = descendant
		}
		var curString *string = &curPattern.tag
		skipTo := -1

		for i, char := range pat {
			if i < skipTo {
				continue
			}

			if combinators[string(char)] > 0 {
				curPattern.combinator = combinators[string(char)]
				patterns = append(patterns, curPattern)
				curPattern = queryPattern{combinator: descendant}
				continue
			}

			switch char {
			case '.':
				curPattern.class = append(curPattern.class, "")
				curString = &curPattern.class[len(curPattern.class)-1]
			case '#':
				if curPattern.id != "" {
					return queryPattern{}, PatternError{"Multiple ids found in pattern"}
				}
				curString = &curPattern.id
			case '[':
				curString = nil
				var key, val string
				var err error
				key, val, skipTo, err = parseAttribute(pat, i+1)
				if err != nil {
					return queryPattern{}, err
				}
				curPattern.attr[key] = val
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

func parseAttribute(base string, index int) (string, string, int, error) {
	attr, value := "", ""
	curString := &attr
	startedAttr, finishedAttr := false, false
	startedVal, finishedVal := false, false
	openQuote, closedQuote := false, false
	for index < len(base) {
		char := rune(base[index])
		switch char {
		case '=':
			curString = &value

		case ']':
			if attr != "" && openQuote == closedQuote {
				return attr, value, index + 1, nil
			}
			if openQuote != closedQuote {
				return attr, value, index + 1, PatternError{"Opened quote was not closed"}
			}
			return attr, value, index + 1, PatternError{"Empty attribute in pattern"}

		case ' ':
			if startedAttr {
				finishedAttr = true
			}
			if startedVal && !closedQuote {
				finishedVal = true
			}

		case '\'':
			fallthrough
		case '"':
			if openQuote {
				closedQuote = true
				finishedVal = true
			} else {
				if curString != &value {
					return "", "", index, PatternError{"Quote found before value started"}
				}
				openQuote = true
			}

		default:
			if curString == &attr && finishedAttr {
				return "", "", index, PatternError{"Multiple spaces in attribute name"}
			}
			if curString == &value && finishedVal {
				return "", "", index, PatternError{"Value extends past end of quotes or of value"}
			}
			if closedQuote {
				return "", "", index, PatternError{"Value after closed quote"}
			}

			if curString == &attr {
				startedAttr = true
			} else if curString == &value {
				startedVal = true
			}

			*curString += string(char)
		}

		index++
	}
	return "", "", index, PatternError{"Attribute was not closed"}
}
