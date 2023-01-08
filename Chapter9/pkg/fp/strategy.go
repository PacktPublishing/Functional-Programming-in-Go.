package fp

var (
	alphabet [26]rune = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func indexOf(r rune, rs [26]rune) int {
	for i := 0; i < len(rs); i++ {
		if r == rs[i] {
			return i
		}
	}
	panic("index not found")
}

type (
	CipherFunc   func(string) string
	DecipherFunc func(string) string
)

type CipherService struct {
	Cipher   CipherFunc
	Decipher DecipherFunc
}

func CaesarCipher(input string, rotation int) string {
	output := ""
	for _, r := range input {
		idx := indexOf(r, alphabet)
		idx += rotation
		idx = idx % 26
		output += string(alphabet[idx])
	}
	return output
}

func CaesarDecipher(input string, rotation int) string {
	output := ""
	for _, r := range input {
		idx := indexOf(r, alphabet)
		idx += (26 - rotation)
		idx = idx % 26
		output += string(alphabet[idx])
	}
	return output
}
