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
