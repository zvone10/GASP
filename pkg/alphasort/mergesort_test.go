package alphasort

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimizedMergeSort(t *testing.T) {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"Empty array", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Already sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sortedArray},
		{"Reverse order", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, sortedArray},
		{"Random order", []int{3, 1, 4, 5, 2, 6, 7, 8, 9, 10}, sortedArray},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := OptimizedMergeSort(tt.arr)
			assert.Equal(t, tt.want, sorted)
		})
	}
}

func TestLargeMergeSort(t *testing.T) {
	arrLen := 100000
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(arrLen)
	}

	sorted := OptimizedMergeSort(arr)
	for i := 1; i < len(sorted); i++ {
		if sorted[i-1] > sorted[i] {
			fmt.Println(sorted)
			t.Fatalf("Array is not sorted at index %d: %d > %d", i, sorted[i-1], sorted[i])
		}
		// Alternatively, using testify's assertion:
		//
		assert.LessOrEqual(t, sorted[i-1], sorted[i], "Array is not sorted at index %d", i)
	}
}

func BenchmarkOptimizedMergeSort(b *testing.B) {
	arrLen := 600000
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(arrLen)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Create a copy to ensure each benchmark iteration sorts the same initial data
		tempArr := make([]int, arrLen)
		copy(tempArr, arr)

		b.StartTimer()
		OptimizedMergeSort(tempArr)
	}
}
