package pkg

import (
	"os"
)

type authorizationFunc func() bool

type Db struct {
	AuthorizationF authorizationFunc
}

func (d *Db) IsAuthorized() bool {
	return d.AuthorizationF()
}

func NewDB() *Db {
	return &Db{
		AuthorizationF: argsAuthorization,
	}
}

func argsAuthorization() bool {
	user := os.Args[1]
	// super secure authorization
	// in reality, this would be a database call
	if user == "admin" {
		return true
	}
	return false
}
