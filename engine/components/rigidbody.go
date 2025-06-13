package components

import (
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type RigidBodyComponent struct {
	Mass            float64
	Velocity        mathx.Vector3
	AngularVelocity mathx.Vector3
	UseGravity      bool
	IsKinematic     bool

	LinearDamp  float64
	AngularDamp float64

	Force  mathx.Vector3
	Torque mathx.Vector3

	Layer uint32
	Mask  uint32

	FreezePosition [3]bool // [X, Y, Z]
	FreezeRotation [3]bool // [X, Y, Z]
}

func (r *RigidBodyComponent) Type() string {
	return constants.ComponentRigidBody
}

// ApplyForce accumulates a force to be processed in the physics update step
func (r *RigidBodyComponent) ApplyForce(force mathx.Vector3) {
	r.Force = r.Force.Add(force)
}

// ApplyTorque accumulates torque to be processed in the physics update step
func (r *RigidBodyComponent) ApplyTorque(torque mathx.Vector3) {
	r.Torque = r.Torque.Add(torque)
}

func init() {
	RegisterComponent(constants.ComponentRigidBody, func(data map[string]interface{}) (RuntimeComponent, error) {
		mass, _ := data[constants.FieldMass].(float64)
		if mass <= 0 {
			mass = 1.0
		}

		gravity, _ := data[constants.FieldUseGravity].(bool)
		isKinematic, _ := data[constants.FieldIsKinematic].(bool)

		linearDamp, _ := data[constants.FieldLinearDamp].(float64)
		angularDamp, _ := data[constants.FieldAngularDamp].(float64)

		layer := uint32(0)
		if val, ok := data[constants.FieldLayer].(int); ok {
			layer = uint32(val)
		}
		mask := uint32(0xFFFFFFFF)
		if val, ok := data[constants.FieldMask].(int); ok {
			mask = uint32(val)
		}

		var freezePos [3]bool
		if freeze, ok := data[constants.FieldFreezePosition].([]interface{}); ok && len(freeze) == 3 {
			for i := 0; i < 3; i++ {
				if b, ok := freeze[i].(bool); ok {
					freezePos[i] = b
				}
			}
		}

		var freezeRot [3]bool
		if freeze, ok := data[constants.FieldFreezeRotation].([]interface{}); ok && len(freeze) == 3 {
			for i := 0; i < 3; i++ {
				if b, ok := freeze[i].(bool); ok {
					freezeRot[i] = b
				}
			}
		}

		velocity := mathx.Vector3{}
		if raw, ok := data[constants.FieldVelocity].(map[string]any); ok {
			velocity, _ = mathx.FromMap3(raw)
		}

		angularVelocity := mathx.Vector3{}
		if raw, ok := data[constants.FieldAngularVelocity].(map[string]any); ok {
			angularVelocity, _ = mathx.FromMap3(raw)
		}

		return &RigidBodyComponent{
			Mass:            mass,
			UseGravity:      gravity,
			IsKinematic:     isKinematic,
			LinearDamp:      linearDamp,
			AngularDamp:     angularDamp,
			Velocity:        velocity,
			AngularVelocity: angularVelocity,
			Layer:           layer,
			Mask:            mask,
			FreezePosition:  freezePos,
			FreezeRotation:  freezeRot,
		}, nil
	})
}
