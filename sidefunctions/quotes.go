package sidefunctions

import "strings"

func Quotes(reloaded []string) []string {
	b := false

	for i := 0; i < len(reloaded); i++ {
		if i != len(reloaded)-1 && !b && reloaded[i][len(reloaded[i])-1] == '\'' {
			if reloaded[i+1][0] != '\'' {
				reloaded[i+1] = "'" + reloaded[i+1]
				reloaded[i] = reloaded[i][:len(reloaded[i])-1]
				reloaded = Clean(strings.Join(reloaded," "))
				i--
				b = true
			}
		} else if i != 0 && b && reloaded[i][0] == '\'' {
			if reloaded[i-1][len(reloaded[i-1])-1] != '\'' {
				reloaded[i-1] = reloaded[i-1] + "'"
				reloaded[i] = reloaded[i][1:]
				reloaded = Clean(strings.Join(reloaded," "))
				i--
				b = false
			}
		}
	}

	return reloaded
}
