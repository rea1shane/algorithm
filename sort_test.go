package algorithm

import (
	"fmt"
	"reflect"
	"testing"
)

var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
var res = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestBubbleSort(t *testing.T) {
	BubbleSort(arr)
	check()
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(arr)
	check()
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(arr)
	check()
}

func TestQuickSort(t *testing.T) {
	QuickSort(arr)
	check()
}

// check 检测排序是否正确
func check() {
	fmt.Println(reflect.DeepEqual(arr, res))
}
