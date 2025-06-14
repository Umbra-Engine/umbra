package components

import (
	"fmt"

	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

// Collider2DType defines supported 2D collider shapes
type Collider2DType string

const (
	Collider2DBox     Collider2DType = "box"
	Collider2DCircle  Collider2DType = "circle"
	Collider2DPill    Collider2DType = "pill"
	Collider2DPolygon Collider2DType = "polygon"
	Collider2DEdge    Collider2DType = "edge"
)

// Collider2DComponent describes a 2D collider's properties
type Collider2DComponent struct {
	ColliderType Collider2DType  // Shape type
	Offset       mathx.Vector2   // Local offset
	Size         mathx.Vector2   // Used for box/pill
	Radius       float64         // Used for circle/pill
	Points       []mathx.Vector2 // Used for polygon/edge
	IsTrigger    bool            // Trigger vs solid collider
	Layer        uint32          // Collision layer
	Mask         uint32          // Collision mask
}

// Type returns the component's type identifier
func (c *Collider2DComponent) Type() string {
	return constants.ComponentCollider2D
}

// init registers the Collider2DComponent with the component registry
func init() {
	RegisterComponent(constants.ComponentCollider2D, func(data map[string]interface{}) (RuntimeComponent, error) {
		// Collider type
		cTypeStr, _ := data[constants.FieldType].(string)
		cType := Collider2DType(cTypeStr)

		// Offset
		offset := mathx.Vector2{}
		if raw, ok := data[constants.FieldOffset].(map[string]any); ok {
			offset, _ = mathx.FromMap2(raw)
		}

		// Size (for box/pill)
		size := mathx.Vector2{}
		if raw, ok := data[constants.FieldSize].(map[string]any); ok {
			size, _ = mathx.FromMap2(raw)
		}

		// Radius (for circle/pill)
		radius := 0.0
		if r, ok := data[constants.FieldRadius].(float64); ok {
			radius = r
		}

		// IsTrigger
		isTrigger := false
		if b, ok := data[constants.FieldIsTrigger].(bool); ok {
			isTrigger = b
		}

		// Layer and mask
		layer := uint32(0)
		if val, ok := data[constants.FieldLayer].(int); ok {
			layer = uint32(val)
		}

		mask := uint32(0xFFFFFFFF)
		if val, ok := data[constants.FieldMask].(int); ok {
			mask = uint32(val)
		}

		// Points (for polygon/edge)
		points := []mathx.Vector2{}
		if raw, ok := data[constants.FieldPoints].([]interface{}); ok {
			for _, item := range raw {
				if pointMap, ok := item.(map[string]any); ok {
					if pt, err := mathx.FromMap2(pointMap); err == nil {
						points = append(points, pt)
					}
				}
			}
		}

		// --- Validation based on collider type ---
		switch cType {
		case Collider2DBox:
			if size.IsZero(1e-5) {
				return nil, fmt.Errorf("Collider2D 'box' requires non-zero size")
			}
		case Collider2DCircle:
			if radius <= 0 {
				return nil, fmt.Errorf("Collider2D 'circle' requires a positive radius")
			}
		case Collider2DPill:
			if radius <= 0 || size.X <= 0 {
				return nil, fmt.Errorf("Collider2D 'pill' requires a positive radius and non-zero width (size.x)")
			}
		case Collider2DPolygon:
			if len(points) < 3 {
				return nil, fmt.Errorf("Collider2D 'polygon' requires at least 3 points")
			}
		case Collider2DEdge:
			if len(points) < 2 {
				return nil, fmt.Errorf("Collider2D 'edge' requires at least 2 points")
			}
		default:
			return nil, fmt.Errorf("unknown Collider2D type: %s", cType)
		}

		// Final component
		return &Collider2DComponent{
			ColliderType: cType,
			Offset:       offset,
			Size:         size,
			Radius:       radius,
			Points:       points,
			IsTrigger:    isTrigger,
			Layer:        layer,
			Mask:         mask,
		}, nil
	})
}
