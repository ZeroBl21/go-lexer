package parser

import (
	"fmt"

	"github.com/ZeroBl21/go-lexer/src/ast"
	"github.com/ZeroBl21/go-lexer/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func newParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)

	p := newParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parseStmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}

// Helpers

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++

	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	kind := p.currentTokenKind()

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s but received %s instead\n",
				lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
