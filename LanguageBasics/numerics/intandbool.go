package main

import "fmt"

func main() {
	var f float64
	f = 3.14152
	var n int64
	n = int64(f)
	fmt.Println(f, n)

	// silent overflow
	var a uint8
	a = 250
	a += 10 // 4, not 260
}
