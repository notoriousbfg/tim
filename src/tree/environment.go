package tree

import (
	"tim/token"
)

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		values:    make(map[string]interface{}),
		Enclosing: enclosing,
	}
}

type Environment struct {
	values    map[string]interface{}
	Enclosing *Environment
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value

	// only modify global if exists
	if _, ok := e.Enclosing.values[name]; ok {
		e.Enclosing.values[name] = value
	}
}

func (e *Environment) Get(token token.Token) interface{} {
	if _, ok := e.values[token.Text]; ok {
		return e.values[token.Text]
	}

	if e.Enclosing != nil {
		if _, ok := e.Enclosing.values[token.Text]; ok {
			return e.Enclosing.values[token.Text]
		}
	}

	panic(NewRuntimeError("Undefined variable '" + token.Text + "'."))
}
