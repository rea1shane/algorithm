package algorithm

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] ^ arr[j+1]
				arr[j+1] = arr[j] ^ arr[j+1]
				arr[j] = arr[j] ^ arr[j+1]
			}
		}
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i] = arr[i] ^ arr[minIndex]
			arr[minIndex] = arr[i] ^ arr[minIndex]
			arr[i] = arr[i] ^ arr[minIndex]
		}
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}
