package main

import "fmt"

func main() {
	var items = []int{4, 5, 6, 7, 8, 0, 1, 2}
	fmt.Println(search(items, 7))
}

func search(nums []int, target int) int {
	sortNums, shift := sortInput(nums)
	index, ok := binarySearch(sortNums, target)
	if !ok {
		return -1
	}
	return (index + shift) % len(nums)
}

func sortInput(input []int) (result []int, shift int) {
	result = make([]int, 0)
	for i, num := range input {
		if i == 0 {
			continue
		}
		if num < input[i-1] {
			shift = i
			result = input[i:len(input)]
			result = append(result, input[:i]...)
			return
		}
	}
	return input, 0
}

func binarySearch(nums []int, target int) (result int, ok bool) {
	mid := len(nums) / 2
	ok = true
	switch {
	case len(nums) == 0:
		result = -1
		ok = false
	case nums[mid] > target:
		result, ok = binarySearch(nums[:mid], target)
	case nums[mid] < target:
		result, ok = binarySearch(nums[mid+1:], target)
		result += mid + 1
	case nums[mid] == target:
		result = mid
	}

	return
}
