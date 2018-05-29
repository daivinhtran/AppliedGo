package main

import "fmt"

func newClosures() (func(), func() int) {
	a := 0
	f1 := func() {
		a = 5
	}
	f2 := func() int {
		return a * 7
	}
	return f1, f2
}

func trace(name string) func() {
	fmt.Println("Entering " + name)
	return func() {
		fmt.Println("Leaving " + name)
	}
}

func f() {
	defer trace("f")()
	fmt.Println("Doing something")
}

func main() {
	f1, f2 := newClosures()
	f1()      // sets "a" to 5
	n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call?
	fmt.Println(n)
	f()
}
