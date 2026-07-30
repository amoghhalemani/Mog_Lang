package main

import (
	"fmt"
	"mog_lang/interpreter"
	"mog_lang/lexer"
	"mog_lang/parser"
	"os"
	"path/filepath"
)

func readCLI() string {
	fmt.Println("Reading CLI arguments...")
	var args []string = os.Args
	if len(args) != 3 {
		return "Error: Not enough arguments provided."
	}
	if args[1] != "runs" {
		return "Error: Invalid command. Please use 'amogh runs <filename>'.mog"
	}

	var filename string = args[2]

	if filepath.Ext(filename) != ".mog" {
		return "Error: Invalid file type. Please provide a .mog file."
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return "Error: Unable to read file."
	}

	return string(content)

}

func main() {
	content := readCLI()
	tokens := lexer.New(content)
	parse := parser.New(tokens)
	program := parse.ParseProgram()
	interp := interpreter.New()
	for _, stmt := range program {
		//type switch to know which statement we're looking at before executing it.
		switch v := stmt.(type) {
		case *parser.LetStatement:
			interp.ExecuteLet(v)
		case *parser.PrintStatement:
			interp.ExecutePrint(v)
		}
	}
}
