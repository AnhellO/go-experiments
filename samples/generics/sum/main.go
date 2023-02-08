package main

import "fmt"

type SliceType interface {
	~int | ~float32 | ~float64
}

func SumSliceValues[T SliceType](s []T) T {
	var carry T
	for _, v := range s {
		carry += v
	}

	return carry
}

func main() {
	fmt.Println("Playing with generics!")
	fmt.Println(SumSliceValues([]float32{1.1, 3.3, 5.6}))
}
