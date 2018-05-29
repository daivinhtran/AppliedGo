package main

import (
	"fmt"
	"os"
	"strings"
)

func acronym(s string) (acr string) {

	// TODO: Your code here
	var words = strings.Split(s, " ")
	//arcs = []rune
	return words[0]
}

func main() {
	s := "Âaz Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
