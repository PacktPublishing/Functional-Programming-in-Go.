package pkg

import "testing"

const RUNS = 20

func BenchmarkIterative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IterativeFact(RUNS)
	}
}

func BenchmarkRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RecursiveFact(RUNS)
	}
}
