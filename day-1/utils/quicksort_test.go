package utils

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	expectedArray := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	actualArray := QuickSortStart([]int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	if !reflect.DeepEqual(expectedArray, actualArray) {
		t.Errorf("Expected array to be %v, but got %v", expectedArray, actualArray)
	}
}
