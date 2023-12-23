package algorithm

import (
	"fmt"
	"testing"
)

var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

func TestBubbleSort(t *testing.T) {
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(arr)
	fmt.Println(arr)
}
