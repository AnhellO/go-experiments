package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		n1, n2 = n2, n1+n2
		return n2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i == 0 || i == 1 {
			fmt.Println(i)
			continue
		}
		fmt.Println(f())
	}
}
