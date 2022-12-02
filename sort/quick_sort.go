package sort

func QuickSort(arr []int) []int {
	array := &arr
	sort(array, 0, len(*array)-1)

	return *array
}

func sort(arr *[]int, beging, end int) {
	if beging < end {
		p := partition(*arr, beging, end)
		sort(arr, beging, p-1)
		sort(arr, p+1, end)
	}
}

func partition(arr []int, beging, end int) int {
	pivot := arr[end]

	i := beging
	for j := beging; j < end; j++ {
		if arr[j] <= pivot {
			vj := arr[j]
			arr[j] = arr[i]
			arr[i] = vj
			i++
		}
	}

	arr[end] = arr[i]
	arr[i] = pivot
	return i
}
