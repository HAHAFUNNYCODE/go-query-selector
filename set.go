package goqs

import "golang.org/x/net/html/atom"

type atomSet map[atom.Atom]bool

func (set atomSet) insert(a atom.Atom) {
	set[a] = true
}

func (set atomSet) remove(a atom.Atom) {
	set[a] = false
}

func (set atomSet) has(a atom.Atom) bool {
	return set[a]
}

//ClassSet is a set type with string elements
type ClassSet map[string]bool

func MakeClassSet(classes []string) ClassSet {
	set := make(ClassSet)
	for _, class := range classes {
		set.Insert(class)
	}
	return set
}

func (set ClassSet) Insert(class string) {
	set[class] = true
}

func (set ClassSet) Remove(class string) {
	set[class] = false
}

func (set ClassSet) Has(class string) bool {
	return set[class]
}

func (set ClassSet) Slice() []string {
	slice := make([]string, 0)
	for class, exists := range set {
		if exists {
			slice = append(slice, class)
		}
	}
	return slice
}
