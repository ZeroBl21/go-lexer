package parser

import (
	"github.com/ZeroBl21/go-lexer/src/ast"
	"github.com/ZeroBl21/go-lexer/src/lexer"
)

type bindingPower int

const (
	defaultBP bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmtHandler func(p *parser) ast.Stmt
type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

type stmtLookup map[lexer.TokenKind]stmtHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

var bpLu = bpLookup{}
var nudLu = nudLookup{}
var ledLu = ledLookup{}
var stmtLu = stmtLookup{}

func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
	bpLu[kind] = bp
	ledLu[kind] = ledFn
}

func nud(kind lexer.TokenKind, nudFn nudHandler) {
	nudLu[kind] = nudFn
}

func stmt(kind lexer.TokenKind, stmtFn stmtHandler) {
	bpLu[kind] = defaultBP
	stmtLu[kind] = stmtFn
}

func createTokenLookups() {
	led(lexer.ASSIGNMENT, assignment, parseAssignmentExpr)
	led(lexer.PLUS_EQUALS, assignment, parseAssignmentExpr)
	led(lexer.MINUS_EQUALS, assignment, parseAssignmentExpr)
	led(lexer.STAR_EQUALS, assignment, parseAssignmentExpr)
	led(lexer.SLASH_EQUALS, assignment, parseAssignmentExpr)

	// Logical

	led(lexer.AND, logical, parseBinaryExpr)
	led(lexer.OR, logical, parseBinaryExpr)
	led(lexer.DOT_DOT, logical, parseBinaryExpr)

	// Relational

	led(lexer.LESS, relational, parseBinaryExpr)
	led(lexer.LESS_EQUALS, relational, parseBinaryExpr)
	led(lexer.GREATER, relational, parseBinaryExpr)
	led(lexer.GREATER_EQUALS, relational, parseBinaryExpr)
	led(lexer.EQUALS, relational, parseBinaryExpr)
	led(lexer.NOT_EQUALS, relational, parseBinaryExpr)

	// Additive & Multiplicative

	led(lexer.PLUS, additive, parseBinaryExpr)
	led(lexer.DASH, additive, parseBinaryExpr)
	led(lexer.STAR, multiplicative, parseBinaryExpr)
	led(lexer.SLASH, multiplicative, parseBinaryExpr)
	led(lexer.PERCENT, multiplicative, parseBinaryExpr)

	// Literals & Symbols

	nud(lexer.NUMBER, parsePrimaryExpr)
	nud(lexer.STRING, parsePrimaryExpr)
	nud(lexer.IDENTIFIER, parsePrimaryExpr)

	nud(lexer.DASH, parsePrefixExpr)

	// Statement

	stmt(lexer.CONST, parseVarDeclStmt)
	stmt(lexer.LET, parseVarDeclStmt)
}
