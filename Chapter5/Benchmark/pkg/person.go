package pkg

type Person struct {
	name string
	age  int
}

//go:noinline
func immutableCreatePerson() Person {
	p := Person{}
	p = immutableSetName(p, "Sean")
	p = immutableSetAge(p, 29)
	return p
}

//go:noinline
func immutableSetName(p Person, name string) Person {
	p.name = name
	return p
}

//go:noinline
func immutableSetAge(p Person, age int) Person {
	p.age = age
	return p
}

//go:noinline
func mutableCreatePerson() *Person {
	p := &Person{}
	mutableSetName(p, "Tom")
	mutableSetAge(p, 31)
	return p
}

//go:noinline
func mutableSetName(p *Person, name string) {
	p.name = name
}

//go:noinline
func mutableSetAge(p *Person, age int) {
	p.age = age
}
