package main

import "fmt"

func f1(s string) bool { return len(s) > 0 }
func f2(s string) bool { return len(s) < 4 }

var funcVar func(string) bool

func funcAsParam(s string, f func(string) bool) bool {
	return f(s + "abcd")
}

func newClosure() func() {
	var a int
	return func() {
		fmt.Println(a)
		a++
	}
}

func caveat() {
	s := "abcd"
	var funcs []func()
	for _, c := range s {
		c := c
		funcs = append(funcs, func() {
			fmt.Println(string(c))
		})
	}
	for _, f := range funcs {
		f()
	}
}

func main() {
	// funcVar = f1
	// fmt.Println(funcVar("abcd"))
	// funcVar = f2
	// fmt.Println(funcVar("abcd"))

	funcVar = func(s string) bool {
		return len(s) > 4
	}
	fmt.Println(funcVar("abcd"))
	// fmt.Println(funcAsParam("", funcVar))

	// var result string = func() string {
	// 	return "abcd"
	// }
	// c := newClosure()
	// c()
	// c()
	// c()
	//caveat()

	// Functions are first-class objects in Go fvar := f()

	// An anonymous function literal can be invoked directly
	res := func() int {
		return 1
	}()
	fmt.Println(res)
}
