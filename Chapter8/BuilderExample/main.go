package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func newPerson() *person {
	return &person{}
}

func (p *person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *person) SetAge(age int) {
	p.age = age
}

func constructor(firstName, lastName string, age int) person {
	return person{firstName, lastName, age}
}

type personBuilder struct {
	person
}

func (pb personBuilder) FirstName(firstName string) personBuilder {
	pb.person.firstName = firstName
	return pb
}
func (pb personBuilder) LastName(lastName string) personBuilder {
	pb.person.lastName = lastName
	return pb
}
func (pb personBuilder) Age(age int) personBuilder {
	pb.person.age = age
	return pb
}

func (pb personBuilder) Build() person {
	return pb.person
}

func main() {
	alice := newPerson()
	alice.SetFirstName("alice")
	alice.SetLastName("elvi")
	alice.SetAge(30)
	fmt.Println(alice)

	bob := personBuilder{}.FirstName("bob").LastName("doog").Age(20).Build()
	fmt.Println(bob)
}
