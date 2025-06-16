package components

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/scripting"
)

// ScriptComponent attaches a script by name to a game object.
// The script will be looked up and executed by the scripting runtime.
type ScriptComponent struct {
	ClassName string
	Enabled   bool
	Params    map[string]any
	Instance  scripting.Script
}

func (s *ScriptComponent) Type() string { return constants.ComponentScript }

func init() {
	RegisterComponent(constants.ComponentScript, func(data map[string]interface{}) (RuntimeComponent, error) {
		className, ok := data[constants.FieldClassName].(string)
		if !ok || className == "" {
			return nil, fmt.Errorf("script component: missing or invalid class name")
		}

		enabled := true
		if val, ok := data[constants.FieldEnabled].(bool); ok {
			enabled = val
		}

		// Everything else is considered user-defined params
		params := make(map[string]any)
		for key, val := range data {
			if key != constants.FieldClassName && key != constants.FieldEnabled {
				params[key] = val
			}
		}

		instance, err := scripting.CreateScript(className, params)
		if err != nil {
			return nil, err
		}

		return &ScriptComponent{
			ClassName: className,
			Enabled:   enabled,
			Params:    params,
			Instance:  instance,
		}, nil
	})
}

func (s *ScriptComponent) Update(dt float64) {
	if s.Enabled && s.Instance != nil {
		s.Instance.OnUpdate(dt)
	}
}

func (s *ScriptComponent) OnStart() {
	if s.Enabled && s.Instance != nil {
		s.Instance.OnStart()
	}
}

func (s *ScriptComponent) OnStop() {
	if s.Enabled && s.Instance != nil {
		s.Instance.OnStop()
	}
}
