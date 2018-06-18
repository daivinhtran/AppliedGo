package main

import (
	"fmt"
)

func worker(done chan struct{}) {
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	close(done)
}

func receiver(n int, c <-chan int) {
	for {
		fmt.Printf("Receiver %d received %d\n", n, <-c)
	}
}

func sender(c chan<- int, done chan<- struct{}) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(done)
}

func main() {
	// done := make(chan struct{})
	// go worker(done)
	// fmt.Println("Waiting for the goroutine")
	// <-done
	// fmt.Println("Done")

	c := make(chan int)
	done := make(chan struct{})
	go receiver(1, c)
	go receiver(2, c)
	go sender(c, done)
	<-done
}
