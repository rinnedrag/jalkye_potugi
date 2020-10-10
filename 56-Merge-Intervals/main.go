package main

import (
	"fmt"
	"sort"
)

// [[1,4],[0,2],[3,5]]
func main() {
	input := [][]int{{2, 3}, {3, 4}, {4, 5}, {1, 10}}
	result := merge(input)
	fmt.Println(result)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] <= intervals[j][0] {
			return true
		}
		return false
	})
	var result = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		interval, merged := mergeIntervals(result[len(result)-1], intervals[i])
		if !merged {
			result = append(result, interval)
		} else {
			result[len(result)-1] = interval
		}

	}

	return result
}

func mergeIntervals(interval1 []int, interval2 []int) ([]int, bool) {
	if interval1[1] >= interval2[1] {
		return interval1, true
	} else if interval1[1] >= interval2[0] {
		return []int{interval1[0], interval2[1]}, true
	}
	return interval2, false
}
