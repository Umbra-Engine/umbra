package components

import (
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type CameraComponent struct {
	IsOrthographic   bool
	FieldOfView      float64
	OrthographicSize float64
	NearClip         float64
	FarClip          float64
	ClearColor       mathx.Vector4
	Priority         int
	ViewportRect     mathx.Vector4
	ClearFlags       string
	Zoom             float64
	CullingMask      uint32

	// Reference by ID instead of pointer
	FollowTargetID string
	Offset         mathx.Vector3
}

func (c *CameraComponent) Type() string {
	return constants.ComponentCamera
}

func init() {
	RegisterComponent(constants.ComponentCamera, func(data map[string]interface{}) (RuntimeComponent, error) {
		isOrtho, _ := data[constants.FieldIsOrthographic].(bool)

		fov, _ := data[constants.FieldFOV].(float64)
		if fov <= 0 {
			fov = 60
		}

		orthoSize, _ := data[constants.FieldOrthographicSize].(float64)
		if orthoSize <= 0 {
			orthoSize = 5
		}

		near, _ := data[constants.FieldNear].(float64)
		if near == 0 {
			near = 0.1
		}
		far, _ := data[constants.FieldFar].(float64)
		if far == 0 {
			far = 1000
		}

		clearColor := mathx.Vector4{X: 0, Y: 0, Z: 0, W: 1}
		if raw, ok := data[constants.FieldClearColor].(map[string]any); ok {
			if parsed, err := mathx.FromMap4(raw); err == nil {
				clearColor = parsed
			}
		}

		priority := 0
		if val, ok := data[constants.FieldPriority].(int); ok {
			priority = val
		}

		zoom, _ := data[constants.FieldZoom].(float64)
		if zoom <= 0 {
			zoom = 1
		}

		cullingMask := uint32(0xFFFFFFFF)
		if val, ok := data[constants.FieldCullingMask].(int); ok {
			cullingMask = uint32(val)
		}

		followID := ""
		if val, ok := data[constants.FieldFollowTargetID].(string); ok {
			followID = val
		}

		offset := mathx.ZeroVector3()
		if raw, ok := data[constants.FieldOffset].(map[string]any); ok {
			offset, _ = mathx.FromMap3(raw)
		}

		viewport := mathx.Vector4{X: 0, Y: 0, Z: 1, W: 1}
		if raw, ok := data[constants.FieldViewportRect].(map[string]any); ok {
			viewport, _ = mathx.FromMap4(raw)
		}

		clearFlags := "color"
		if val, ok := data[constants.FieldClearFlags].(string); ok {
			clearFlags = val
		}

		return &CameraComponent{
			IsOrthographic:   isOrtho,
			FieldOfView:      fov,
			OrthographicSize: orthoSize,
			NearClip:         near,
			FarClip:          far,
			ClearColor:       clearColor,
			Priority:         priority,
			Zoom:             zoom,
			CullingMask:      cullingMask,
			FollowTargetID:   followID,
			Offset:           offset,
			ViewportRect:     viewport,
			ClearFlags:       clearFlags,
		}, nil
	})
}
