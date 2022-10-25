package main

import "github.com/PacktPublishing/Chapter5/Monads/pkg"

func main() {
	pkg.Just(1)
}

func getFromMap(m map[string]int, key string) pkg.Maybe[int] {
	if value, ok := m[key]; ok {
		return pkg.Just[int](value)
	} else {
		return pkg.Nothing[int]()
	}
}
