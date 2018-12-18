package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan int)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//defaultCount := 0
	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println("c1 : ", i)
			case s := <-c2:
				fmt.Println("c2 : ", s)
			//case now :=<-time.After(50 * time.Millisecond):
			//	fmt.Println("timeout", now)
			case c3<- 10000:
			//default:
			//	if defaultCount < 3 {
			//		fmt.Println("default")
			//	}
			//	defaultCount++
			}
		}
	}()

	go func() {
		for {
			select {
			case i := <-c3:
				time.Sleep(1 * time.Second)
				fmt.Println("c3 : ", i)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
