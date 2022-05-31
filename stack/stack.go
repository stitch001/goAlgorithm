package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	vals []interface{}
}

type MaxMinStack struct {
	Stack
	MaxVal interface{}
	MinVal interface{}
}

func CreateStack() *Stack {
	var s = Stack{}
	s.vals = make([]interface{}, 0)
	return &s
}

func (s *Stack) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack) Push(e interface{}) {
	s.vals = append(s.vals, e)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.vals) == 0 {
		return nil, errors.New("Pop from empty stack")
	}
	lastVal := s.vals[len(s.vals)-1]
	s.vals = s.vals[0 : len(s.vals)-1]
	return lastVal, nil
}

func Run() {
	s := CreateStack()
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Push(3)
	s.Push(4)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Printf("s: %v\n", s)
}
