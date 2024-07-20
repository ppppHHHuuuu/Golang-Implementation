package utils

type Operation interface {
	Size() uint
	IsEmpty() bool
	Clear()
	Print()
	Pop() (int, error)
	Peek() (int, error)
}
