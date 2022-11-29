package search

import (
	"math"
)

/*
 * This methods receives an array of integers and a element to search on this array.
 * It will return the index of the the element in the array (not necessarely the lower index).
 * If it does not find the element, it will return -1.
 * Obs.: It requires a ordered array.
 * Parameters:
 * 	- arr []int (ex.: []int{1,2,3,5,5,6,7,8,9})
 * 	- e int (ex.: 5)
 * 	-> output: 3
 */
func FindItemBinarySearch(arr []int, element int) int {
	lastIndex := len(arr) - 1
	return recursiveBinarySearch(arr, element, 0, lastIndex)
}

/**
 * This methods receives an array of integers and a element to search on this array.
 * It will return the index of the first occurrence of the element in the ordered array.
 * If it does not find the element, it will return -1.
 * Obs.: It requires a ordered array.
 * Parameters:
 * 	- arr []int (ex.: []int{1,2,3,5,5,6,7,8,9})
 * 	- e int (ex.: 5)
 * 	-> output: 4
 */
func FindFirstItemBinarySearch(arr []int, element int) int {
	return abstractBinarySearch(arr, element, firstElement)
}

/**
 * This methods receives an array of integers and a element to search on this array.
 * It will return the index of the last occurrence of the element in the ordered array.
 * If it does not find the element, it will return -1.
 * Obs.: It requires a ordered array.
 * Parameters:
 * 	- arr []int (ex.: []int{1,2,3,5,5,6,5,8,9})
 * 	- e int (ex.: 5)
 * 	-> output: 6
 */
func FindLastItemBinarySearch(arr []int, element int) int {
	return abstractBinarySearch(arr, element, lastElement)
}

type searchType int

const (
	firstElement searchType = iota
	anyElement
	lastElement
)

/**
 * An abstract function for binary search in array search for the first occurrence, the last one
 * or the occurrence in any position.
 */
func abstractBinarySearch(arr []int, element int, searchType searchType) int {
	left := 0
	right := len(arr) - 1
	lastFound := -1

	middle := int(math.Floor((float64(left) + float64(right)) / 2))
	lastIndexLeft := middle - 1
	lastIndexRight := middle - 1

	for left <= right {
		curr := int(math.Floor((float64(left) + float64(right)) / 2))

		if arr[curr] < element {
			left = curr + 1
		} else if arr[curr] > element {
			right = curr - 1
		} else {
			switch searchType {
			case firstElement:
				right = curr - 1
				lastFound = curr
			case anyElement:
				return curr
			case lastElement:
				left = curr + 1
				lastFound = curr
			}
		}

		if curr == lastIndexLeft || curr == lastIndexRight {
			return lastFound
		}
	}

	return lastFound
}

/**
 * That is just another way to do binary search. It uses recursion instead of a for loop.
 */
func recursiveBinarySearch(arr []int, element, begin, end int) int {
	if begin <= end {
		middle := int(math.Floor((float64(begin) + float64(end)) / 2))

		if arr[middle] == element {
			return middle
		}

		if element < arr[middle] {
			return recursiveBinarySearch(arr, element, begin, middle-1)
		} else {
			return recursiveBinarySearch(arr, element, middle+1, end)
		}
	}

	return -1
}
