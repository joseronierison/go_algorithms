package search

import (
	"math"
)

func findFirstItemBinarySearch(arr []int, e int) int {
	left := 0
	right := len(arr) - 1
	lastFound := -1

	middle := int(math.Floor((float64(left) + float64(right)) / 2))
	lastIndexLeft := middle - 1
	lastIndexRight := middle - 1

	for left <= right {
		curr := int(math.Floor((float64(left) + float64(right)) / 2))

		if arr[curr] < e {
			left = curr + 1
		} else if arr[curr] > e {
			right = curr - 1
		} else {
			right = curr - 1
			lastFound = curr
		}

		if curr == lastIndexLeft || curr == lastIndexRight {
			return lastFound
		}
	}

	return lastFound
}
