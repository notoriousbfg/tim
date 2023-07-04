package parser_test

import (
	"reflect"
	"testing"
	"tim/lexer"
	"tim/parser"
	"tim/token"
	"tim/tree"
)

type StatementCase struct {
	InputString string
	Statements  []tree.Stmt
}

func TestExpressions(t *testing.T) {
	cases := map[string]StatementCase{
		"equality: double equal": {
			InputString: "3 == 3",
			Statements: []tree.Stmt{
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
			Statements: []tree.Stmt{
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
			Statements: []tree.Stmt{
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
			Statements: []tree.Stmt{
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
			Statements: []tree.Stmt{
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
			Statements: []tree.Stmt{
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
			if !reflect.DeepEqual(testcase.Statements, parsedExpression) {
				t.Fatalf("expressions do not match: expected: %+v, actual: %+v", testcase.Statements, parsedExpression)
			}
		})
	}
}

func TestStatements(t *testing.T) {
	cases := map[string]StatementCase{
		"list": {
			InputString: "(1, 2, 3)",
			Statements: []tree.Stmt{
				tree.ListStmt{
					Items: []tree.Stmt{
						tree.ExpressionStmt{
							Expr: tree.Literal{
								Value: 1.00,
							},
						},
						tree.ExpressionStmt{
							Expr: tree.Literal{
								Value: 2.00,
							},
						},
						tree.ExpressionStmt{
							Expr: tree.Literal{
								Value: 3.00,
							},
						},
					},
				},
			},
		},
		"variable declaration": {
			InputString: "(myVariable: \"testvariable\")",
			Statements: []tree.Stmt{
				tree.ListStmt{
					Items: []tree.Stmt{
						tree.VariableStmt{
							Name: token.Token{
								Type:     token.IDENTIFIER,
								Text:     "myVariable",
								Literal:  "myVariable",
								Position: 1,
								Line:     1,
							},
							Initializer: tree.ExpressionStmt{
								Expr: tree.Literal{
									Value: "testvariable",
								},
							},
						},
					},
				},
			},
		},
		"list function": {
			InputString: "(\"hello\").print()",
			Statements: []tree.Stmt{
				tree.ListStmt{
					Items: []tree.Stmt{
						tree.ExpressionStmt{
							Expr: tree.Literal{
								Value: "hello",
							},
						},
					},
					Functions: []tree.Stmt{
						tree.CallStmt{
							Callee: tree.Variable{
								Name: token.Token{
									Type:     token.IDENTIFIER,
									Text:     "print",
									Literal:  "print",
									Position: 10,
									Line:     1,
								},
							},
							ClosingParen: token.Token{
								Type:     token.RIGHT_PAREN,
								Text:     ")",
								Literal:  ")",
								Position: 16,
								Line:     1,
							},
						},
					},
				},
			},
		},
		"nested list function": {
			InputString: "((\"hello\", \"tim\").join(\",\")).print()",
			Statements: []tree.Stmt{
				tree.ListStmt{
					Items: []tree.Stmt{
						tree.ListStmt{
							Items: []tree.Stmt{
								tree.ExpressionStmt{
									Expr: tree.Literal{
										Value: "hello",
									},
								},
								tree.ExpressionStmt{
									Expr: tree.Literal{
										Value: "tim",
									},
								},
							},
							Functions: []tree.Stmt{
								tree.CallStmt{
									Callee: tree.Variable{
										Name: token.Token{
											Type:     token.IDENTIFIER,
											Text:     "join",
											Literal:  "join",
											Position: 18,
											Line:     1,
										},
									},
									ClosingParen: token.Token{
										Type:     token.RIGHT_PAREN,
										Text:     ")",
										Literal:  ")",
										Position: 26,
										Line:     1,
									},
									Arguments: []tree.Expr{
										tree.Literal{
											Value: ",",
										},
									},
								},
							},
						},
					},
					Functions: []tree.Stmt{
						tree.CallStmt{
							Callee: tree.Variable{
								Name: token.Token{
									Type:     token.IDENTIFIER,
									Text:     "print",
									Literal:  "print",
									Position: 29,
									Line:     1,
								},
							},
							ClosingParen: token.Token{
								Type:     token.RIGHT_PAREN,
								Text:     ")",
								Literal:  ")",
								Position: 35,
								Line:     1,
							},
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
			if !reflect.DeepEqual(testcase.Statements, parsedExpression) {
				t.Fatalf("expressions do not match: expected: %+v, actual: %+v", testcase.Statements, parsedExpression)
			}
		})
	}
}
