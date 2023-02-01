package main

import (
	"fmt"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func main() {
	names := pie.FilterNot([]string{"Bob", "Sally", "John", "Jane"},
		func(name string) bool {
			return strings.HasPrefix(name, "J")
		})

	fmt.Println(names) // "[Bob Sally]"

	MyDogs := []Dog{
		Dog{
			"Bucky",
			1,
		},
		Dog{
			"Keeno",
			15,
		},
		Dog{
			"Tala",
			16,
		},
		Dog{
			"Amigo",
			7,
		},
		Dog{
			"Keeno",
			15,
		},
	}
	result := pie.Of(MyDogs).
		Filter(func(d Dog) bool {
			return d.Age > 10
		}).Map(func(d Dog) Dog {
		d.Name = strings.ToUpper(d.Name)
		return d
	}).
		SortUsing(func(a, b Dog) bool {
			return a.Age < b.Age
		})
	fmt.Printf("out: %v\n", result)

	fmt.Println("examples with lo")

	result2 :=
		lo.Map(lo.Uniq(MyDogs), func(d Dog, i int) Dog {
			d.Name = strings.ToUpper(d.Name)
			return d
		})
	fmt.Printf("%v\n", result2)

	fmt.Println("examples with Lo in parallel")
	result3 :=
		lop.Map(lo.Uniq(MyDogs), func(d Dog, i int) Dog {
			d.Name = strings.ToUpper(d.Name)
			return d
		})

	fmt.Printf("%v\n", result3)

}

type Dogs []Dog

type Dog struct {
	Name string
	Age  int
}
