// The AST is what allows us to start to make sense of tokens. The Abstract Syntax Tree (AST)
// is a data structure that represents the meaning and structure of the code.
// Main Job: Capture relationships and hierarchy

// The AST takes the flat sequence of tokens from the lexer and organizes them into a tree
// that represents the grammatical structure of the program — which pieces are statements, which are expressions, and how they nest. This structure is what later gets evaluated.

package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	// ! GO TRICK:
	// * We use a bytes.Buffer so that when we concatenate to a string, we don't
	// * have to create a new location in memory every time we cat. We would have to
	// * create a new location in memory everytime because...
	// ! STRINGS IN GO ARE IMMUTABLE
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ============================================================================
type ReturnStatement struct {
	Token token.Token // the 'return' token or token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ============================================================================
type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

// Function Receiver Structure
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ============================================================================
type ExpressionStatement struct {
	Token token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// ============================================================================
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal } 
func (i *Identifier) String() string { return i.Value }

// ============================================================================
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }