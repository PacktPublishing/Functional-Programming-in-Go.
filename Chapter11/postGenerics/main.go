package main

import (
	"fmt"
	"strings"

	"github.com/elliotchance/pie/v2"
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
}

type Dogs []Dog

type Dog struct {
	Name string
	Age  int
}
