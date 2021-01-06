package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1, _ := b64.StdEncoding.DecodeString("c29tZSBzdHJpbmc=")
	err := ioutil.WriteFile("dat1.pdf", d1, 0644)
	check(err)

	f, err := os.Create("dat2.pdf")
	check(err)

	defer f.Close()

	d2, _ := b64.StdEncoding.DecodeString("c29tZSBvdGhlciBzdHJpbmcNCg==")
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
