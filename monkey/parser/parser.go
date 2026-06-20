package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer	
	errors []string

	curToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFn map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
} 

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFn[tokenType] = fn
}
