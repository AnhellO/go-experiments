package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func _walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		_walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	// 1st part of the exercise
	ch := make(chan int)

	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 2nd part of the exercise
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Printf("Tree(1) vs Tree(1): %v\n", Same(t1, t2)) // Returns true
	t1 = tree.New(1)
	t2 = tree.New(2)
	fmt.Printf("Tree(1) vs Tree(2): %v\n", Same(t1, t2)) // Returns true
}
