package main

import (
	"fmt"
	"os"

	"sidefunctions/sidefunctions"
)

func main() {
	filenames := os.Args
	if len(filenames) != 3 {
		fmt.Println("invalid input")
		return
	}
	str, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("error", err)
		return
	}
	new:=sidefunctions.MakeSpaces(string(str))
	new2:=sidefunctions.Reload(new)
new2=sidefunctions.Quotes(new2)
new2=sidefunctions.Punc(new2)
new2=sidefunctions.ATooAn(new2)
	

	
	err = os.WriteFile(filenames[2], []byte(sidefunctions.SliceToString(new2)), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
