package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("c<-i", i)
		}
		close(c)
	}()

	useChanRange(c)
	// useVariable(c)
}

func useChanRange(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func useVariable(c chan int) {
	a, ok := <-c
	fmt.Println(a, ok)
	close(c)
	b, ok := <-c
	fmt.Println(b, ok)
}
