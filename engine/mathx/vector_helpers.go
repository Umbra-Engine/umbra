package mathx

import (
	"fmt"
	"math"
)

// Vector 2 Helpers //

func FromMap2(data map[string]any) (Vector2, error) {
	x, okX := data["x"].(float64)
	y, okY := data["y"].(float64)

	if !okX || !okY {
		return Vector2{}, fmt.Errorf("Invalid vector2 format: expected float64 x, y fields.")
	}

	return Vector2{X: x, Y: y}, nil
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2) Div(scalar float64) Vector2 {
	if scalar == 0 {
		panic("division by zero in Vector2.Div()")
	}
	return Vector2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Cross(other Vector2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vector2) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vector2) Normalize() Vector2 {
	length := v.Length()
	if length == 0 {
		return Vector2{0, 0}
	}
	return v.Div(length)
}

func (v Vector2) Equals(other Vector2, epsilon float64) bool {
	return math.Abs(v.X-other.X) < epsilon &&
		math.Abs(v.Y-other.Y) < epsilon
}

func (v Vector2) IsZero(epsilon float64) bool {
	return math.Abs(v.X) < epsilon &&
		math.Abs(v.Y) < epsilon
}

func ZeroVector2() Vector2 {
	return Vector2{X: 0, Y: 0}
}

// Vector 3 Helpers //

func FromMap3(data map[string]any) (Vector3, error) {
	x, okX := data["x"].(float64)
	y, okY := data["y"].(float64)
	z, okZ := data["z"].(float64)

	if !okX || !okY || !okZ {
		return Vector3{}, fmt.Errorf("Invalid vector3 format: expected float64 x, y, z fields.")
	}

	return Vector3{X: x, Y: y, Z: z}, nil
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vector3) Mul(scalar float64) Vector3 {
	return Vector3{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func (v Vector3) Div(scalar float64) Vector3 {
	if scalar == 0 {
		panic("division by zero in Vector3.Div()")
	}
	return Vector3{
		X: v.X / scalar,
		Y: v.Y / scalar,
		Z: v.Z / scalar,
	}
}

func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vector3) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v Vector3) Normalize() Vector3 {
	length := v.Length()
	if length == 0 {
		return Vector3{0, 0, 0}
	}
	return v.Div(length)
}

func (v Vector3) Equals(other Vector3, epsilon float64) bool {
	return math.Abs(v.X-other.X) < epsilon &&
		math.Abs(v.Y-other.Y) < epsilon &&
		math.Abs(v.Z-other.Z) < epsilon
}

func (v Vector3) IsZero(epsilon float64) bool {
	return math.Abs(v.X) < epsilon &&
		math.Abs(v.Y) < epsilon &&
		math.Abs(v.Z) < epsilon
}

func ZeroVector3() Vector3 {
	return Vector3{X: 0, Y: 0, Z: 0}
}

// Vector4 Helpers //

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

func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
		W: v.W + other.W,
	}
}

func (v Vector4) Sub(other Vector4) Vector4 {
	return Vector4{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Vector4) Mul(scalar float64) Vector4 {
	return Vector4{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
		W: v.W * scalar,
	}
}

func (v Vector4) Div(scalar float64) Vector4 {
	if scalar == 0 {
		panic("division by zero in Vector4.Div()")
	}
	return Vector4{
		X: v.X / scalar,
		Y: v.Y / scalar,
		Z: v.Z / scalar,
		W: v.W / scalar,
	}
}

func (v Vector4) Dot(other Vector4) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

func (v Vector4) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2))
}

func (v Vector4) Normalize() Vector4 {
	length := v.Length()
	if length == 0 {
		return Vector4{0, 0, 0, 0}
	}
	return v.Div(length)
}

func (v Vector4) Equals(other Vector4, epsilon float64) bool {
	return math.Abs(v.X-other.X) < epsilon &&
		math.Abs(v.Y-other.Y) < epsilon &&
		math.Abs(v.Z-other.Z) < epsilon &&
		math.Abs(v.W-other.W) < epsilon
}

func (v Vector4) IsZero(epsilon float64) bool {
	return math.Abs(v.X) < epsilon &&
		math.Abs(v.Y) < epsilon &&
		math.Abs(v.Z) < epsilon &&
		math.Abs(v.W) < epsilon
}

func ZeroVector4() Vector4 {
	return Vector4{X: 0, Y: 0, Z: 0, W: 0}
}
