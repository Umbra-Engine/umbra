package components

import (
	"fmt"

	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type TransformComponent struct {
	Position mathx.Vector3
	Rotation mathx.Vector3
	Scale    mathx.Vector3
}

func (t *TransformComponent) Type() string {
	return constants.ComponentTransform
}

func init() {
	RegisterComponent(constants.ComponentTransform, func(data map[string]interface{}) (RuntimeComponent, error) {
		// Fetch position
		posRaw, ok := data[constants.FieldPosition].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("missing or invalid position field")
		}

		pos, err := mathx.FromMap3(posRaw)
		if err != nil {
			return nil, err
		}

		// Fetch rotation
		rotRaw, ok := data[constants.FieldRotation].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("missing or invalid rotation field")
		}

		rot, err := mathx.FromMap3(rotRaw)
		if err != nil {
			return nil, err
		}

		// Fetch scale
		scaleRaw, ok := data[constants.FieldScale].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("missing or invalid scale field")
		}

		scale, err := mathx.FromMap3(scaleRaw)
		if err != nil {
			return nil, err
		}

		return &TransformComponent{
			Position: pos,
			Rotation: rot,
			Scale:    scale,
		}, nil
	})
}
