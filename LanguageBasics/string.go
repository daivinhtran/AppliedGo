package main

import "fmt"

func main() {
	// var s string = "This string contains"
	// var b byte = s[2]
	// fmt.Println("s[2] =", b)         // 105
	// fmt.Println("s[2] =", string(b)) // i

	// fmt.Println(s[1:4])
	// immutable
	s1 := "I am immutable"
	s2 := s1
	s3 := s1[5:14]

	s1 = "I am a new, and " + s1 + ", too"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// muttable
	var m []byte // slice of bytes
}
