package main

func main() {
	// Intentional deadlock!
	ch := make(chan int)
	ch <- 1
}
