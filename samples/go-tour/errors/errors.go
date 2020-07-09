package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	del := 0.0000001
	sqrt, prevSqrt := x, x*2
	for prevSqrt-sqrt > del {
		prevSqrt = sqrt
		sqrt = sqrt - (sqrt*sqrt-x)/(2*sqrt)
	}

	return sqrt, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
