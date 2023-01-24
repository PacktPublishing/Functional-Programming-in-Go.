package pkg

type Node[A any] func(<-chan A) <-chan A

// could abstract generators and collectors
func ChainPipes[A any](in <-chan A, nodes ...Node[A]) []A {
	var out <-chan A
	for _, node := range nodes {
		out = node(in)
	}
	return CollectorNode(out)
}

func GeneratorNode[A any](nums ...A) <-chan A {
	out := make(chan A)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func FilterNode[A any](in <-chan A, predicate Predicate[A]) <-chan A {
	out := make(chan A)
	go func() {
		for n := range in {
			if predicate(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func MapNode[A any](in <-chan A, mapf MapFunc[A]) <-chan A {
	out := make(chan A)
	go func() {
		for n := range in {
			out <- mapf(n)
		}
		close(out)
	}()
	return out
}

func CollectorNode[A any](in <-chan A) []A {
	output := []A{}
	for n := range in {
		output = append(output, n)
	}
	return output
}
