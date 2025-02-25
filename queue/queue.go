package queue

import "errors"

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size uint32
}

type node[T any] struct {
	Data T
	Next *node[T]
}

func New[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Push(value T) {
	newNode := &node[T]{
		Data: value,
	}

	if q.size == 0 {
		q.head = newNode
		q.tail = newNode
		q.size++

		return
	}

	q.tail.Next = newNode
	q.tail = newNode
	q.size++
}

func (q *Queue[T]) Pop() {
	if q.size == 0 {
		return
	}

	temp := q.head
	q.head = q.head.Next
	temp.Next = nil
	q.size--

	if q.head == nil {
		q.tail = nil
	}
}

func (q *Queue[T]) Size() uint32 {
	return q.size
}

func (q *Queue[T]) Front() (T, error) {
	if q.size == 0 {
		var zeroValue T

		return zeroValue, errors.New("queue is empty")
	}

	return q.head.Data, nil
}

func (q *Queue[T]) Back() (T, error) {
	if q.size == 0 {
		var zeroValue T

		return zeroValue, errors.New("queue is empty")
	}

	return q.tail.Data, nil
}
