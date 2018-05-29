package main

import "fmt"

func verify(i int) (int, error) {
	if i < 0 || i > 10 {
		return 0, fmt.Errorf("verify: %d is outside the allowed range", i)
	}
	return i, nil
}

func propagate(i int) error {
	if _, err := verify(i); err != nil {
		return fmt.Errorf("propagate: %s", err)
	}
	return nil
}

func unexpectedError(p *int) {
	if p == nil {
		panic("p must not be nil")
	}
}

func main() {
	// n, err := verify(12)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//fmt.Println(propagate(24))

	defer func() {
		res := recover()
		if res != nil {
			fmt.Println("Recovered:", res)
		}
	}()
	unexpectedError(nil)
}
