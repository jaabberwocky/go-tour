package main

import (
	"fmt"
	"strconv"

	"golang.org/x/tour/tree" // imported from tour
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func goWalk(t *tree.Tree, ch chan int) {
	defer close(ch)
	Walk(t, ch)
}

func collectValues(vals *[]int, ch chan int) {
	for v := range ch {
		*vals = append(*vals, v)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go goWalk(t1, ch1)
	go goWalk(t2, ch2)

	vals1, vals2 := []int{}, []int{}

	collectValues(&vals1, ch1)
	collectValues(&vals2, ch2)

	return EqualSlice(vals1, vals2)
}

func EqualSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func PrintTree(t *tree.Tree) string {
	vals, ch := []int{}, make(chan int)
	go goWalk(t, ch)

	for v := range ch {
		vals = append(vals, v)
	}
	return convertValsToString(vals)
}

func convertValsToString(val []int) string {
	s := "Tree:"
	for _, v := range val {
		s += " " + strconv.Itoa(v)
	}
	return s
}

func main() {
	t1, t1b, t2, t3 := tree.New(1), tree.New(1), tree.New(2), tree.New(3)

	fmt.Println(PrintTree(t1))
	fmt.Println(PrintTree(t2))
	fmt.Println(PrintTree(t3))

	fmt.Println(Same(t1, t1b)) // should be true
	fmt.Println(Same(t1, t2))  // should be false

}
