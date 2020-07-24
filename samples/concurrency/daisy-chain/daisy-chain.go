package main

import "fmt"

func f(left, right chan int) {
	num := <-right
	fmt.Println(num)
	left <- 1 + num
}

func main() {
	const n = 10
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}
