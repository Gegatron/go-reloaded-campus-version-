package sidefunctions

import "strings"

func BeginsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}