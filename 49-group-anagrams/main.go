package main

import (
	"fmt"
	"sort"
)

//Runtime: 28 ms, faster than 66.94% of Go online submissions for Group Anagrams.
//Memory Usage: 8.7 MB, less than 26.49% of Go online submissions for Group Anagrams.

func main() {
	example := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(example))
}

type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}

func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	var outputMap = make(map[string][]string)
	var temp = make(RuneSlice, 0)
	for _, str := range strs {
		temp = []rune(str)
		sort.Sort(temp)
		outputMap[string(temp)] = append(outputMap[string(temp)], str)
	}

	var output = make([][]string, len(outputMap))
	var counter = 0
	for _, item := range outputMap {
		output[counter] = item
		counter++
	}

	return output
}
