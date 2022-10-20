package pkg

type Person struct {
	name string
	age  int
}

func immutableCreatePerson() Person {
	p := Person{}
	p = immutableSetName(p, "Sean")
	p = immutableSetAge(p, 29)
	return p
}

func immutableSetName(p Person, name string) Person {
	p.name = name
	return p
}

func immutableSetAge(p Person, age int) Person {
	p.age = age
	return p
}

func mutableCreatePerson() *Person {
	p := &Person{}
	mutableSetName(p, "Tom")
	mutableSetAge(p, 31)
	return p
}

func mutableSetName(p *Person, name string) {
	p.name = name
}

func mutableSetAge(p *Person, age int) {
	p.age = age
}
