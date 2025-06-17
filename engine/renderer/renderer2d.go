package renderer

import (
	"github.com/Umbra-Engine/umbra/engine/components"
	"github.com/Umbra-Engine/umbra/engine/scene"
)

// DrawScene2D renders all visible sprites in the scene.
func DrawScene2D(s *scene.Scene) {
	for _, obj := range s.Objects {
		drawObject2DRecursive(obj)
	}
}

// drawObject2DRecursive draws a single GameObject and its children (if any).
func drawObject2DRecursive(obj scene.GameObject) {
	var transform *components.TransformComponent
	var sprite *components.SpriteRendererComponent

	for _, comp := range obj.RuntimeComponents {
		switch c := comp.(type) {
		case *components.TransformComponent:
			transform = c
		case *components.SpriteRendererComponent:
			sprite = c
		}
	}

	if transform != nil && sprite != nil && sprite.Visible {
		DrawSprite(*transform, *sprite)
	}

	// Recurse into children
	for _, child := range obj.Children {
		drawObject2DRecursive(child)
	}
}

// DrawSprite is a stub for now — in the future this should use your rendering backend.
func DrawSprite(transform components.TransformComponent, sprite components.SpriteRendererComponent) {
	// TODO: Integrate with actual drawing system
	// Example debug output:
	println("Drawing sprite:", sprite.Texture, "at", transform.Position)
}
