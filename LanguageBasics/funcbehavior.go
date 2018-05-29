package main

import (
	"fmt"
	"os"
)

func faculty(n int) int {
	if n == 0 { // break condition
		return 1
	}
	return n * faculty(n-1)
}

// Ok..
func f() error {
	f, err := os.Open("funcbehavior.go")
	if err != nil {
		return err
	}
	fmt.Println(f.Name())
	if f.Name() == "funcbehavior.go" {
		f.Close()
		return nil
	}
	//...
	f.Close()
	return nil
}

// Better: deferred func
func f() error {
	f, err := os.Open("funcbehavior.go")
	if err != nil {
		return err
	}
	defer f.Close() // is called whenever f() exits
	fmt.Println(f.Name())
	if f.Name() == "funcbehavior.go" {
		return nil
	}
	// ..
	return nil
}

func main() {

}
