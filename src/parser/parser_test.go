package parser_test

import (
	"reflect"
	"testing"
	"tim/lexer"
	"tim/parser"
	"tim/token"
	"tim/tree"
)

type ExpressionCase struct {
	InputString string
	Expression  []tree.Stmt
}

func TestExpressions(t *testing.T) {
	cases := map[string]ExpressionCase{
		"equality: double equal": {
			InputString: "3 == 3",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Binary{
						Left: tree.Literal{
							Value: 3.00,
						},
						Operator: token.Token{
							Type:     token.DOUBLE_EQUAL,
							Text:     "==",
							Literal:  "==",
							Position: 2,
							Line:     1,
						},
						Right: tree.Literal{
							Value: 3.00,
						},
					},
				},
			},
		},
		"comparison: greater than": {
			InputString: "3 > 2",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Binary{
						Left: tree.Literal{
							Value: 3.00,
						},
						Operator: token.Token{
							Type:     token.GREATER,
							Text:     ">",
							Literal:  ">",
							Position: 2,
							Line:     1,
						},
						Right: tree.Literal{
							Value: 2.00,
						},
					},
				},
			},
		},
		"term: addition": {
			InputString: "2 + 4",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Binary{
						Left: tree.Literal{
							Value: 2.00,
						},
						Operator: token.Token{
							Type:     token.PLUS,
							Text:     "+",
							Literal:  "+",
							Position: 2,
							Line:     1,
						},
						Right: tree.Literal{
							Value: 4.00,
						},
					},
				},
			},
		},
		"factor: multiplication": {
			InputString: "2 * 4",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Binary{
						Left: tree.Literal{
							Value: 2.00,
						},
						Operator: token.Token{
							Type:     token.STAR,
							Text:     "*",
							Literal:  "*",
							Position: 2,
							Line:     1,
						},
						Right: tree.Literal{
							Value: 4.00,
						},
					},
				},
			},
		},
		"unary: minus": {
			InputString: "-4",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Unary{
						Operator: token.Token{
							Type:     token.MINUS,
							Text:     "-",
							Literal:  "-",
							Position: 0,
							Line:     1,
						},
						Right: tree.Literal{
							Value: 4.00,
						},
					},
				},
			},
		},
		"primary: identifier": {
			InputString: "myVariable",
			Expression: []tree.Stmt{
				tree.ExpressionStmt{
					Expr: tree.Variable{
						Name: token.Token{
							Type:     token.IDENTIFIER,
							Text:     "myVariable",
							Literal:  "myVariable",
							Position: 0,
							Line:     1,
						},
					},
				},
			},
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsedExpression := p.Parse()
			if !reflect.DeepEqual(testcase.Expression, parsedExpression) {
				t.Fatalf("expressions do not match: expected: %+v, actual: %+v", testcase.Expression, parsedExpression)
			}
		})
	}
}
