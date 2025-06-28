package goreloaded

import "strings"

func Quotes(reloaded []string) []string {
	str := strings.Join(reloaded, " ")

	index := 0
	new := ""
	b := true
	for i := 0; i < len(str); i++ {
		if b && str[i] == '\'' {
			if i == len(str)-1 {
				new += string(str[i])
			}
			new += " "
			index = i
			b = false

			continue

		}
		if !b && str[i] == '\'' {

			new += "'" + strings.TrimSpace(str[index+1:i]) + "'" + " "
			b = true
			continue
		}
		if b {
			new += string(str[i])
		}
		if !b && i == len(str)-1 {
			new += str[index:]
		}
	}
	reloaded = Clean(new)
	return reloaded
}
