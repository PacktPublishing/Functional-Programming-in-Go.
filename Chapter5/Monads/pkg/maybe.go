package pkg

import "errors"

type Maybe[A any] interface {
	Get() (A, error)
	GetOrElse(def A) A // gets A, or returns a default value
	IsPresent() bool
	IfPresent(f func(a A)) // if present, execute the function, otherwise do nothing
}

type JustMaybe[A any] struct {
	value A
}

func (j JustMaybe[A]) Get() (A, error) {
	return j.value, nil
}

func (j JustMaybe[A]) GetOrElse(def A) A {
	return j.value
}
func (j JustMaybe[A]) IsPresent() bool {
	return true
}
func (j JustMaybe[A]) IfPresen(f func(a A)) {
	f(j.value)
}

func Just[A any](a A) JustMaybe[A] {
	return JustMaybe[A]{value: a}
}

type NothingMaybe[A any] struct{}

func Nothing[A any]() Maybe[A] {
	return NothingMaybe[A]{}
}
func (n NothingMaybe[A]) Get() (A, error) {
	return *new(A), errors.New("no value")
}

func (n NothingMaybe[A]) GetOrElse(def A) A {
	return def
}
func (n NothingMaybe[A]) IsPresent() bool {
	return false
}
func (m NothingMaybe[A]) IfPresent(f func(a A)) {
	return
}
