package tree

import "tim/token"

func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

type Environment struct {
	values map[string]interface{}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(token token.Token) interface{} {
	if _, ok := e.values[token.Text]; ok {
		return e.values[token.Text]
	}

	panic(NewRuntimeError("Undefined variable '" + token.Text + "'."))
}
