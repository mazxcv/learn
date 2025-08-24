package atbash

import "strings"

// func CipherAtbash make Atbash-cipher to symbol
func CipherAtbash(s rune) rune {
	if s >= 'a' && s <= 'z' {
		return 'z' - (s - 'a')
	}
	return s
}

// func prepareForAtbash make string of atbash-word
func prepareForAtbash(s string) string {
	const maxAtbashLengthWord = 5
	var sb strings.Builder
	s = strings.ToLower(s)
	flCnt := 0
	for _, v := range strings.Trim(strings.ToLower(s), ",.:;") {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {

			if flCnt == maxAtbashLengthWord {
				sb.WriteRune(' ')
				flCnt = 0
			}
			sb.WriteRune(v)
			flCnt++
		}
	}

	return sb.String()

}

// func Atbash prepare atbash cipher to string
func Atbash(s string) string {
	var sb strings.Builder
	for _, v := range prepareForAtbash(s) {
		sb.WriteRune(CipherAtbash(v))
	}
	return sb.String()
}
