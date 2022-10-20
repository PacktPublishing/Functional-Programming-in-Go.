package pkg

import "testing"

func BenchmarkImmutablePerson(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		immutableCreatePerson()
	}
}

func BenchmarkMutablePerson(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		mutableCreatePerson()
	}
}
