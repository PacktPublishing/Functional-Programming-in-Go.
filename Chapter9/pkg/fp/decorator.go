package fp

import "log"

func LogCipher(cipher CipherFunc) CipherFunc {
	return func(input string) string {
		log.Printf("ciphering: %s\n", input)
		return cipher(input)
	}
}

func LogDecipher(decipher DecipherFunc) DecipherFunc {
	return func(input string) string {
		log.Printf("deciphering: %s\n", input)
		return decipher(input)
	}
}
