package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(numbers)
	test := numbers[0:3:5]
	fmt.Println(test)
	test = append(test, 44)
	fmt.Println(test)
	test = append(test, 55)
	fmt.Println(test)
	fmt.Println(numbers)
}
