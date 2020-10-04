package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-2, 1, -2, 7, -5, 8, -10, 1}
	fmt.Println(maxSubArray(input))
}

func maxSubArray(nums []int) int {
	var output = nums[0]
	var currentMax = math.MinInt32
	for i := 1; i < len(nums); i++ {
		currentMax = max(currentMax, 0) + nums[i]
		output = max(output, currentMax)
	}

	return output
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
