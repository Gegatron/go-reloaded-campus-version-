package main

import (
	"fmt"
	"os"
	"path/filepath"
	
	"goreloaded/goreloaded"
)

func main() {
	filenames := os.Args
	if len(filenames) != 3 {
		fmt.Println("invalid input")
		return
	}
	if filenames[1] == filenames[2] {
		fmt.Println("invalid input")
		return
	}
	for i := 1; i < 3; i++ {
		if filepath.Ext(filenames[i]) != ".txt" {
			fmt.Println("invalid input")
			return
		}
	}
	str, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	reloaded:= goreloaded.Trait(string(str))
	temp:=""
	for temp!=reloaded {
		temp=reloaded
		reloaded=goreloaded.Trait(string(reloaded))
		
	}
	err = os.WriteFile(filenames[2], []byte(reloaded), 0o644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
