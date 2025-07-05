package main

import "fmt"

func sayHello(numPtr *int) {
	for i := 0; i < 10; i++ {
		*numPtr++
	}
}

func main() {
	num := 0
	sayHello(&num)
	fmt.Println(num)
	nums := []int{1, 2, 3, 4, 5}
	multiply(&nums)
	fmt.Println(nums)
}

func multiply(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
}
