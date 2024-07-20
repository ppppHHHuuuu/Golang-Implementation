package main

import (
	"sort"
)

func sortArray(a []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}
func sortDefault(a []int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}
func modifyVar(a *int) {
	*a = *a + 1

}

func main() {
	k := 2
	w := 0
	index := 0
	profit := []int{1, 2, 3}
	//sortArray(profit)
	capital := []int{0, 1, 1}

	//sortArray(capital)
	for index <= k {
		for i := index; i < len(capital); i++ {
			if w < capital[index] {
				// update total capital
				w += profit[index]
				break
			}
		}
		index += 1
		//
	}
}
