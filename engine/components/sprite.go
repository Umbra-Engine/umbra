package components

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/constants"
)

type SpriteRendererComponent struct {
	Texture string
	Layer   int
	FlipX   bool
	FlipY   bool
	Visible bool
}

func (s *SpriteRendererComponent) Type() string { return constants.ComponentSpriteRenderer }

func init() {
	RegisterComponent(constants.ComponentSpriteRenderer, func(data map[string]interface{}) (RuntimeComponent, error) {
		texture, ok := data[constants.FieldTexture].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid texture.")
		}

		layer := 0
		if l, ok := data[constants.FieldLayer].(int); ok {
			layer = l
		}

		flipX, _ := data[constants.FieldFlipX].(bool)
		flipY, _ := data[constants.FieldFlipY].(bool)
		visible, _ := data[constants.FieldVisible].(bool)

		return &SpriteRendererComponent{
			Texture: texture,
			Layer:   layer,
			FlipX:   flipX,
			FlipY:   flipY,
			Visible: visible,
		}, nil
	})
}
