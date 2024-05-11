package main

import (
	"os"

	"github.com/ZeroBl21/go-lexer/src/lexer"
	"github.com/ZeroBl21/go-lexer/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/04.lang")
	tokens := lexer.Tokenize(string(bytes))

	// for _, token := range tokens {
	// 	token.Debug()
	// }

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
