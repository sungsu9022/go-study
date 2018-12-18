package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := num(1, 2)
	out := sumByChan(c)

	fmt.Println("sumByhan func : ", <-out)
	fmt.Println("sum func : ", <-sum(1, 2))
}

func num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()
	return out
}

func sumByChan(c <-chan int) <-chan int {
	out := make(chan int) // 채널 생성
	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼내고 꺼낸 값을 모두 더함.
			r = r + i
		}
		out <- r
	}()
	return out // 채널 반환
}

func sum(a, b int) <-chan int {
	out := make(chan int) // 채널 생성
	go func() {
		out <- a + b // 채널에 a와 b의 합을 보냄
	}()
	return out // 채널 반환
}
