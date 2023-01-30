package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter11/pregenerics/pkg"
	"github.com/elliotchance/pie/pie"
)

func main() {
	out := pie.Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		Map(func(i int) int { return i * i })

	fmt.Printf("result: %v\n", out)

	MyDogs := []pkg.Dog{
		pkg.Dog{
			"Bucky",
			1,
		},
		pkg.Dog{
			"Keeno",
			15,
		},
		pkg.Dog{
			"Tala",
			16,
		},
		pkg.Dog{
			"Amigo",
			7,
		},
	}

	results := pkg.Dogs(MyDogs).
		Filter(func(d pkg.Dog) bool {
			return d.Age > 10
		}).SortUsing(func(a, b pkg.Dog) bool {
		return a.Age < b.Age
	})
	fmt.Printf("results: %v\n", results)
}
