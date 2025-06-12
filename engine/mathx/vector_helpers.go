package mathx

import (
	"fmt"
)

func FromMap2(data map[string]any) (Vector2, error) {
	x, okX := data["x"].(float64)
	y, okY := data["y"].(float64)

	if !okX || !okY {
		return Vector2{}, fmt.Errorf("Invalid vector2 format: expected float64 x, y fields.")
	}

	return Vector2{X: x, Y: y}, nil
}

func FromMap3(data map[string]any) (Vector3, error) {
	x, okX := data["x"].(float64)
	y, okY := data["y"].(float64)
	z, okZ := data["z"].(float64)

	if !okX || !okY || !okZ {
		return Vector3{}, fmt.Errorf("Invalid vector3 format: expected float64 x, y, z fields.")
	}

	return Vector3{X: x, Y: y, Z: z}, nil
}

func FromMap4(data map[string]any) (Vector4, error) {
	x, okX := data["x"].(float64)
	y, okY := data["y"].(float64)
	z, okZ := data["z"].(float64)
	w, okW := data["w"].(float64)

	if !okX || !okY || !okZ || !okW {
		return Vector4{}, fmt.Errorf("Invalid vector4 format: expected float64 x, y, z, w fields.")
	}

	return Vector4{X: x, Y: y, Z: z, W: w}, nil
}
