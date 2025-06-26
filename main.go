package main

import (
	"fmt"
	"os"
	"path/filepath"

	"sidefunctions/sidefunctions"
)

func main() {
	filenames := os.Args
	if len(filenames) != 3 {
		fmt.Println("invalid input")
		return
	}

	for i := 1; i < 3; i++ {
		if filepath.Ext(filenames[i]) != ".txt" {
			fmt.Println("invalid input")
			return
		}
	}

	// if filenames[1][len(filenames[1])-4:len(filenames[1])] != ".txt" || filenames[2][len(filenames[2])-4:len(filenames[2])] != ".txt" {
	// 	fmt.Println("invalid input")
	// 	return
	// }

	str, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fixed := sidefunctions.Trait(string(str))
	err = os.WriteFile(filenames[2], []byte(fixed), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
