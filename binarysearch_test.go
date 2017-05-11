package binarysearch

import "testing"

const Max = 1000000

func generateSlice() []int {
	s := make([]int, Max)
	for x := 0; x < Max; x++ {
		s[x] = x
	}
	return s
}

func BenchmarkSearch(b *testing.B) {
	s := generateSlice()
	b.Run("Div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BinarySearchDiv(s, Max)
		}
	})

	b.Run("Shift", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BinarySearch(s, Max)
		}
	})
}
