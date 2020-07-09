package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s2 := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s2[y] = uint8(x * y)
		}
		s[y] = s2
	}

	return s
}

func main() {
	pic.Show(Pic)
}
