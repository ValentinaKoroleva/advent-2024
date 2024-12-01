package main

import (
	"fmt"
	"sort"
)

// To find the total distance between the left list and the right list
func findTotalDistance(left []int, right []int) {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += Abs(left[i] - right[i])
	}
	fmt.Println(totalDistance)
}

// To find the similarity between the left list and the right list
func findSimilarity(left []int, right []int) {
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	// find frequency of each element in right array
	freq := make(map[int]int)
	for i := 0; i < len(right); i++ {
		freq[right[i]]++
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		if left[i] >= right[0] && left[i] <= right[len(right)-1] {
			similarity += left[i] * freq[left[i]]
		}
	}
	fmt.Println(similarity)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
