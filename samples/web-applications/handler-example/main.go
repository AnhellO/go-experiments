package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func timeHandlerFun(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func timeHandlerClosure(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

func main() {
	mux := http.NewServeMux()

	// Using an existing handler
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	// Using a custom handler
	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	// Using a function handler
	thFun := http.HandlerFunc(timeHandlerFun)
	mux.Handle("/time-func", thFun)

	// Using a function handler with shortcut
	mux.HandleFunc("/time-func/shortcut", timeHandlerFun)

	// Using a function handler with a closure
	mux.HandleFunc("/time-func/closure", timeHandlerClosure(time.RFC1123))

	// Using DefaultServeMux
	// http.Handle("/time", timeHandlerClosure(time.RFC1123))
	// log.Println("Listening...")
	// http.ListenAndServe(":3000")

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
