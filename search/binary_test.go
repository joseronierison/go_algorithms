package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchOfAnInexistentElement(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 6}

	index := findFirstItemBinarySearch(arr, 9)

	assert.Equal(t, -1, index)
}

func TestSearchOfAnUniqElement(t *testing.T) {
	arr := []int{1, 2, 3, 4, 6}

	index := findFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfDuplicatedElement(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 6}

	index := findFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 2, index)
}

func TestSearchOfManyDuplicatedElement(t *testing.T) {
	arr := []int{3, 3, 3, 3, 3, 3, 3}

	index := findFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 0, index)
}

func TestSearchOfManyDuplicatedOnRightElement(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 3, 3}

	index := findFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 5, index)
}

func TestSearchOfElementRightPosition(t *testing.T) {
	arr := []int{1, 2, 2, 2, 2, 1, 3}

	index := findFirstItemBinarySearch(arr, 3)

	assert.Equal(t, 6, index)
}
