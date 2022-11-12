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
	Expression  tree.Expr
}

func TestParser(t *testing.T) {
	cases := map[string]ExpressionCase{
		"basic addition": {
			InputString: "2 + 4",
			Expression: tree.Binary{
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
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsedExpression := p.Parse()
			if !reflect.DeepEqual(testcase.Expression, parsedExpression) {
				t.Fatal("expressions do not match", tree.Print(testcase.Expression), tree.Print(parsedExpression))
			}
		})
	}
}
