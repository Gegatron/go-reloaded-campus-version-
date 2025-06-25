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
	fixed:=sidefunctions.Trait(string(str))
	err = os.WriteFile(filenames[2], []byte(fixed), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	
}
