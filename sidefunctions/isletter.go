package sidefunctions

import "unicode"

func IsLetter(s rune) bool {
	return unicode.Is(unicode.Latin, (s)) && unicode.IsLetter(s)
}
