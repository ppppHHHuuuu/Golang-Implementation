//package main
//
//import (
//	"fmt"
//	"sort"
//	"time"
//)
//
//func getParents(n int, edges [][]int) [][]int {
//	result := make([][]int, n)
//	for i, _ := range edges {
//		temp := -1
//		for j, value := range edges[i] {
//			if j == 0 {
//				temp = value
//			} else {
//				result[value] = append(result[value], temp)
//			}
//		}
//	}
//
//	return result
//}
//
//func setAndSort(arr []int) []int {
//	var resultArr []int
//	freq := make(map[int]int)
//	for i := 0; i < len(arr); i++ {
//		freq[arr[i]] += 1
//		if freq[arr[i]] > 1 {
//			continue
//		} else {
//			resultArr = append(resultArr, arr[i])
//		}
//	}
//	sort.Ints(resultArr)
//
//	return resultArr
//}
//func getAncestorDFS(parents [][]int, i int) []int {
//	numParent := len(parents[i])
//	if numParent == 0 {
//		return []int{}
//	}
//
//	var ancestor []int
//	ancestor = append(ancestor, parents[i]...)
//	for index, _ := range parents[i] {
//		key := parents[i][index]
//		//ancestor = append(ancestor, parents[key]...)
//		ancestor = append(ancestor, getAncestorDFS(parents, key)...)
//	}
//	//ancestor = setAndSort(ancestor)
//	return ancestor
//}
//func getAncestorBFS(parents [][]int, i int) ([]int, error) {
//	var ancestor []int
//	ancestor = append(ancestor, parents[i]...)
//
//	stack := Stack.NewStack(30)
//	for _, value := range parents[i] {
//		err := stack.Push(value)
//		if err != nil {
//			return nil, err
//		}
//	}
//	for !stack.IsEmpty() {
//		key, _ := stack.Pop()
//		if len(parents[key]) != 0 {
//			for index := 0; index < len(parents[key]); index++ {
//				stack.Push(parents[key][index])
//			}
//		}
//		ancestor = append(ancestor, parents[key]...)
//	}
//
//	return ancestor, nil
//}
//func getAncestorsFromParents(n int, parents [][]int) [][]int {
//	ancestors := make([][]int, n)
//	for i, _ := range parents {
//		ancestors[i], _ = getAncestorBFS(parents, i)
//		ancestors[i] = setAndSort(ancestors[i])
//	}
//	return ancestors
//}
//func getAncestors(n int, edges [][]int) [][]int {
//	result := getParents(n, edges)
//	ancestors := getAncestorsFromParents(n, result)
//	return ancestors
//}
//func main() {
//	start := time.Now()
//
//	var n = 8
//	//var edges = [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
//	var edges = [][]int{{0, 7}, {7, 6}, {0, 3}, {6, 3}, {5, 4}, {1, 5}, {2, 7}, {3, 5}, {3, 1}, {0, 5}, {7, 5}, {2, 1}, {1, 4}, {6, 1}}
//	result := getParents(n, edges)
//	ancestors := getAncestorsFromParents(n, result)
//	elapsed := time.Since(start)
//
//	fmt.Println(result)
//	fmt.Println(ancestors)
//	fmt.Printf("Sum function took %s", elapsed)
//
//}
