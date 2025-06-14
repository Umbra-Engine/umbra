package components

import (
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type RigidBody2DComponent struct {
	Mass            float64
	Velocity        mathx.Vector2
	AngularVelocity float64
	UseGravity      bool
	IsKinematic     bool

	LinearDamp  float64
	AngularDamp float64

	Force  mathx.Vector2
	Torque float64

	Layer uint32
	Mask  uint32

	FreezePosition [2]bool // [X, Y]
	FreezeRotation bool
}

func (r *RigidBody2DComponent) Type() string {
	return constants.ComponentRigidBody2D
}

func (r *RigidBody2DComponent) ApplyForce(force mathx.Vector2) {
	r.Force = r.Force.Add(force)
}

func (r *RigidBody2DComponent) ApplyTorque(torque float64) {
	r.Torque += torque
}

func init() {
	RegisterComponent(constants.ComponentRigidBody2D, func(data map[string]interface{}) (RuntimeComponent, error) {
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

		var freezePos [2]bool
		if freeze, ok := data[constants.FieldFreezePosition].([]interface{}); ok && len(freeze) == 2 {
			for i := range 2 {
				if b, ok := freeze[i].(bool); ok {
					freezePos[i] = b
				}
			}
		}

		freezeRot := false
		if val, ok := data[constants.FieldFreezeRotation].(bool); ok {
			freezeRot = val
		}

		velocity := mathx.Vector2{}
		if raw, ok := data[constants.FieldVelocity].(map[string]any); ok {
			velocity, _ = mathx.FromMap2(raw)
		}

		angularVelocity := 0.0
		if val, ok := data[constants.FieldAngularVelocity].(float64); ok {
			angularVelocity = val
		}

		return &RigidBody2DComponent{
			Mass:            mass,
			Velocity:        velocity,
			AngularVelocity: angularVelocity,
			UseGravity:      gravity,
			IsKinematic:     isKinematic,
			LinearDamp:      linearDamp,
			AngularDamp:     angularDamp,
			Layer:           layer,
			Mask:            mask,
			FreezePosition:  freezePos,
			FreezeRotation:  freezeRot,
		}, nil
	})
}
