package main

import "fmt"

func main() {
	fmt.Println("Hello,", "gopher") // Hello,<space>gopher

	fmt.Println("cat" + "bird") // no space between
	a := 10
	b := true
	c := "ten"
	// %t for boolean, %s for string, %d for integer
	fmt.Printf("It is %t that %d is spelled '%s'\n", b, a, c)
}
