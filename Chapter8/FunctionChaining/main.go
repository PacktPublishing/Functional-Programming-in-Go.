package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter8/FunctionChaining/pkg"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doubled := pkg.Map(ints, func(i int) int { return i * 2 })
	larger10 := pkg.Filter(doubled, func(i int) bool { return i >= 10 })
	sum := pkg.Sum(larger10)

	fmt.Println(sum)

	fmt.Println(chaining())

}

func oneliner() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := pkg.Sum(
		pkg.Filter(
			pkg.Map(ints,
				func(i int) int { return i * 2 }),
			func(i int) bool { return i >= 10 }))
	fmt.Println(sum)
}

type ints []int

func (i ints) Map(f func(i int) int) ints {
	return pkg.Map(i, f)
}

func (i ints) Filter(f func(i int) bool) ints {
	return pkg.Filter(i, f)
}

func (i ints) Sum() int {
	return pkg.Sum(i)
}

func chaining() int {
	input := ints([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	return input.Map(func(i int) int { return i * 2 }).
		Filter(func(i int) bool { return i >= 10 }).
		Sum()
}
