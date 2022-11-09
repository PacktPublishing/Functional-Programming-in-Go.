package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	multiplyMapDemo()
	dogMapDemo()
	flatMapDemo()
	sumDemo()
	reduceWithStartDemo()

	exampleAirlineData()
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

func multiplyMapDemo() {
	ints := []int{1, 1, 2, 3, 5, 8, 13}
	result := pkg.Map(ints, func(i int) int {
		return i * 2
	})
	fmt.Printf("%v\n", result)
}

func dogMapDemo() {
	dogs := []Dog{
		Dog{"Bucky", Havanese, Male},
		Dog{"Tipsy", Poodle, Female},
	}
	result := pkg.Map(dogs, func(d Dog) Dog {
		if d.Gender == Male {
			d.Name = "Mr. " + d.Name
		} else {
			d.Name = "Mrs. " + d.Name
		}
		return d
	})
	fmt.Printf("%v\n", result)
}

func flatMapDemo() {
	ints := []int{1, 2, 3}
	result := pkg.FlatMap(ints, func(n int) []int {
		out := []int{}
		for i := 0; i < n; i++ {
			out = append(out, i)
		}
		return out
	})
	fmt.Printf("%v\n", result)
}

func sumDemo() {
	ints := []int{1, 2, 3, 4}
	result := pkg.Sum(ints)
	fmt.Printf("%v\n", result)
}

func reduceWithStartDemo() {
	words := []string{"hello", "world", "universe"}
	result := pkg.ReduceWithStart(words, "first", func(s1, s2 string) string {
		return s1 + ", " + s2
	})
	fmt.Printf("%v\n", result)
}

type Entry struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`

	// time
	Time struct {
		Label     string `json:"Label"`
		Month     uint   `json:"Month"`
		MonthName string `json:"Month Name"`
		Year      uint   `json:"Year"`
	} `json:"Time"`

	// statistics
	Statistics struct {
		NumberOfDelays struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Weather                int `json:"Weather"`
		} `json:"# of Delays"`
		// Carriers
		Carriers struct {
			Names string `json:"Names"`
			Total int    `json:"Total"`
		} `json:"Carriers"`

		// Flights
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed   int `json:"Delayed"`
			Diverted  int `json:"Diverted"`
			OnTime    int `json:"On Time"`
			Total     int `json:"Total"`
		} `json:"Flights"`

		MinutesDelayed struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Weather                int `json:"Weather"`
		} `json:"Minutes Delayed"`
	} `json:"Statistics"`
}

func exampleAirlineData() {
	bytes, err := ioutil.ReadFile("./resources/airlines.json")
	if err != nil {
		panic(err)
	}

	var entries []Entry
	err = json.Unmarshal(bytes, &entries)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", entries[0])

}
