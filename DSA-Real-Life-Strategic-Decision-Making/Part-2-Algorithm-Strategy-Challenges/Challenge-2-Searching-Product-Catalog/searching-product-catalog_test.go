package main

import (
	"math"
	"testing"
)

func TestSearchingProductCatalog(t *testing.T) {
	arr := []int{6, 13, 14, 25, 33, 43, 51, 53, 64, 72, 84, 93, 95, 96, 97}

	if value := BinarySearch(arr, 72); value != 72 {
		t.Errorf("Expected: 72, Got: %d", value)
	}

	if value := BinarySearch(arr, 73); value != -1 {
		t.Errorf("Expected: -1, Got: %d", value)
	}
}

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := int(math.Floor(float64(low+high) / 2))

		if arr[mid] == target {
			return target
		}

		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
