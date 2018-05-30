package main

import (
	"fmt"
)

func appendOne(s *[]int) {
	*s = append(*s, 1)
}

func changeSlice1(s []int) {
	s[0] = 7
}

func main() {
	// s1 := make([]int, 4, 8)
	// // s1 = make([]int, 4, 8) // capacity is twice the initial size
	// s2 := s1
	// fmt.Printf("Before appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	// appendOne(&s1)
	// fmt.Printf("After appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	// s1[0] = 2
	// s1 := []int{1}
	// fmt.Println("s1 before changeSlice1:", s1)
	// changeSlice1(s1)
	// fmt.Println("s1 after changeSlice1:", s1)
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	fmt.Println(cap(src))
	fmt.Println(src)
	dest1 := append(src, 3)
	fmt.Println(len(src), cap(src))
	fmt.Println(src)
	dest2 := append(src, 4)
	fmt.Println(src, dest1, dest2)
}
