package scene

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type Scene struct {
	Name    string       `yaml:"name"`
	Objects []GameObject `yaml:"objects"`
}

type GameObject struct {
	Name       string        `yaml:"name"`
	Position   mathx.Vector3 `yaml:"position"`
	Rotation   mathx.Vector3 `yaml:"rotation"`
	Scale      mathx.Vector3 `yaml:"scale"`
	Components []Component   `yaml:"components"`
	Children   []GameObject  `yaml:"children"`
}

type Component struct {
	Type string         `yaml:"type"`
	Data map[string]any `yaml:",inline"`
}

func (g *GameObject) InitDefaults() {
	// Default Scale
	if g.Scale == (mathx.Vector3{}) {
		g.Scale = mathx.Vector3{X: 1, Y: 1, Z: 1}
	}

	// Default rotation to zero explicitly
	if g.Rotation == (mathx.Vector3{}) {
		g.Rotation = mathx.Vector3{X: 0, Y: 0, Z: 0}
	}

	// Default position
	if g.Position == (mathx.Vector3{}) {
		g.Position = mathx.Vector3{X: 0, Y: 0, Z: 0}
	}

	// Recursively default children
	for i := range g.Children {
		g.Children[i].InitDefaults()
	}
}

func (s *Scene) Validate() []error {
	var errs []error
	nameSet := make(map[string]struct{})

	for _, obj := range s.Objects {
		errs = append(errs, validateObjectRecursive(obj, nameSet)...)
	}

	if s.Name == "" {
		errs = append(errs, fmt.Errorf("scene name is missing"))
	}

	return errs
}

func validateObjectRecursive(obj GameObject, nameSet map[string]struct{}) []error {
	var errs []error

	if obj.Name == "" {
		errs = append(errs, fmt.Errorf("game object at position %+v is missing a name", obj.Position))
	}

	if !isNameUnique(obj.Name, nameSet) {
		errs = append(errs, fmt.Errorf("%s is not a unique name for game object", obj.Name))
	}

	if len(obj.Children) > 0 {
		for _, child := range obj.Children {
			errs = append(errs, validateObjectRecursive(child, nameSet)...)
		}
	}

	return errs
}

func isNameUnique(name string, nameSet map[string]struct{}) bool {
	if _, exists := nameSet[name]; !exists {
		nameSet[name] = struct{}{}
		return true
	}

	return false
}
