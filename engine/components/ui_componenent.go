package components

import (
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type UIComponent struct {
	Anchor        string // "top-left", "top", "top-right", "left", "center", "right", "bottom-left", "bottom", "bottom-right"
	Pivot         mathx.Vector2
	Position      mathx.Vector2
	Size          mathx.Vector2
	Rotation      float64
	Scale         mathx.Vector2
	Visible       bool
	Interactable  bool
	RaycastTarget bool
	Layer         uint32
	Mask          uint32
	ClipChildren  bool
	AutoLayout    string        // "horizontal", "vertical", "none"
	Padding       mathx.Vector4 // top(X), bottom(Y), left(Z), right(W)
}

func (uic *UIComponent) Type() string { return constants.ComponentUIComponent }

func init() {
	RegisterComponent(constants.ComponentUIComponent, func(data map[string]interface{}) (RuntimeComponent, error) {
		anchor := "center"
		if val, ok := data[constants.FieldAnchor].(string); ok {
			anchor = val
		}

		pivot := mathx.Vector2{}
		if raw, ok := data[constants.FieldPivot].(map[string]any); ok {
			if parsed, err := mathx.FromMap2(raw); err == nil {
				pivot = parsed
			}
		}

		position := mathx.Vector2{}
		if raw, ok := data[constants.FieldPosition].(map[string]any); ok {
			if parsed, err := mathx.FromMap2(raw); err == nil {
				position = parsed
			}
		}

		size := mathx.Vector2{X: 1, Y: 1}
		if raw, ok := data[constants.FieldSize].(map[string]any); ok {
			if parsed, err := mathx.FromMap2(raw); err == nil {
				size = parsed
			}
		}

		rotation := 0.0
		if val, ok := data[constants.FieldRotation].(float64); ok {
			rotation = val
		}

		scale := mathx.Vector2{}
		if raw, ok := data[constants.FieldScale].(map[string]any); ok {
			if parsed, err := mathx.FromMap2(raw); err == nil {
				scale = parsed
			}
		}

		visible, _ := data[constants.FieldVisible].(bool)
		interactable, _ := data[constants.FieldInteractable].(bool)
		raycastTarget, _ := data[constants.FieldRaycastTarget].(bool)

		layer := uint32(0)
		if val, ok := data[constants.FieldLayer].(uint32); ok {
			layer = val
		}

		mask := uint32(0)
		if val, ok := data[constants.FieldMask].(uint32); ok {
			mask = val
		}

		clipChildren, _ := data[constants.FieldClipChildren].(bool)

		autoLayout := "none"
		if val, ok := data[constants.FieldAutoLayout].(string); ok {
			autoLayout = val
		}

		padding := mathx.Vector4{}
		if raw, ok := data[constants.FieldPadding].(map[string]any); ok {
			if parsed, err := mathx.FromMap4(raw); err == nil {
				padding = parsed
			}
		}

		validAnchors := map[string]bool{
			"top-left": true, "top": true, "top-right": true,
			"left": true, "center": true, "right": true,
			"bottom-left": true, "bottom": true, "bottom-right": true,
		}
		if !validAnchors[anchor] {
			anchor = "center"
		}

		validLayouts := map[string]bool{"horizontal": true, "vertical": true, "none": true}
		if !validLayouts[autoLayout] {
			autoLayout = "none"
		}

		return &UIComponent{
			Anchor:        anchor,
			Pivot:         pivot,
			Position:      position,
			Size:          size,
			Rotation:      rotation,
			Scale:         scale,
			Visible:       visible,
			Interactable:  interactable,
			RaycastTarget: raycastTarget,
			Layer:         layer,
			Mask:          mask,
			ClipChildren:  clipChildren,
			AutoLayout:    autoLayout,
			Padding:       padding,
		}, nil
	})
}
