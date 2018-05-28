package main

import (
	"fmt"
	"os"
)

var (
	g       int
	h       string
	i       = 1234
	j, k, l bool
)

func main() {
	var a int
	var b = 10
	var c = 10
	var d, e, f bool

	var (
		i       = 1234
		j, k, l bool
	)

	m := 1
	n, o := 2, 3
	a = 11
	e, f = f, e
	a, p := 100, 200

	fmt.Println("Hello, World!")
	fmt.Println(a)

	// blank identifier
	// always check error returns; they're provided for a reason
	if _, err := os.Stat(path); os.IsNotExist(err) { // Good!
		fmt.Printf("%s does not exist\n", path)
	}

	// Bad!
	fi, _ := os.Stat(path)
	if fi.IsDir() { // will crash if path does not exist
		fmt.Printf("%s is a directory\n", path)
	}

}
