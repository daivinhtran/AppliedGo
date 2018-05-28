package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {

	// }

	// // while loop
	// for i < 10 {

	// }

	//
	// for {
	// 	for { // Innerloop
	// 		if enough {
	// 			break // Innerloop
	// 		}
	// 	}
	// }

	// Used label to break the Outerloop
	// OuterLoop:
	// for {
	// 	for { // Innerloop
	// 		if enough {
	// 			break OuterLoop
	// 		}
	// 	}
	// }

	s := "Georgia Tech"
	for index, element := range s {
		// element is a copy of s[index]
		fmt.Println(index, ":", string(element), ", ")
	}

	for _, element := range s {
		fmt.Println(string(element))
	}

	for a := range s { // a in index in this case
		fmt.Println(a)
	}

	// modify content of slice using index
	intArr := []int{1, 2, 3}
	for index, _ := range intArr {
		intArr[index] = 5
	}
	fmt.Println(intArr)
}
