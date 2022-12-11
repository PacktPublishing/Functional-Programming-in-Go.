package pkg

type Number interface {
	uint8 | uint16 | uint32 | uint64 | uint |
		int8 | int16 | int32 | int64 | int |
		float32 | float64
}

type (
	reduceFunc[A any] func(a1, a2 A) A
)

// Reduce iteratively applies a combining function to the elements of the slice
func Reduce[A any](input []A, reducer reduceFunc[A]) A {
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

func ReduceRight[A any](input []A, reducer reduceFunc[A]) A {
	if len(input) == 0 {
		// return default zero
		return *new(A)
	}
	if len(input) == 1 {
		return input[0]
	}

	result := input[len(input)-1]

	for i := len(input) - 2; i > 0; i-- {
		result = reducer(result, input[i])
	}

	return result
}

func ReduceWithStart[A any](input []A, startValue A, reducer reduceFunc[A]) A {
	if len(input) == 0 {
		return startValue
	}
	if len(input) == 1 {
		return reducer(startValue, input[0])
	}
	result := reducer(startValue, input[0])
	for _, element := range input[1:] {
		result = reducer(result, element)
	}
	return result
}

func ReduceWithStartRight[A any](input []A, startValue A, reducer reduceFunc[A]) A {
	if len(input) == 0 {
		return startValue
	}
	if len(input) == 1 {
		return reducer(startValue, input[0])
	}
	result := reducer(startValue, input[len(input)-1])
	for i := len(input) - 2; i > 0; i-- {
		result = reducer(result, input[i])
	}
	return result
}

func Sum[A Number](input []A) A {
	return Reduce(input, func(a1, a2 A) A { return a1 + a2 })
}

func Prod[A Number](input []A) A {
	return Reduce(input, func(a1, a2 A) A { return a1 * a2 })
}
