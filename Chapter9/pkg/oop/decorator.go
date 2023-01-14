package oop

import (
	"log"
)

type CipherLogDecorator struct {
	CipherI CipherStrategy
}

func (c CipherLogDecorator) Cipher(input string) string {
	log.Printf("ciphering: %s\n", input)
	return c.CipherI.Cipher(input)
}

func (c CipherLogDecorator) Decipher(input string) string {
	log.Printf("deciphering: %s\n", input)
	return c.CipherI.Decipher(input)
}
