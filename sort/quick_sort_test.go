package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSort123Array(t *testing.T) {
	arr := []int{3, 2, 1}
	expectedValue := []int{1, 2, 3}

	sortedArray := QuickSort(arr)

	assert.Equal(t, expectedValue, sortedArray)
}

func TestShouldSortEmptyArray(t *testing.T) {
	arr := []int{}
	expectedValue := []int{}

	sortedArray := QuickSort(arr)

	assert.Equal(t, expectedValue, sortedArray)
}

func TestShouldLargeNumberArray(t *testing.T) {
	arr := []int{9993, 9996, 9992, 0, 9998, 5555, 1000}
	expectedValue := []int{0, 1000, 5555, 9992, 9993, 9996, 9998}

	sortedArray := QuickSort(arr)

	assert.Equal(t, expectedValue, sortedArray)
}
