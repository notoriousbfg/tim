package tree

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

func (e *Environment) Get(name string) interface{} {
	return e.values[name]
}
