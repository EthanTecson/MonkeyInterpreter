// The AST is what allows us to start to make sense of tokens. The Abstract Syntax Tree (AST)
// is a data structure that represents the meaning and structure of the code.
// Main Job: Capture relationships and hierarchy

// The AST takes the flat sequence of tokens from the lexer and organizes them into a tree
// that represents the grammatical structure of the program — which pieces are statements, which are expressions, and how they nest. This structure is what later gets evaluated.

package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // The token.LET
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
