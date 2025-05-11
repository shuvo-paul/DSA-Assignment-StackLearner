package main

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 50, 49}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

	result := InsertionSort(arr)

	if len(result) != len(expected) {
		t.Errorf("expected: %v, Got: %v", len(expected), len(result))
		return
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected: %v, Got: %v", expected[i], result[i])
			return
		}
	}
}

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
	return arr
}
