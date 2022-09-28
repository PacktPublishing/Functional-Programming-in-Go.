package main

import "fmt"

type calculator struct {
	x int
	y int
}

/*
 calculate supports operations: +,-,/,*
*/
func calculate(a, b int, operation string) int {
	switch operation {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mult(a, b)
	case "/":
		return div(a, b)
	default:
		panic("operation not supported")
	}
}

type calculateFunc func(int, int) int

var (
	operations = map[string]calculateFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a + b
}

func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func calculateWithMap(a, b int, opString string) int {
	if operation, ok := operations[opString]; ok {
		return operation(a, b)
	}
	panic("operation not supported")
}

func main() {
	fmt.Println("Hello Go, from iPad Pro!")
}
