package sidefunctions

import "strings"

func Reload(str string) []string {
	c := Clean(str)
	bar := []string{}
	n := 0
	var temp []string
	cou := 0
	for _, car := range c {
		if cou == 0 && (car == "(up)" || car == "(low)" || car == "(cap)" || car == "(bin)" || car == "(hex)" || strings.Contains(car, "(up,") || strings.Contains(car, "(low,") || strings.Contains(car, "(cap,")) {
			continue
		}
		temp = append(temp, car)
		cou++

	}
	c = temp
	s := ""
	for i:=0;i<len(c);i++ {
		if strings.Contains(c[i], "(up,") {
			bar = MultiChanges(bar, GetNumber(c[i]), "up", i)
			c[i] = ""
			continue
		} else if strings.Contains(c[i], "(low,") {
			bar = MultiChanges(bar, GetNumber(c[i]), "low", i)
			c[i] = ""
			continue
		} else if strings.Contains(c[i], "(cap,") {
			bar = MultiChanges(bar, GetNumber(c[i]), "cap", i)
			c[i] = ""
			continue
		}
		switch c[i] {
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
			
			c[i-1] = ToHex(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--

		case "(bin)":

			c[i-1] = ToBin(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--
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
		default:
			goku = append(goku, c[u])
		}
	}
	goku = Quotes(goku)
	return goku
}
