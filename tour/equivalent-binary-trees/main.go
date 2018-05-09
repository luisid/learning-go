package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		go Walk(t.Right, ch)
	}

	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int)
	t2Chan := make(chan int)

	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)

	t1Sum := 0
	t2Sum := 0

	for i := 0; i < 20; i++ {
		select {
		case num := <-t1Chan:
			t1Sum += num
		case num := <-t2Chan:
			t2Sum += num
		}
	}
	return t1Sum == t2Sum
}

func main() {
	println(Same(tree.New(10), tree.New(15)))
}
