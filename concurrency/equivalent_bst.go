package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return ;
	}

	Walk(t.Left, ch)
	ch <- t.Value 
	Walk(t.Right, ch)
}

func getTreeStr(ch chan int, s *string) {
	for i := range ch {
		(*s) += string(i)
		(*s) += " "
	}
	fmt.Printf("tree: %s\n", *s)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() { 
		Walk(t1, c1)
		close(c1)
	}()
	go func() {
		Walk(t2, c2)
		close(c2)
	}()

	for {
		v1, ok1 := <- c1
		v2, ok2 := <- c2

		fmt.Printf("Node from tree1 : %d vs. Node from tree2 : %d\n", v1, v2)
		if v1 != v2 || ok1 != ok2 {
			fmt.Println("Not equivalent!")
			return false
		}

		if !ok1 {
			break;
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}