package pkg

type MapFunc[A any] func(A) A

// Map is a transformation whereby the type of the element and structure of the container is preserved
func Map[A any](input []A, m MapFunc[A]) []A {
	output := make([]A, len(input))

	for i, element := range input {
		output[i] = m(element)
	}

	return output
}

// F stands for "Functor", not "Flat", it maps between types and maintains structure of the
// container
func FMap[A, B any](input []A, m func(A) B) []B {
	output := make([]B, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}

// First maps a single element to a slice of elements (new container structure)
// then flattens the slice by a single level (same container structure as input)
func FlatMap[A any](input []A, m func(A) []A) []A {
	output := []A{}

	for _, element := range input {
		newElements := m(element)
		for _, newEl := range newElements {
			output = append(output, newEl)
		}
	}

	return output
}
