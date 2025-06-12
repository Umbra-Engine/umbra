package components

import (
	"fmt"
)

type RuntimeComponent interface {
	Type() string
}

type ComponentFactory func(data map[string]interface{}) (RuntimeComponent, error)

var registry = map[string]ComponentFactory{}

func RegisterComponent(name string, factory ComponentFactory) {
	registry[name] = factory
}

func BuildRuntimeComponent(compType string, data map[string]interface{}) (RuntimeComponent, error) {
	if factory, ok := registry[compType]; ok {
		return factory(data)
	}
	return nil, fmt.Errorf("unrecognized component type: %s", compType)
}
