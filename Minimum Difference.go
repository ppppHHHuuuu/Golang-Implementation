package main

import (
	"fmt"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	diffThreeSmallest := nums[2] - nums[0]
	diffThreeLargest := nums[len(nums)-1] - nums[len(nums)-3]
	if diffThreeSmallest < diffThreeLargest {
		return nums[len(nums)-4] - nums[0]
	}
	return nums[len(nums)-1] - nums[3]
}
func main() {
	//nums := []int{1, 5, 7, 10, 20, 40, 300}
	nums := []int{6, 6, 0, 1, 1, 4, 6}
	fmt.Println(minDifference(nums))
}
