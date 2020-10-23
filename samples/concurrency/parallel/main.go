package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func checkExpire() {
	for {
		// do some job
		fmt.Println(time.Now().UTC())
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go checkExpire()
	http.HandleFunc("/", handler) // http://127.0.0.1:8080/Go
	http.ListenAndServe(":8080", nil)
}
