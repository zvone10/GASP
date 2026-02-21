package alphasort

import (
	"math/rand"
	"testing"
)

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
