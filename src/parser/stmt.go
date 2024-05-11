package parser

import (
	"github.com/ZeroBl21/go-lexer/src/ast"
	"github.com/ZeroBl21/go-lexer/src/lexer"
)

func parseStmt(p *parser) ast.Stmt {
	if stmtFn, exists := stmtLu[p.currentTokenKind()]; exists {
		return stmtFn(p)
	}

	expr := parseExpr(p, defaultBP)
	p.expect(lexer.SEMICOLON)

	return ast.ExpressionStmt{
		Expression: expr,
	}
}

func parseVarDeclStmt(p *parser) ast.Stmt {
	IsConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER,
		"Inside variable declaration expected to find variable name").Value
	p.expect(lexer.ASSIGNMENT)
	assignedValue := parseExpr(p, assignment)
	p.expect(lexer.SEMICOLON)

	return ast.VarDeclStmt{
		IsConstant: IsConstant,
		Identifier: varName,
		Value:      assignedValue,
	}
}
