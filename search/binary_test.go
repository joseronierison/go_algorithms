package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// pure binary search tests
func TestSearchOfAnInexistentElementSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 6}

	index := FindItemBinarySearch(arr, 9)

	assert.Equal(t, -1, index)
}

func TestSearchOfAnUniqElementSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6}

	index := FindItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfDuplicatedElementSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 6}

	index := FindItemBinarySearch(arr, 3)

	assert.Equal(t, 3, index)
}

func TestSearchOfManyDuplicatedElementSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{3, 3, 3, 3, 3, 3, 3}

	index := FindItemBinarySearch(arr, 3)

	assert.Equal(t, 3, index)
}

func TestSearchOfManyDuplicatedOnRightElementSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 3, 3}

	index := FindItemBinarySearch(arr, 3)

	assert.Equal(t, 5, index)
}

func TestSearchOfElementRightPositionSearchingForAnyOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 1, 3}

	index := FindItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}

// binary search for first occurrence (worst complexity case) tests
func TestSearchOfAnInexistentElementSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 6}

	index := FindFirstItemBinarySearch(arr, 9)

	assert.Equal(t, -1, index)
}

func TestSearchOfAnUniqElementSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6}

	index := FindFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfDuplicatedElementSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 6}

	index := FindFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfManyDuplicatedElementSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{3, 3, 3, 3, 3, 3, 3}

	index := FindFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 0, index)
}

func TestSearchOfManyDuplicatedOnRightElementSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 3, 3}

	index := FindFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 5, index)
}

func TestSearchOfElementRightPositionSearchingForFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 1, 3}

	index := FindFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}

// binary search for last occurrence (worst complexity case) tests
func TestSearchOfAnInexistentElementSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 6}

	index := FindLastItemBinarySearch(arr, 9)

	assert.Equal(t, -1, index)
}

func TestSearchOfAnUniqElementSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6}

	index := FindLastItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfDuplicatedElementSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 6}

	index := FindLastItemBinarySearch(arr, 3)

	assert.Equal(t, 4, index)
}

func TestSearchOfManyDuplicatedElementSearchingForLastOccurrence(t *testing.T) {
	arr := []int{3, 3, 3, 3, 3, 3, 3}

	index := FindLastItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}

func TestSearchOfManyDuplicatedOnRightElementSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 3, 3}

	index := FindLastItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}

func TestSearchOfElementRightPositionSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 1, 3}

	index := FindLastItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}

func TestSearchOfElementLeftPositionSearchingForLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 3, 4}

	index := FindLastItemBinarySearch(arr, 1)

	assert.Equal(t, 0, index)
}
