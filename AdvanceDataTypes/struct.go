package main

import (
	"strings"
)

type List struct {
	Value int
	next  *List
}

func main() {
	var l := List {
		Value: 1
	}
}

// visibility
package mylib
func Export() {}
func internal() {}

// same for struct
type s struct { // not comparable
	Export string
	internal []byte // can only be access within the same package
}

func uppercase(p Planet) Planet { // p = copy of the struct
	p.Name = strings.ToUpper(p.Name)
	return p
}

func rename(p *Planet) { // point to the same struct
	p.Name = strings.ToUpper(p.Name)
}