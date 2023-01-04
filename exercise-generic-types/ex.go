package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l List[T]) printAllNodes() {
	fmt.Print("[")
	isFirst := true

	for l.next != nil {
		if isFirst {
			fmt.Printf("%v", l.val)
			isFirst = !isFirst
		} else {
			fmt.Printf(", %v", l.val)
		}
		l = *l.next
	}
	fmt.Printf(", %v]\n", l.val)
}

func (l List[T]) getLength() (len int) {
	len = 0
	for l.next != nil {
		len += 1
		l = *l.next
	}
	len += 1
	return
}

func (l List[T]) search(searchValue T) bool {
	for l.next != nil {
		if l.val == searchValue {
			return true
		}
		l = *l.next
	}
	if l.val == searchValue {
		return true
	}
	return false
}

func main() {
	l := List[int]{nil, 5}
	l2 := List[int]{nil, 12}
	l3 := List[int]{nil, 0}
	l4 := List[int]{nil, 45}
	l.next = &l2
	l2.next = &l3
	l3.next = &l4

	l.printAllNodes()
	fmt.Println(l.getLength())
	fmt.Println(l.search(0))
	fmt.Println(l.search(-1))
	fmt.Println(l.search(5))

}
