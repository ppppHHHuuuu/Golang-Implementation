package main

import (
	"fmt"
	"sort"
)

func findMaxNum(nums []int) int {
	maxNum := -1
	for i := 0; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func assignNumsToMap(nums []int, maxNum int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < maxNum; i++ {
		m[i] = 0
	}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	return m
}
func assignNumsToArray(nums []int, maxNum int) []int {
	arr := make([]int, 0, maxNum)
	for i := 0; i <= maxNum; i++ {
		arr = append(arr, 0)
	}
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += 1
	}
	return arr
}
func mainWithMap(nums []int) {
	maxNum := findMaxNum(nums)
	m := assignNumsToMap(nums, maxNum)

	count := 0
	temp := 0

	keys := make([]int, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
		m[k] += temp
		fmt.Println("m ", m)

		temp = 0
		for m[k] > 1 {
			m[k]--
			count += 1
			temp += 1
		}
	}
	fmt.Println(count)
}
func mainWithArray(nums []int) {
	maxNum := findMaxNum(nums)
	array := assignNumsToArray(nums, maxNum)

	count := 0
	temp := 0

	//for index, value := range array {
	//	fmt.Println(index, " ", value)
	//
	//	if index == len(array)-1 {
	//
	//		if value > 1 {
	//			array = append(array, 0)
	//		}
	//		fmt.Println("array new", array)
	//
	//	}
	//	array[index] += temp
	//	fmt.Println("array ", array)
	//
	//	temp = 0
	//	for array[index] > 1 {
	//		array[index]--
	//		count += 1
	//		temp += 1
	//	}
	//}
	for index := 0; index < len(array); index++ {
		fmt.Println(index, " ", array[index])
		array[index] += temp
		fmt.Println("array ", array)
		temp = 0

		if index == len(array)-1 {

			if array[index] > 1 {
				array = append(array, 0)
			}
			fmt.Println("array new", array)

		}
		for array[index] > 1 {
			array[index]--
			count += 1
			temp += 1
		}
	}
	fmt.Println(count)
}

//func main() {
//	//nums := []int{3, 2, 1, 2, 1, 7}
//	//nums := []int{2, 1, 1, 1}
//	//mainWithArray(nums)
//	//fmt.Println(assignNumsToArray(nums, 2))
//	var counts = make([]int, 0, 2000)
//	fmt.Println((counts))
//}
