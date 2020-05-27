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

//ClassSet is a set type with string elements.
//It has the following methods:
//Insert(string),
//Remove(string),
//Has(string)->bool,
//Slice()->[]string
type ClassSet map[string]bool

//MakeClassSet turns a slice of strings into a non-repeating
//set containing each element from the slice
func MakeClassSet(classes []string) ClassSet {
	set := make(ClassSet)
	for _, class := range classes {
		set.Insert(class)
	}
	return set
}

//Insert puts the string class into the set.
//If the class is already there, there will be no effect.
func (set ClassSet) Insert(class string) {
	set[class] = true
}

//Remove takes a class out of a set.
//If the class is not there, there will be no effect.
func (set ClassSet) Remove(class string) {
	set[class] = false
}

//Has returns if a class exists inside of a set
func (set ClassSet) Has(class string) bool {
	return set[class]
}

//Slice returns a slice containing each element
//that resides in the set
func (set ClassSet) Slice() []string {
	slice := make([]string, 0)
	for class, exists := range set {
		if exists {
			slice = append(slice, class)
		}
	}
	return slice
}
