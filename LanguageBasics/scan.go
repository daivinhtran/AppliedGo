package main

import "fmt"

func main() {
	var s string
	var n int

	fmt.Println("Please enter a word and a number.")
	cnt, err := fmt.Scanf("a %s %d", &s, &n)
	fmt.Println(cnt, s, n, err)
	if err != nil {
		fmt.Println("Scan failed:", err.Error())
	}
}
