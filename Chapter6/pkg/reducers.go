package pkg

type Number interface {
	uint | int
}

// Reduce iteratively applies a combining function to the elements of the slice
func Reduce[A any](input []A, reducer func(A1, A2 A) A) A {
	if len(input) == 0 {
		// return default zero
		return *new(A)
	}
	if len(input) == 1 {
		return input[0]
	}

	result := input[0]

	for _, element := range input[1:] {
		result = reducer(result, element)
	}

	return result
}

func Sum[A Number](input []A) A {
	return Reduce(input, func(a1, a2 A) A { return a1 + a2 })
}

func Prod[A Number](input []A) A {
	return Reduce(input, func(a1, a2 A) A { return a1 * a2 })
}
