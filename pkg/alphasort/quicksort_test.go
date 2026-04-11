package alphasort

import "testing"

func TestOptimizedQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"Empty array", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Already sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reverse order", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Random order", []int{3, 1, 4, 5, 2, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := OptimizedQuickSort(tt.arr)
			for i := range tt.arr {
				if sorted[i] != tt.want[i] {
					t.Errorf("got %v, want %v", sorted, tt.want)
					break
				}
			}
		})
	}
}
