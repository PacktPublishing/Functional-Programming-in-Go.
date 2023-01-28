package pkg

import (
	"io/ioutil"
	"strings"
)

type (
	Node[A any]          func(<-chan A) <-chan A
	GeneratorNode[A any] func() <-chan A
	CollectorNode[A any] func(<-chan A) []A
)

func CurriedFilterNode[A any](p Predicate[A]) Node[A] {
	return func(in <-chan A) <-chan A {
		out := make(chan A)
		go func() {
			for n := range in {
				if p(n) {
					out <- n
				}
			}
			close(out)
		}()
		return out
	}
}

func CurriedMapNode[A any](mapFn MapFunc[A]) Node[A] {
	return func(in <-chan A) <-chan A {
		out := make(chan A)
		go func() {
			for n := range in {
				out <- mapFn(n)
			}
			close(out)
		}()
		return out
	}
}

// could abstract generators and collectors
func ChainPipes[A any](in <-chan A, nodes ...Node[A]) []A {
	for _, node := range nodes {
		in = node(in)
	}
	return Collector(in)
}

func ChainPipes2[A any](gn GeneratorNode[A], nodes ...Node[A]) []A {
	in := gn()
	for _, node := range nodes {
		in = node(in)
	}
	return Collector(in)
}

func Generator[A any](input ...A) <-chan A {
	out := make(chan A)
	go func() {
		for _, element := range input {
			out <- element
		}
		close(out)
	}()
	return out
}

func CurriedCat(filepath string) func() <-chan string {
	return func() <-chan string {
		out := make(chan string)
		f, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		go func() {
			lines := strings.Split(string(f), "\n")
			for _, line := range lines {
				out <- line
			}
			close(out)
		}()
		return out
	}

}

func Cat(filepath string) <-chan string {
	out := make(chan string)
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	go func() {
		lines := strings.Split(string(f), "\n")
		for _, line := range lines {
			out <- line
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

func Collector[A any](in <-chan A) []A {
	output := []A{}
	for n := range in {
		output = append(output, n)
	}
	return output
}
