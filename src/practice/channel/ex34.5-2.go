package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)

	go func() {
		for {
			i := <-c1
			fmt.Println("첫번째 고루틴 c1 :", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c1 <- 20
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10:
			case i := <-c1:
				fmt.Println("select c1 : ", i)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
