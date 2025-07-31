package linkedlist

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll == nil {
		t.Error("linekd list not instantiated")
	}
}

func TestNewNode(t *testing.T) {
	node := NewNode(1)

	if node == nil {
		t.Error("node not instantiated")
	}

	if node.Value != 1 {
		t.Errorf("wrong node value %v", node.Value)
	}
}

func TestAppendNoError(t *testing.T) {
	ll := NewLinkedList[int]()

	err := ll.Append(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAppendOneValue(t *testing.T) {
	ll := NewLinkedList[int]()

	node := NewNode(1)

	ll.Append(1)

	if *ll.Head != *node {
		t.Errorf("wrong head! %v", *ll.Head)
	}

	if *ll.Tail != *node {
		t.Errorf("wrong tail! %v", *ll.Tail)
	}
}

func TestAppendMultipleValues(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.Append(1)
	ll.Append(2)

	if ll.Head.Next.Value != 2 {
		t.Errorf("wrong head next value! %v expected 2", ll.Head.Next.Value)
	}

	if ll.Head.Value != 1 {
		t.Errorf("wrong head value! %v expected 1", ll.Head.Value)
	}

	if ll.Tail.Value != 2 {
		t.Errorf("wrong tail value! %v expected 2", ll.Tail.Value)
	}
}

func TestPop(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)

	value, _ := ll.Pop()

	if value != 2 {
		t.Errorf("expected 2 got %v", value)
	}

	if ll.Tail.Value != 1 || ll.Tail.Next != nil {
		t.Errorf("old tail not deleted")
	}
}

func TestPopEmptyList(t *testing.T) {
	ll := NewLinkedList[int]()

	_, err := ll.Pop()

	if err == nil {
		t.Errorf("error expeected. cant pop from empty list")
	}
}

func TestGetIndex(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	v, _ := ll.GetByIndex(1)

	if v != 2 {
		t.Errorf("2 expected got %v", v)
	}
}

func TestGetNotExistingIndex(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	_, err := ll.GetByIndex(999)

	if err == nil {
		t.Error("expected error. not available index")
	}
}

func TestGetIndexFromEmptyList(t *testing.T) {
	ll := NewLinkedList[int]()

	_, err := ll.GetByIndex(32)

	if err == nil {
		t.Error("expected error. cant find in empty list")
	}
}

func TestIsEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll.IsEmpty() != true {
		t.Error("empty list should return true")
	}
	ll.Append(1)
	if ll.IsEmpty() != false {
		t.Error("not empty list should return false")
	}
}
