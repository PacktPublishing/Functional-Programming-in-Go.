package pkg

import "testing"

func BenchmarkImmutablePerson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		immutableCreatePerson()
	}
}

func BenchmarkMutablePerson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutableCreatePerson()
	}
}
