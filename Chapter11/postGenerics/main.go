package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/samber/mo"
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

	fmt.Println("examples with Mo")

	maybe := mo.Some(Dog{"Bucky", 1})
	getOrElse := maybe.OrElse(Dog{})
	fmt.Println(getOrElse)

	maybe2 := mo.None[Dog]()
	getOrElse2 := maybe2.OrElse(Dog{"Default", -1})
	fmt.Println(getOrElse2)

	ok := mo.Ok(MyDogs[0])
	result1 := ok.OrElse(Dog{})
	err1 := ok.Error()

	fmt.Println(result1, err1)

	err := errors.New("dog not found")
	ok2 := mo.Err[Dog](err)
	result4 := ok2.OrElse(Dog{"Default", -1})
	err2 := ok2.Error()

	fmt.Println(result4, err2)

}

type Dogs []Dog

type Dog struct {
	Name string
	Age  int
}
