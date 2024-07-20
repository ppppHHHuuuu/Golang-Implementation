package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	check := false
	for i := 1; i <= c; i++ {
		fmt.Println(i)
		remain := c - i*i
		fmt.Println(remain)
		if remain > 0 && math.Sqrt(float64(remain))-math.Round(math.Sqrt(float64(remain))) == 0 {
			fmt.Println(remain, " ", i)
			check = true
		}
	}
	return check
}

func main() {
	fmt.Println(judgeSquareSum(4))
}
