package goqs

import "strings"

type queryPattern struct {
	tag, tagInv     []string
	class, classInv []string
	id, idInv       []string
	valid           bool
}

func makeQueryPattern(base string) []queryPattern {
	patSlice := make([]queryPattern, 0)
	baseSplit := strings.Split(base, ",")
	for _, v := range baseSplit {
		patSlice = append(patSlice, parsePattern(v))
	}
	return patSlice
}

func parsePattern(base string) queryPattern {
	return queryPattern{}
}
