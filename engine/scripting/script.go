package scripting

import "fmt"

type Script interface {
	OnStart()
	OnUpdate(deltaTime float64)
	OnStop()
}

type Factory func(params map[string]interface{}) Script

var scriptRegistry = make(map[string]Factory)

func RegisterScript(name string, factory Factory) {
	scriptRegistry[name] = factory
}

func CreateScript(name string, params map[string]interface{}) (Script, error) {
	if factory, ok := scriptRegistry[name]; ok {
		return factory(params), nil
	}
	return nil, fmt.Errorf("script not found: %s", name)
}
