package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// default value for flags
	verbose := flag.Bool("verbose", false, "Print details")
	count := flag.Int("count", 1, "Number of iterations")

	var max int
	flag.IntVar(&max, "max", 10, "Maximum count")
	// flagStringVar

	// define flag var
	// then parse
	// count := flag.Int(...)
	// flag.Int(...) returns a pointer to a variable
	// flag.StringVar(...) receives an existing pointer to a variable
	// flag.Parse() : read flag variables

	flag.Parse()

	if *verbose == true {
		fmt.Printf("%#v: %d\n", count, *count)
	} else {
		fmt.Printf("%d\n", *count)
	}

	if len(os.Args) > max {
		fmt.Println("Found", len(os.Args), "parameters.",
			"Only", max, "are allowed.")
	}
}
