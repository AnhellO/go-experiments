package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
