package parser

type LetStatement struct {
	Name  string
	Value string
}

type PrintStatement struct {
	Name string
}

// statement Interface

type Statement interface {
	statementNode()
}

func (ls *LetStatement) statementNode()   {}
func (ps *PrintStatement) statementNode() {}
