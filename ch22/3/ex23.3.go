package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {

	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	// for e := stack.v.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	v := stack.Pop()

	for v != nil {
		fmt.Printf("%v -> ", v)
		v = stack.Pop()
	}

}
