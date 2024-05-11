package ast

import "github.com/ZeroBl21/go-lexer/src/lexer"

// ---
// Literal Expressions
// ---

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (s StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (s SymbolExpr) expr() {}

// ---
// Complex Expressions
// ---

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (b BinaryExpr) expr() {}

type PrefixExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (b PrefixExpr) expr() {}

type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Value    Expr
}

func (b AssignmentExpr) expr() {}
