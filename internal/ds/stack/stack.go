package stack

type Stack[T any] interface {
	Push(value T)
	Pop() (value T, err error)
	IsEmpty() bool
}
