package main

import (
	"fmt"
)

func squares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func cubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 666
	sqrch := make(chan int)
	cubech := make(chan int)

	go squares(number, sqrch)
	go cubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
