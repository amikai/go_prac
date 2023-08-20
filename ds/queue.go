package ds

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Peek() (T, error)
	Empty() bool
}
