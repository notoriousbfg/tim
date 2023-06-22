package tree

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) interface{}
	VisitPrintStmt(stmt PrintStmt) interface{}
}

type ExpressionStmt struct {
	Expr Expr
}

func (es ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(es)
}

type PrintStmt struct {
	Expr Expr
}

func (ps *PrintStmt) Print(expr Expr) {
	ps.Expr = expr
}

func (ps PrintStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmt(ps)
}
