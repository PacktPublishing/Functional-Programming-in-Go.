package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter8/LazyEvaluation/pkg"
)

func main() {
	largerThan10Mil := func(i int) bool {
		return i > 10_000_000
	}
	res := ints(IntRange(0, 100)).
		Map(Factorial).
		Filter(largerThan10Mil).
		Head()

	fmt.Printf("%v\n", res)

}

type ints []int

func (i ints) Head() pkg.Maybe[int] {
	return pkg.Head(i)
}

func (i ints) Map(f func(i int) int) ints {
	return pkg.Map(i, f)
}

func (i ints) Filter(f func(i int) bool) ints {
	return pkg.Filter(i, f)
}
func IntRange(start, end int) []int {
	out := []int{}
	for i := start; i <= end; i++ {
		out = append(out, i)
	}
	return out
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
