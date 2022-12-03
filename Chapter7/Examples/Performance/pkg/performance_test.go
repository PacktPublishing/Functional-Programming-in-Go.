package pkg

import "testing"

func BenchmarkIterative100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IterativeFact(10)
	}
}

func BenchmarkRecursive100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RecursiveFact(10)
	}
}
