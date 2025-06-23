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
	proto := sidefunctions.MakeSpaces(string(str))
	fixed := sidefunctions.Reload(proto)
	fixed = sidefunctions.Quotes(fixed)
	fixed = sidefunctions.Punc(fixed)
	fixed = sidefunctions.ATooAn(fixed)

	err = os.WriteFile(filenames[2], []byte(sidefunctions.SliceToString(fixed)), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
