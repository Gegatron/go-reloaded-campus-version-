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
	
	geted := sidefunctions.Clean(string(str))
	s:=string(str)
	
	fmt.Println(s)
	for i := 0; i < len(geted); i++ {
		fmt.Println(geted[i])
	}
	
	err = os.WriteFile(filenames[2], []byte(sidefunctions.SliceToString(geted)), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
