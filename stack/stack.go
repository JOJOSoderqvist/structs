package stack

import "errors"

type Stack[T any] struct {
	head *node[T]
}

type node[T any] struct {
	Data T
	Next *node[T]
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	newNode := node[T]{
		Data: value,
	}

	if s.head == nil {
		s.head = &newNode

		return
	}

	newNode.Next = s.head
	s.head = &newNode
}

func (s *Stack[T]) Pop() {
	if s.head == nil {
		return
	}

	temp := s.head
	s.head = s.head.Next
	temp.Next = nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.head == nil {
		var zeroValue T

		return zeroValue, errors.New("empty stack")
	}

	return s.head.Data, nil
}

func (s *Stack[T]) Len() uint32 {
	var stackLen uint32
	temp := s.head
	for temp != nil {
		stackLen++
		temp = temp.Next
	}

	return stackLen
}

func (s *Stack[T]) Empty() bool {
	return s.head == nil
}
