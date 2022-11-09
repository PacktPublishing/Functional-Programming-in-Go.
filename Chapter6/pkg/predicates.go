package pkg

/*

this package contains predicate based functions

*/

type Predicate[A any] func(A) bool

func Filter[A any](input []A, pred Predicate[A]) []A {
	output := []A{}
	for _, element := range input {
		if pred(element) {
			output = append(output, element)
		}
	}

	return output
}

func Any[A any](input []A, pred Predicate[A]) bool {
	for _, element := range input {
		if pred(element) {
			return true
		}
	}
	return false
}

func All[A any](input []A, pred Predicate[A]) bool {
	for _, element := range input {
		if !pred(element) {
			return false
		}
	}
	return true
}

func DropWhile[A any](input []A, pred Predicate[A]) []A {
	out := []A{}
	drop := true
	for _, element := range input {
		if !pred(element) {
			drop = false
		}
		if !drop {
			out = append(out, element)
		}
	}
	return out
}

func TakeWhile[A any](input []A, pred Predicate[A]) []A {
	out := []A{}
	for _, element := range input {
		if pred(element) {
			out = append(out, element)
		} else {
			return out
		}
	}
	return out
}
