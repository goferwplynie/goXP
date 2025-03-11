package models

import (
	"errors"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
	}
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}

func (ll *LinkedList[T]) Append(value T) error {
	if ll.Head == nil {
		node := NewNode(value)

		if node == nil {
			return errors.New("node not created")
		}

		ll.Head = node
		ll.Tail = node
		return nil
	} else {
		node := NewNode(value)

		if node == nil {
			return errors.New("node not created")
		}

		ll.Tail.Next = node
		ll.Tail = node
	}

	return nil
}

func (ll *LinkedList[T]) Pop() (value T, err error) {
	if ll.Head == nil {
		var zeroValue T
		return zeroValue, errors.New("cant pop from empty list")
	}
	currentNode := ll.Head
	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}
	value = currentNode.Next.Value

	currentNode.Next = nil

	ll.Tail = currentNode

	return value, nil
}

func (ll *LinkedList[T]) GetByIndex(index int) (value T, err error) {
	if ll.Head == nil {
		var zeroValue T
		return zeroValue, errors.New("cant find in empty list")
	}

	currentNode := ll.Head
	for i := 0; i < index; i++ {
		if index-i > 0 && currentNode.Next == nil {
			var zeroValue T
			return zeroValue, errors.New("index out of range")
		}
		currentNode = currentNode.Next
	}
	return currentNode.Value, nil
}

func (ll *LinkedList[T]) IsEmpty() bool {
	if ll.Head == nil {
		return true
	}
	return false
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
