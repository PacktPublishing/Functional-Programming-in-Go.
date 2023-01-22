package pkg

import (
	"fmt"
	"math"
)

type (
	Predicate[A any]   func(A) bool
	MapFunc[A any]     func(A) A
	FMapFunc[A, B any] func(A) B
)

func Filter[A any](input []A, p Predicate[A], out chan []A) {
	output := []A{}
	for _, element := range input {
		if p(element) {
			output = append(output, element)
		}
	}
	out <- output
}

func ConcurrentFilter[A any](input []A, p Predicate[A], batchSize int) []A {
	output := []A{}

	out := make(chan []A)
	threadCount := int(math.Ceil(float64(len(input)) / float64(batchSize)))
	fmt.Printf("goroutines: %d\n", threadCount)
	for i := 0; i < threadCount; i++ {
		fmt.Println("spun up thread")
		if ((i + 1) * batchSize) < len(input) {
			go Filter(input[i*batchSize:(i+1)*batchSize], p, out)
		} else {
			go Filter(input[i*batchSize:], p, out)
		}
	}

	for i := 0; i < threadCount; i++ {
		filtered := <-out
		fmt.Printf("got data: %v\n", filtered)
		output = append(output, filtered...)
	}
	return output
}

// Map is a transformation whereby the type of the element and structure of the container is preserved
func Map[A any](input []A, m MapFunc[A], out chan []A) {
	output := make([]A, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	out <- output
}

func ConcurrentMap[A any](input []A, mapFn MapFunc[A], batchSize int) []A {
	output := []A{}

	out := make(chan []A)
	threadCount := int(math.Ceil(float64(len(input)) / float64(batchSize)))
	fmt.Printf("goroutines: %d\n", threadCount)
	for i := 0; i < threadCount; i++ {
		fmt.Println("spun up thread")
		if ((i + 1) * batchSize) < len(input) {
			go Map(input[i*batchSize:(i+1)*batchSize], mapFn, out)
		} else {
			go Map(input[i*batchSize:], mapFn, out)
		}
	}

	for i := 0; i < threadCount; i++ {
		mapped := <-out
		fmt.Printf("got data: %v\n", mapped)
		output = append(output, mapped...)
	}
	return output
}

func FMap[A, B any](input []A, m func(A) B, out chan []B) {
	output := make([]B, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	out <- output
}

func ConcurrentFMap[A, B any](input []A, fMapFn FMapFunc[A, B], batchSize int) []B {
	output := []B{}

	out := make(chan []B)
	threadCount := int(math.Ceil(float64(len(input)) / float64(batchSize)))
	fmt.Printf("goroutines: %d\n", threadCount)
	for i := 0; i < threadCount; i++ {
		fmt.Println("spun up thread")
		if ((i + 1) * batchSize) < len(input) {
			go FMap(input[i*batchSize:(i+1)*batchSize], fMapFn, out)
		} else {
			go FMap(input[i*batchSize:], fMapFn, out)
		}
	}

	for i := 0; i < threadCount; i++ {
		mapped := <-out
		fmt.Printf("got data: %v\n", mapped)
		output = append(output, mapped...)
	}
	return output
}
