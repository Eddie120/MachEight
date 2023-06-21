package main

import (
	"fmt"
)

func createKeys(complement, num int) (string, string) {
	return fmt.Sprintf("%x-%x", complement, num), fmt.Sprintf("%x-%x", num, complement)
}

func sum(nums []int, target int) []string { // BigO = n
	var table = map[int]bool{}
	var hasUsed = map[string]bool{}
	var output = make([]string, 0)

	for i := 0; i < len(nums); i++ { // O(n)
		num := nums[i]
		table[num] = true
	}

	for i := 0; i < len(nums); i++ { // O(n)
		num := nums[i]
		complement := target - num
		// this is used to track used values in the hash table
		// disable the same key in inverse way
		key, inverseKey := createKeys(complement, num)
		if table[complement] && !hasUsed[key] {
			hasUsed[key] = true
			hasUsed[inverseKey] = true
			output = append(output, fmt.Sprintf("%d,%d", complement, num))
		}
	}

	return output
}
