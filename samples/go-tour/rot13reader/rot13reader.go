package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rt.r.Read(b)

	for i := 0; i < n; i++ {
		if b[i] >= 65 && b[i] <= 77 || b[i] >= 97 && b[i] <= 109 {
			b[i] += 13
		} else if b[i] > 77 && b[i] <= 90 || b[i] > 109 && b[i] <= 122 {
			b[i] -= 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
