package stack

type IStack interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
	Empty() bool
	Size() int
}
