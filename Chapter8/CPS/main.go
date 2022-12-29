package main

import "fmt"

func factorial(n int, cont func(int)) {
	if n == 0 {
		cont(1)
	} else {
		factorial(n-1, func(v int) {
			cont(n * v)
		})
	}
}

func main() {
	factorial(5, func(v int) {
		fmt.Println(v)
	})

	callback := func(input int, b bool) {
		if b {
			fmt.Printf("the number %v is even\n", input)
		} else {
			fmt.Printf("the number %v is odd\n", input)
		}
	}

	for i := 0; i < 10; i++ {
		go isEven(i, callback)
	}
	_ := <-make(chan int)
}

func isEven(i int, callback func(int, bool)) {
	if i%2 == 0 {
		callback(i, true)
	} else {
		callback(i, false)
	}
}
