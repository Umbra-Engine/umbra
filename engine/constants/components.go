package constants

// ComponentType holds known component type strings
const (
	ComponentSpriteRenderer = "SpriteRenderer"
	ComponentCamera         = "Camera"
	ComponentTransform      = "Transform"
	ComponentScript         = "Script"
	ComponentRigidBody      = "RigidBody"
)

// Common Component Field Keys
const (
	// Shared fields
	FieldActive = "active"
	FieldLayer  = "layer"
	FieldMask   = "mask"

	// Sprite Renderer fields
	FieldTexture = "texture"
	FieldFOV     = "fov"

	// Transform fields
	FieldPosition = "position"
	FieldScale    = "scale"
	FieldRotation = "rotation"

	// Camera fields
	FieldNear         = "near"
	FieldFar          = "far"
	FieldOrthographic = "orthographic"

	// Script fields
	FieldEnabled   = "enabled"
	FieldClassName = "classname"
	FieldParams    = "params"

	// Rigid Body fields
	FieldMass            = "mass"
	FieldVelocity        = "velocity"
	FieldAngularVelocity = "angular_velocity"
	FieldUseGravity      = "use_gravity"
	FieldIsKinematic     = "is_kinematic"
	FieldLinearDamp      = "linear_damp"
	FieldAngularDamp     = "angular_damp"
	FieldForce           = "force"
	FieldTorque          = "torque"
	FieldFreezePosition  = "freeze_position"
	FieldFreezeRotation  = "freeze_rotation"
)
