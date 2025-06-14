package constants

// ComponentType holds known component type strings
const (
	ComponentSpriteRenderer = "SpriteRenderer"
	ComponentCamera         = "Camera"
	ComponentTransform      = "Transform"
	ComponentScript         = "Script"
	ComponentRigidBody2D    = "RigidBody2D"
	ComponentCollider2D     = "Collider2D"
)

// Common Component Field Keys
const (
	// Shared fields
	FieldActive = "active"
	FieldLayer  = "layer"
	FieldMask   = "mask"
	FieldFOV    = "fov"
	FieldOffset = "offset"

	// Sprite Renderer fields
	FieldTexture = "texture"

	// Transform fields
	FieldPosition = "position"
	FieldScale    = "scale"
	FieldRotation = "rotation"

	// Camera fields
	FieldIsOrthographic   = "is_orthographic"
	FieldOrthographicSize = "orthographic_size"
	FieldNear             = "near"
	FieldFar              = "far"
	FieldClearColor       = "clear_color"
	FieldPriority         = "priority"
	FieldClearFlags       = "clear_flags"
	FieldViewportRect     = "viewport_rect"
	FieldZoom             = "zoom"
	FieldCullingMask      = "culling_mask"
	FieldFollowTargetID   = "follow_target_id"

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

	// Collider Fields
	FieldType      = "type"
	FieldSize      = "size"
	FieldRadius    = "radius"
	FieldPoints    = "points"
	FieldIsTrigger = "is_trigger"
)
