package parser

import (
	"mog_lang/lexer"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken lexer.Token
}

func New(lex *lexer.Lexer) *Parser {
	return &Parser{lexer: lex, currentToken: lex.NextToken()}
}

func (p *Parser) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

func (p *Parser) ParseLetStatement() *LetStatement {
	p.nextToken()
	name := p.currentToken.Value
	p.nextToken()
	p.nextToken()
	value := p.currentToken.Value
	p.nextToken()
	return &LetStatement{Name: name, Value: value}
}

func (p *Parser) ParsePrintStatement() *PrintStatement {
	p.nextToken()
	p.nextToken()
	name := p.currentToken.Value
	p.nextToken()
	p.nextToken()
	return &PrintStatement{Name: name}
}

func (p *Parser) ParseStatement() Statement {
	if p.currentToken.Value == "let" {
		return p.ParseLetStatement()
	}
	if p.currentToken.Value == "print" {
		return p.ParsePrintStatement()
	}
	return nil
}

func (p *Parser) ParseProgram() []Statement {
	var statements []Statement
	for p.currentToken.Type != lexer.EOF {
		present := p.ParseStatement()
		statements = append(statements, present)
	}
	return statements
}
