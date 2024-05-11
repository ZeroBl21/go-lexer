package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (e ExpressionStmt) stmt() {}

type VarDeclStmt struct {
	Identifier string
	IsConstant bool
	Value Expr
	// TODO
	// ExplicitType Type
}

func (e VarDeclStmt) stmt() {}
