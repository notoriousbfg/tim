package tree

type Environment struct {
	values map[string]interface{}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}
