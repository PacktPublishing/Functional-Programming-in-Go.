package main

import "fmt"

type city uint

var (
	Brussels city = iota
	Paris    city
	London   city
	Glasgow  city
	Madrid   city
)

var (
	railroads = [][]city{
		Brussels: []city{Paris},
		Paris:    []city{Brussels, London, Madrid},
		London:   []city{Paris, Glasgow},
	}
)

func main() {
	fmt.Println("hello world")
}
