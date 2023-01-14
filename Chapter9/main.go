package main

import (
	"fmt"

	"github.com/PacktPublishing/Chapter9/pkg/fp"
	"github.com/PacktPublishing/Chapter9/pkg/oop"
)

/*
func main() {
	svc := oop.CipherService{}
	svc.Strategy = oop.CaesarCipher{Rotation: 10}
	fmt.Println(svc.Cipher("helloworld"))
	svc.Strategy = oop.AtbashCipher{}
	fmt.Println(svc.Cipher("helloworld"))

	fpSvc := fp.CipherService{
		CipherFn: func(input string) string {
			return fp.CaesarCipher(input, 10)
		},
		DecipherFn: func(input string) string {
			return fp.CaesarDecipher(input, 10)
		},
	}
	fmt.Println(fpSvc.Cipher("helloworld"))

	fpSvc.CipherFn = fp.AtbashCipher
	fpSvc.DecipherFn = fp.AtbashDecipher

	fmt.Println(fpSvc.Cipher("helloworld"))
	fmt.Println(fpSvc.Decipher(fpSvc.Cipher("hello")))
}
*/

func main() {
	cld := oop.CipherLogDecorator{
		CipherI: oop.CaesarCipher{Rotation: 10},
	}
	ciphered := cld.Cipher("helloworld")
	fmt.Println(ciphered)

	caesarCipher := func(input string) string {
		return fp.CaesarCipher(input, 10)
	}
	caesarDecipher := func(input string) string {
		return fp.CaesarDecipher(input, 10)
	}
	fpSvc := fp.CipherService{
		CipherFn:   fp.LogCipher(caesarCipher),
		DecipherFn: fp.LogDecipher(caesarDecipher),
	}

	fmt.Println(fpSvc.Cipher("hello"))
}
