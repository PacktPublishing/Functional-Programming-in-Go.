package pkg

//go:generate pie Dogs.*
type Dogs []Dog

type Dog struct {
	Name string
	Age  int
}
