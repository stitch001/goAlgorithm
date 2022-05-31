package list

import (
	"GoAlgo/myUtils"
	"fmt"
)

type listNode struct {
	Val  interface{}
	Next *listNode
}
type LinkedList struct {
	head *listNode
}

func createLinkedList(vals []interface{}) *LinkedList {
	var pre *listNode
	for i := range vals {
		t := listNode{Val: vals[len(vals)-1-i]}
		t.Next = pre
		pre = &t
	}
	return &LinkedList{head: pre}
}

func (l *LinkedList) printAll() {
	p := l.head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
}

func (l *LinkedList) reverse() {
	var pre, p *listNode
	p = l.head
	for p != nil {
		q := p.Next
		p.Next = pre
		pre = p
		p = q
	}
	l.head = pre
}

func Run() {
	strings := myUtils.ToInterfaceArray("A", "B", "C", "D")
	l1 := createLinkedList(strings)
	l1.printAll()
	fmt.Println("")
	numbers := myUtils.ToInterfaceArray(5, 6, 7, 8, 9)
	l2 := createLinkedList(numbers)
	l2.printAll()
	l2.reverse()
	fmt.Println("")
	l2.printAll()
}
