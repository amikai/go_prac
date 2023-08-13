package ds

type Stack[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	Empty() bool
}
