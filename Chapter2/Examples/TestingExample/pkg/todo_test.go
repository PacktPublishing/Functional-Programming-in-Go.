package pkg_test

import (
	"testing"

	"github.com/DylanMeeus/Book-FP/pkg"
)

func TestTodoWrite(t *testing.T) {
	todo := pkg.Todo{
		Db: &pkg.Db{
			AuthorizationF: func() bool { return true },
		},
	}
	todo.Write("hello")
	if todo.Text != "hello" {
		t.Errorf("Expected 'hello' but got %v\n", todo.Text)
	}
	todo.Append(" world")
	if todo.Text != "hello world" {
		t.Errorf("Expected 'hello world' but got %v\n", todo.Text)
	}
}
