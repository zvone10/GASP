package alphasort

import (
	"cmp"
)

// applySortingNetwork applies a fixed sequence of compare-swaps for small N.
// In the research, these are assembly-optimized. Here, we simulate the
// logic of optimal sorting networks for sizes 3 through 8.
func applySortingNetwork[T cmp.Ordered](arr []T) {
	n := len(arr)

	// Helper swap function (inlined in assembly, explicit here)
	swap := func(i, j int) {
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	switch n {
	case 3:
		swap(0, 1)
		swap(1, 2)
		swap(0, 1)
	case 4:
		swap(0, 1)
		swap(2, 3)
		swap(0, 2)
		swap(1, 3)
		swap(1, 2)
	case 5:
		swap(0, 1)
		swap(3, 4)
		swap(2, 4)
		swap(2, 3)
		swap(0, 3)
		swap(0, 2)
		swap(1, 4)
		swap(1, 3)
		swap(1, 2)
	case 6:
		// Network for N=6 (12 comparators)
		swap(0, 1)
		swap(2, 3)
		swap(4, 5)
		swap(0, 2)
		swap(3, 5)
		swap(1, 4)
		swap(0, 1)
		swap(2, 3)
		swap(4, 5)
		swap(1, 2)
		swap(3, 4)
		swap(2, 3)
	case 7:
		// Network for N=7 (derived from N=8 network)
		swap(0, 1)
		swap(2, 3)
		swap(4, 5)
		// 6-7 omitted
		swap(0, 2)
		swap(1, 3)
		swap(4, 6)
		// 5-7 omitted
		swap(1, 2)
		swap(5, 6)
		swap(0, 4)
		swap(1, 5)
		swap(2, 6)
		// 3-7 omitted
		swap(2, 4)
		swap(3, 5)
		swap(1, 2)
		swap(3, 4)
		swap(5, 6)
	case 8:
		// Network for N=8 (19 comparators)
		swap(0, 1)
		swap(2, 3)
		swap(4, 5)
		swap(6, 7)
		swap(0, 2)
		swap(1, 3)
		swap(4, 6)
		swap(5, 7)
		swap(1, 2)
		swap(5, 6)
		swap(0, 4)
		swap(3, 7)
		swap(1, 5)
		swap(2, 6)
		swap(1, 4)
		swap(3, 6)
		swap(2, 4)
		swap(3, 5)
		swap(3, 4)
	}
}

// shouldApplyNetwork checks if the size fits the "6-to-8" configuration
// recommended for Merge Sort [6][7], or 3-8 general usage.
func shouldApplyNetwork(n int) bool {
	// The paper highlights "6-to-8" as the sweet spot for Merge Sort,
	// but "3-to-5" works well for Quick Sort on sorted data.
	// We enable 3-8 here to cover all optimizations discussed.
	return n >= 3 && n <= 8
}

// OptimizedMergeSort implements the hybrid merge sort described in the paper.
func OptimizedMergeSort[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	// OPTIMIZATION: Check Base Case Network
	// If the array size is between 3 and 8, use the network immediately [3].
	if shouldApplyNetwork(len(arr)) {
		applySortingNetwork(arr)
		return arr
	}

	// Standard Divide and Conquer
	mid := len(arr) / 2
	left := OptimizedMergeSort(arr[:mid])
	right := OptimizedMergeSort(arr[mid:])

	return merge(left, right)
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
