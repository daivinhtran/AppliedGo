package main

import "fmt"

func f(n int, m int, s string) {
	fmt.Println("A function with param:", n, m, s)
}

func f2(s ...string) { // arbitary number of params
	for _, str := range s {
		fmt.Print(str + " ")
	}
	fmt.Println()
}

// return values
func f3() int {
	return 99
}

func f4() (int, string, error) {
	return 0, "ok", nil
}

func f5() (n int, s string) {
	// n and s can be used without declaring
}

func main() {
	f(37, 29, "ok")
	f2("Clear", "is", "better", "than", "clever")
}
