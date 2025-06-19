package sidefunctions

func MultiChanges(real []string, num int, str string ,index int) []string {
	if num<=0 {
		return real
	}
	if len(real)<num {
		num=len(real)
	}
	ind:=len(real)-num
	for i:=range real {
		if ind<=i && ind<index {
			real[i]=str
		}
	}
	return real
}
