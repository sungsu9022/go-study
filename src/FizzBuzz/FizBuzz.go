package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		case i%5 == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Println(i)
		}
	}
	fmt.Printf("test")
}
