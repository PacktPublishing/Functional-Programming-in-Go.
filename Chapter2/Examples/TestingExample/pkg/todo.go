package pkg

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Todo struct {
	Text string
	Db   *Db
}

func NewTodo() Todo {
	return Todo{
		Text: "",
		Db:   NewDB(),
	}
}

func (t *Todo) Write(s string) {
	if t.Db.IsAuthorized() {
		t.Text = s
	} else {
		panic("user not authorized to write")
	}
}

func (t *Todo) Append(s string) {
	if t.Db.IsAuthorized() {
		t.Text += s
	} else {
		panic("user not authorized to append")
	}
}
