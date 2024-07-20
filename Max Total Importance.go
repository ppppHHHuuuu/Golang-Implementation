package main

import (
	"fmt"
	"sort"
)

// with each city get its number of road
func freqOfCities(roads [][]int) map[int]int {
	var freq = make(map[int]int)

	for i, _ := range roads {
		for _, value := range roads[i] {
			freq[value] += 1
		}
	}
	return freq
}
func orderedImportance(n int, freq map[int]int) []int {
	var arr []int
	var lenFreq = len(freq)
	for _ = range n - lenFreq {
		arr = append(arr, 0)
	}
	for _, value := range freq {
		arr = append(arr, value)
	}
	sort.Ints(arr)
	return arr
}
func maximumImportances(arr []int) int64 {
	var count = 1
	var res = int64(0)

	for _, value := range arr {
		res += int64(count) * int64(value)
		count += 1
	}
	return res
}
func maximumImportance(n int, roads [][]int) int64 {
	var freqs map[int]int = freqOfCities(roads)
	var orderImportance []int = orderedImportance(n, freqs)
	var res int64 = maximumImportances(orderImportance)
	return res
}

func main() {
	//roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	//roads := [][]int{{0, 3}, {2, 4}, {1, 3}}
	roads := [][]int{{0, 1}}
	fmt.Println(maximumImportance(5, roads))
}
