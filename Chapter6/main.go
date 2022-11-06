package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter6/pkg"
)

type (
	Name   string
	Breed  int
	Gender int
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible genders
const (
	Male Gender = iota
	Female
)

func main() {
	dogs := []Dog{
		Dog{"Bucky", Havanese, Male},
		Dog{"Tipsy", Poodle, Female},
	}
	result := pkg.Filter(dogs, func(d Dog) bool {
		return d.Breed == Havanese
	})
	fmt.Printf("%v\n", result)

	takeWhileDemo()
	dropWhileDemo()

}

func takeWhileDemo() {
	ints := []int{1, 1, 2, 3, 5, 8, 13}
	result := pkg.TakeWhile(ints, func(i int) bool {
		return i%2 != 0
	})
	fmt.Printf("%v\n", result)
}

func dropWhileDemo() {
	ints := []int{1, 1, 2, 3, 5, 8, 13}
	result := pkg.DropWhile(ints, func(i int) bool {
		return i%2 != 0
	})
	fmt.Printf("%v\n", result)
}
