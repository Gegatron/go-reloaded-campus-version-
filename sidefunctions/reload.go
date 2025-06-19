package sidefunctions

import "strings"

func Reload(str string) []string {
	c := Clean(str)
	bar := []string{}
	n := 0
	for i, v := range c {
		if strings.Contains(v,"(up,") {
			bar=MultiChanges(bar,GetNumber(v),"up",i)
			c[i]=""
			continue
		}else if strings.Contains(v,"(low,") {
			bar=MultiChanges(bar,GetNumber(v),"low",i)
			c[i]=""
			continue
		}else if strings.Contains(v,"(cap,") {
			bar=MultiChanges(bar,GetNumber(v),"cap",i)
			c[i]=""
			continue
		}
		switch v {
		case "(up)":
			bar[n-1] = "up"
			c[i] = ""

		case "(low)":
			bar[n-1] = "low"
			c[i] = ""

		case "(cap)":
			bar[n-1] = "cap"
			c[i] = ""

		case "(hex)":
			bar[n-1] = "hex"
			c[i] = ""

		case "(bin)":
			bar[n-1] = "bin"
			c[i] = ""

		default:
			bar = append(bar, "nochange")
			n++

		}

	}
	str = SliceToString(c)
	c = Clean(str)
	goku := []string{}
	for u, x := range bar {
		switch x {
		case "up":
			goku = append(goku, ToUpper(c[u]))
		case "low":
			goku = append(goku, ToLower(c[u]))
		case "cap":
			goku = append(goku, Capitalize(c[u]))
		case "hex":
			goku = append(goku, ToHex(c[u]))
		case "bin":
			goku = append(goku, ToBin(c[u]))
		default:
			goku = append(goku, c[u])
		}
	}
	goku=Quotes(goku)
	return goku
}