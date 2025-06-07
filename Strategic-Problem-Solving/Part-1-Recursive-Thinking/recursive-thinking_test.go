package main

import (
	"testing"
)

func TestMysteriousSequence(t *testing.T) {
	result := MysteriousSequence(15)
	if result != 16384 {
		t.Errorf("Expected: 16384, Got: %d", result)
	}
}

func TestMysteriousSequenceWithMemo(t *testing.T) {
	result := MysteriousSequenceWithMemo(15)

	if result != 16384 {
		t.Errorf("Expected: 16384, Got: %d", result)
	}
}

func TestMysteriousSequenceWithTab(t *testing.T) {
	result := MysteriousSequenceWithTab(15)

	if result != 16384 {
		t.Errorf("Expected: 16384, Got: %d", result)
	}
}

func BenchmarkMysteriousSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MysteriousSequence(20)
	}
}

func BenchmarkMysteriousSequenceWithMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MysteriousSequenceWithMemo(20)
	}
}

func BenchmarkMysteriousSequenceWithTab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MysteriousSequenceWithTab(20)
	}
}

// Time: O(2^n); Space: O(1)
func MysteriousSequence(n int) int {
	if n <= 2 {
		if n == 1 {
			return 1
		}

		if n == 2 {
			return 2
		}

		return 0
	}

	return MysteriousSequence(n-1) + (2 * (MysteriousSequence(n - 2)))
}

// Time: O(n); Space: O(n)
func MysteriousSequenceWithMemo(n int) int {

	if n <= 2 {
		return -1
	}
	var ms func(int) int
	temp := make([]int, n+1)
	temp[1] = 1
	temp[2] = 2

	ms = func(n int) int {
		if n <= 2 {
			if n == 2 {
				return 2
			}
			if n == 1 {
				return 1
			}
			return 0
		}
		if temp[n] != 0 {
			return temp[n]
		}
		temp[n] = ms(n-1) + (2 * ms(n-2))
		return temp[n]
	}

	return ms(n)
}

// Time: O(n); Space: O(1)
func MysteriousSequenceWithTab(n int) int {
	if n < 1 {
		return -1
	}
	prev := 1
	current := 2

	for i := 3; i <= n; i++ {
		next := current + (2 * prev)
		prev = current
		current = next
	}
	return current
}
