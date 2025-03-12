package stack

type Stack[T any] interface {
	Append(value T) error
	Pop() (value T, err error)
	IsEmpty() bool
	// for iterating
	Range() <-chan T
}
