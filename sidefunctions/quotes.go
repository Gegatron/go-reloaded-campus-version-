package sidefunctions

func Quotes(reloaded []string) []string {
	b := false

	for i := 0; i < len(reloaded); i++ {
		if reloaded[i] == "'" && !b && i != len(reloaded)-1 {
			if reloaded[i+1]!="'" {
				reloaded[i+1] = "'" + reloaded[i+1]
			reloaded[i] = ""

			b = true
			}
			
		} else if reloaded[i] == "'" && b {
			if reloaded[i-1]!="'" {
			reloaded[i-1] = reloaded[i-1] + "'"
			reloaded[i] = ""

			b = false
		}
		}
	}

	return reloaded
}
