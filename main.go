package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 3, 4, 5}
	slice = slice[len(slice)-1:]
	fmt.Printf("Sum function took %v", slice)

}
