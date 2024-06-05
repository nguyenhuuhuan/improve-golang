package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 1, 1, 3}
	fmt.Println(findMinimumAscendingSubsequence(arr))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findMinimumAscendingSubsequence(nums []int) []int {
	// Create a map to store the last index of each element in the array
	lastIndexMap := make(map[int]int)
	// Iterate through the array to update the last index of each element
	for i, num := range nums {
		lastIndexMap[num] = i
	}
	// Initialize variables to keep track of the minimum and maximum index
	minIndex := len(nums)
	maxIndex := -1
	// Iterate through the array again to find the minimum and maximum index of repeating elements
	for i, num := range nums {
		if i != lastIndexMap[num] {
			// Element is repeating, update minIndex and maxIndex
			minIndex = min(minIndex, i)
			maxIndex = max(maxIndex, i)
		}
	}
	// Check if there was any repeating element
	if maxIndex == -1 {
		return []int{-1}
	}
	// Return the minimum length ascending subsequence
	return nums[minIndex : maxIndex+1]
}

func maxTrailing(arr []int32) int32 {
	// Write your code here
	min := arr[0]
	diff := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] > min {
			diff = int(math.Max(float64(arr[i]-min), float64(diff)))
		}
		min = int32(math.Min(float64(min), float64(arr[i])))
	}
	return int32(diff)
}
