package constants

// ComponentType holds known component type strings
const (
	ComponentSpriteRenderer = "SpriteRenderer"
	ComponentCamera         = "Camera"
	ComponentTransform      = "Transform"
	ComponentScript         = "Script"
	ComponentRigidBody2D    = "RigidBody2D"
	ComponentCollider2D     = "Collider2D"
	ComponentLight          = "Light"
	ComponentAudioSource    = "AudioSource"
	ComponentTextRenderer   = "TextRenderer"
	ComponentUIComponent    = "UIComponent"
)

// Common Component Field Keys
const (
	// Shared fields
	FieldActive   = "active"
	FieldLayer    = "layer"
	FieldMask     = "mask"
	FieldFOV      = "fov"
	FieldOffset   = "offset"
	FieldPriority = "priority"
	FieldColor    = "color"

	// Sprite Renderer fields
	FieldTexture = "texture"
	FieldFlipX   = "flip_x"
	FieldFlipY   = "flip_y"

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

	// Light Fields
	FieldIntensity   = "intensity"
	FieldRange       = "range"
	FieldAngle       = "angle"
	FieldCastShadows = "cast_shadows"
	FieldFalloff     = "falloff"

	// Audio Source Fields
	FieldClip        = "clip"
	FieldVolume      = "volume"
	FieldPitch       = "pitch"
	FieldLoop        = "loop"
	FieldPlayOnStart = "play_on_start"
	FieldIsMuted     = "is_muted"
	FieldSpatial     = "spatial"
	FieldMinDistance = "min_distance"
	FieldMaxDistance = "max_distance"
	FieldRolloff     = "rolloff"
	FieldOutputGroup = "output_group"

	// Text Renderer Fields
	FieldText                = "text"
	FieldFont                = "font"
	FieldFontSize            = "font_size"
	FieldHorizontalAlignment = "horizontal_alignment"
	FieldVerticalAlignment   = "vertical_alignment"
	FieldLineSpacing         = "line_spacing"
	FieldMaxWidth            = "max_width"
	FieldWrap                = "wrap"
	FieldZIndex              = "z_index"
	FieldBillboard           = "billboard"

	// UI Component Field
	FieldAnchor        = "anchor"
	FieldPivot         = "pivot"
	FieldVisible       = "visible"
	FieldInteractable  = "interactable"
	FieldRaycastTarget = "raycast_target"
	FieldClipChildren  = "clip_children"
	FieldAutoLayout    = "auto_layout"
	FieldPadding       = "padding"
)
