package main

import (
	"fmt"
)

func digits(number int, ch chan int) {
	for number != 0 {
		digit := number % 10
		ch <- digit
		number /= 10
	}

	close(ch)
}

func calcSquares(number int, square chan int) {
	sum := 0
	ch := make(chan int)
	go digits(number, ch)

	for digit := range ch {
		sum += digit * digit
	}

	square <- sum
}

func calcCubes(number int, cube chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}

	cube <- sum
}

func main() {
	number := 999
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	fmt.Printf("Final output: %d\n", squares+cubes)
}
