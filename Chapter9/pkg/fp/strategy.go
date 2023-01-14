package fp

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

type (
	CipherFunc   func(string) string
	DecipherFunc func(string) string
)

type CipherService struct {
	CipherFn   CipherFunc
	DecipherFn DecipherFunc
}

func (c CipherService) Cipher(input string) string {
	return c.CipherFn(input)
}

func (c CipherService) Decipher(input string) string {
	return c.DecipherFn(input)
}

func CaesarCipher(input string, rotation int) string {
	output := ""
	for _, r := range input {
		if idx, ok := indexOf(r, alphabet); ok {
			idx += rotation
			idx = idx % 26
			output += string(alphabet[idx])
		} else {
			output += string(r)
		}
	}
	return output
}

func CaesarDecipher(input string, rotation int) string {
	output := ""
	for _, r := range input {
		if idx, ok := indexOf(r, alphabet); ok {
			idx += (26 - rotation)
			idx = idx % 26
			output += string(alphabet[idx])
		} else {
			output += string(r)
		}
	}
	return output
}

func AtbashCipher(input string) string {
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

var AtbashDecipher = AtbashCipher
