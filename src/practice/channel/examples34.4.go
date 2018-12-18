package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("producer :", i)
		c <- i
	}

	c <- 100

	// fmt.Println(<-c)
}

func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
	//c <- 1
}
