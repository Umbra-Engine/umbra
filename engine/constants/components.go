package constants

// ComponentType holds known component type strings
const (
	ComponentSpriteRenderer = "SpriteRenderer"
	ComponentCamera         = "Camera"
	ComponentTransform      = "Transform"
	ComponentScript         = "Script"
)

// Common Component Field Keys
const (
	FieldActive = "active"

	// SpriteRenderer fields
	FieldTexture = "texture"
	FieldLayer   = "layer"
	FieldFOV     = "fov"

	// Transform fields
	FieldPosition = "position"
	FieldScale    = "scale"
	FieldRotation = "rotation"

	// Camera fields
	FieldNear         = "near"
	FieldFar          = "far"
	FieldOrthographic = "orthographic"
)
