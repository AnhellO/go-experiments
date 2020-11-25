package main

import (
	"fmt"
)

func EvenNumbersTillEight(even chan int) {
	i := 2
	for i < 9 {
		even <- i
		i = i + 2
	}
	close(even)
}

func OddNumberTillEight(odd chan int) {
	i := 1
	for i < 9 {
		odd <- i
		i = i + 2
	}
	close(odd)
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	go EvenNumbersTillEight(even)
	go OddNumberTillEight(odd)
	for {
		even, ok1 := <-even
		odd, ok2 := <-odd
		if ok1 == false && ok2 == false {
			break
		}
		fmt.Println("Received ", even, ok1, odd, ok2)
	}
}
