package pkg

import "errors"

type Maybe[A any] interface {
	Get() (A, error)
	GetOrElse(def A) A // gets A, or returns a default value
	IsPresent() bool
}

type JustMaybe[A any] struct {
	value A
}

type NothingMaybe[A any] struct{}

func (j JustMaybe[A]) Get() (A, error) {
	return j.value, nil
}

func (j JustMaybe[A]) GetOrElse(def A) A {
	return j.value
}
func (j JustMaybe[A]) IsPresent() bool {
	return true
}

func Just[A any](a A) JustMaybe[A] {
	return JustMaybe[A]{value: a}
}

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

func fmap[A, B any](m Maybe[A], mapFunc func(A) B) Maybe[B] {
	switch m.(type) {
	case JustMaybe[A]:
		j := m.(JustMaybe[A])
		return JustMaybe[B]{
			value: mapFunc(j.value),
		}
	case NothingMaybe[A]:
		return NothingMaybe[B]{}
	default:
		panic("unknown type")
	}

}
