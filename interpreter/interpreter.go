package interpreter

import (
	"fmt"
	"mog_lang/parser"
)

type Interpreter struct {
	variables map[string]string
}

func New() *Interpreter {
	return &Interpreter{variables: make(map[string]string)}
}

func (i *Interpreter) ExecuteLet(stmt *parser.LetStatement) {
	i.variables[stmt.Name] = stmt.Value
}

func (i *Interpreter) ExecutePrint(stmt *parser.PrintStatement) {
	fmt.Printf("%s\n", i.variables[stmt.Name])
}
