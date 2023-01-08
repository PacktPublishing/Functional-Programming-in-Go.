package oop

var (
	alphabet [26]rune = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func indexOf(r rune, rs [26]rune) (int, bool) {
	for i := 0; i < len(rs); i++ {
		if r == rs[i] {
			return i, true
		}
	}
	return -1, false
}

type CipherStrategy interface {
	Cipher(string) string
	Decipher(string) string
}

type CipherService struct {
	Strategy CipherStrategy
}

func (c CipherService) Cipher(input string) string {
	return c.Strategy.Cipher(input)
}
func (c CipherService) Decipher(input string) string {
	return c.Strategy.Decipher(input)
}

type CaesarCipher struct {
	Rotation int
}

func (c CaesarCipher) Cipher(input string) string {
	output := ""
	for _, r := range input {
		if idx, ok := indexOf(r, alphabet); ok {
			idx += c.Rotation
			idx = idx % 26
			output += string(alphabet[idx])
		} else {
			output += string(r)
		}
	}
	return output
}

func (c CaesarCipher) Decipher(input string) string {
	output := ""
	for _, r := range input {
		if idx, ok := indexOf(r, alphabet); ok {
			idx += (26 - c.Rotation)
			idx = idx % 26
			output += string(alphabet[idx])
		} else {
			output += string(r)
		}
	}
	return output
}

type AtbashCipher struct {
}

func (a AtbashCipher) Cipher(input string) string {
	output := ""
	for _, r := range input {
		if idx, ok := indexOf(r, alphabet); ok {
			idx = 25 - idx
			output += string(alphabet[idx])
		} else {
			output += string(r)
		}
	}
	return output
}

func (a AtbashCipher) Decipher(input string) string {
	return a.Cipher(input)
}

type CustomCipher struct {
}

func (c CustomCipher) Cipher(input string) string {
	return ""
}

func (c CustomCipher) Decipher(input string) string {
	return ""
}
