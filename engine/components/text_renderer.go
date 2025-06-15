package components

import (
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type TextRendererComponent struct {
	Text                string
	Font                string
	FontSize            int
	Color               mathx.Vector4
	HorizontalAlignment string // "left", "center", "right"
	VerticalAlignment   string // "top", "middle", "bottom"
	LineSpacing         float64
	MaxWidth            float64
	Wrap                bool
	ZIndex              int
	Billboard           bool
	Layer               uint32
	Mask                uint32
}

func (tr *TextRendererComponent) Type() string { return constants.ComponentTextRenderer }

func init() {
	RegisterComponent(constants.ComponentTextRenderer, func(data map[string]interface{}) (RuntimeComponent, error) {
		text, _ := data[constants.FieldText].(string)

		font := "Ubuntu"
		if val, ok := data[constants.FieldFont].(string); ok {
			font = val
		}

		fontSize := 12
		if val, ok := data[constants.FieldFontSize].(int); ok && val > 0 {
			fontSize = val
		}

		color := mathx.Vector4{X: 1, Y: 1, Z: 1, W: 1}
		if raw, ok := data[constants.FieldColor].(map[string]any); ok {
			if parsed, err := mathx.FromMap4(raw); err == nil {
				color = parsed
			}
		}

		horizontalAlignment := "left"
		if val, ok := data[constants.FieldHorizontalAlignment].(string); ok {
			horizontalAlignment = val
		}

		verticalAlignment := "top"
		if val, ok := data[constants.FieldVerticalAlignment].(string); ok {
			verticalAlignment = val
		}

		lineSpacing := 1.0
		if val, ok := data[constants.FieldLineSpacing].(float64); ok && val >= 0 {
			lineSpacing = val
		}

		maxWidth := 1.0
		if val, ok := data[constants.FieldMaxWidth].(float64); ok && val >= 0 {
			maxWidth = val
		}

		wrap, _ := data[constants.FieldWrap].(bool)
		billboard, _ := data[constants.FieldBillboard].(bool)

		zIndex := 1
		if val, ok := data[constants.FieldZIndex].(int); ok {
			zIndex = val
		}

		layer := uint32(0)
		if val, ok := data[constants.FieldLayer].(uint32); ok {
			layer = val
		}

		mask := uint32(0)
		if val, ok := data[constants.FieldMask].(uint32); ok {
			mask = val
		}

		validHAlign := map[string]bool{"left": true, "center": true, "right": true}
		if !validHAlign[horizontalAlignment] {
			horizontalAlignment = "left"
		}

		validVAlign := map[string]bool{"top": true, "middle": true, "bottom": true}
		if !validVAlign[verticalAlignment] {
			verticalAlignment = "top"
		}

		return &TextRendererComponent{
			Text:                text,
			Font:                font,
			FontSize:            fontSize,
			Color:               color,
			HorizontalAlignment: horizontalAlignment,
			VerticalAlignment:   verticalAlignment,
			LineSpacing:         lineSpacing,
			MaxWidth:            maxWidth,
			Wrap:                wrap,
			ZIndex:              zIndex,
			Billboard:           billboard,
			Layer:               layer,
			Mask:                mask,
		}, nil
	})
}
