package alphasort

import "cmp"

// OptimizedQuickSort is the entry point for the recursive function.
func OptimizedQuickSort[T cmp.Ordered](arr []T) []T {
	//copy of arr to avoid modifying the original input
	copyArr := make([]T, len(arr))
	copy(copyArr, arr)
	optimizedQuickSortHelper(copyArr, 0, len(copyArr)-1)
	return copyArr
}

func optimizedQuickSortHelper[T cmp.Ordered](arr []T, low, high int) {
	size := high - low + 1

	if size <= 1 {
		return
	}

	// OPTIMIZATION: Check Base Case Network
	// If the partition fits a network size, sort it directly and return [4].
	if shouldApplyNetwork(size) {
		applySortingNetwork(arr[low : high+1])
		return
	}

	// Standard Recursive Step with Median-of-Three and Hoare Partition
	if low < high {
		pivotIndex := medianOfThree(arr, low, high)
		partitionIndex := hoarePartition(arr, low, high, pivotIndex)

		optimizedQuickSortHelper(arr, low, partitionIndex)
		optimizedQuickSortHelper(arr, partitionIndex+1, high)
	}
}

// medianOfThree selects a pivot to avoid worst-case scenarios [4].
func medianOfThree[T cmp.Ordered](arr []T, low, high int) int {
	mid := low + (high-low)/2
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	return mid
}

// hoarePartition implements the partition scheme cited in the methodology [8].
func hoarePartition[T cmp.Ordered](arr []T, low, high int, pivotIdx int) int {
	pivot := arr[pivotIdx]
	i := low - 1
	j := high + 1

	for {
		// Find element on left that should be on right
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		// Find element on right that should be on left
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		// Swap elements
		arr[i], arr[j] = arr[j], arr[i]
	}
}
