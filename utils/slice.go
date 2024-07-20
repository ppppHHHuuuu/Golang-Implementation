package utils

import (
	"errors"
	"fmt"
)

type Slice struct {
	slice []int
}

func (mySlice Slice) Print() {
	fmt.Println(mySlice)
}

func (mySlice Slice) Append(a int) {
	mySlice.slice = append(mySlice.slice, a)
}
func (mySlice Slice) LoopingByRange() {
	for index, value := range mySlice.slice {
		fmt.Printf("index %v at value %v", index, value)
	}
}
func (mySlice Slice) LoopingByCStyle() {
	for i := 0; i < len(mySlice.slice); i++ {
		fmt.Printf("index %v at value %v", i, mySlice.slice[i])
	}
}
func (mySlice Slice) Get(index int) (int, error) {
	if index >= len(mySlice.slice) || index < 0 {
		return 0, errors.New("Index out of range")
	}
	return mySlice.slice[index], nil
}
func (mySlice Slice) Update(index int, value int) {
	if index >= len(mySlice.slice) || index < 0 {
		error := fmt.Errorf("Index %v out of range ", index)
		fmt.Println(error)
	}
	mySlice.slice[index] = value
}
func (mySlice Slice) Length() int {
	return len(mySlice.slice)
}

func main() {
	mySlice :=
		Slice{slice: []int{2, 4, 5, 1}}
	mySlice.Append(3)
	mySlice.Print()
}
