package main

import "fmt"

func main() {
	a, b := 1, 1
	if a == b {
		fmt.Printf("They are the same!")
	} else {
		fmt.Printf("They are different")
	}

	// this practice is commonly used in error handling
	if err := f(x); err != nil {
		fmt.Println(err.Error())
	}
}
