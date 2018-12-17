package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i :=0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		if i%2 == 0 {
			time.Sleep(1 * time.Second)
		}
		<-done
		fmt.Println("메인 : ", i)
	}

	fmt.Scanln()
}
