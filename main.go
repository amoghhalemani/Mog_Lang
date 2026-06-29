package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readCLI() {
	fmt.Println("Reading CLI arguments...")
	var args []string = os.Args
	if len(args) != 3 {
		fmt.Println("Error: Not enough arguments provided.")
		return
	}
	if args[1] != "runs" {
		fmt.Println("Error: Invalid command. Please use 'amogh runs <filename>'.mog")
		return
	}

	var filename string = args[2]

	if filepath.Ext(filename) != ".mog" {
		fmt.Println("Error: Invalid file type. Please provide a .mog file.")
		return
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: Unable to read file.")
		return
	}

	fmt.Println(string(content))

}

func main() {
	readCLI()
	fmt.Println("Hello, World!")
}
