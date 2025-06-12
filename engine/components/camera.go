package components

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/constants"
)

type CameraComponent struct {
	FOV          float64
	Near         float64
	Far          float64
	Orthographic bool
}

func (c *CameraComponent) Type() string {
	return constants.ComponentCamera
}

func init() {
	RegisterComponent(constants.ComponentCamera, func(data map[string]interface{}) (RuntimeComponent, error) {
		fov, ok := data[constants.FieldFOV].(float64)
		if !ok {
			return nil, fmt.Errorf("camera: missing or invalid fov")
		}
		if fov <= 0 || fov >= 180 {
			return nil, fmt.Errorf("camera: FOV must be between 0 and 180 degrees")
		}

		near, ok := data[constants.FieldNear].(float64)
		if !ok {
			near = 0.1
		}

		far, ok := data[constants.FieldFar].(float64)
		if !ok {
			far = 1000.0
		}

		ortho := false
		if val, ok := data[constants.FieldOrthographic].(bool); ok {
			ortho = val
		}

		return &CameraComponent{
			FOV:          fov,
			Near:         near,
			Far:          far,
			Orthographic: ortho,
		}, nil
	})
}
