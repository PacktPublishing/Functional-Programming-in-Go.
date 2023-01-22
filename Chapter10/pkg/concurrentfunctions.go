package pkg

import (
	"fmt"
	"math"
)

type Predicate[A any] func(A) bool

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
	fmt.Println(threadCount)
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
