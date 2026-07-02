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
