package main

import (
	"fmt"
	"os"
	"strings"
)

func print_prompt() {
	fmt.Print("db>")
}

func main() {
	for {
		var input string
		print_prompt()
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input")
			os.Exit(-1)
		}

		if strings.Compare(input, ".exit") == 0 {
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command %s\n", input)
		}
	}
}
