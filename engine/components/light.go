package components

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/constants"
	"github.com/Umbra-Engine/umbra/engine/mathx"
)

type LightType string

const (
	LightPoint       LightType = "point"
	LightDirectional LightType = "directional"
	LightSpot        LightType = "spot"
)

type FalloffType string

const (
	FalloffNone      FalloffType = "none"
	FalloffLinear    FalloffType = "linear"
	FalloffQuadratic FalloffType = "quadratic"
)

type LightComponent struct {
	LightType    LightType
	Color        mathx.Vector3
	Intensity    float64
	Range        float64 // for point and spot
	Angle        float64 // spotlight angle (in degrees)
	CastsShadows bool
	Layer        uint32
	Mask         uint32
	Falloff      FalloffType
}

func (l *LightComponent) Type() string { return constants.ComponentLight }

func init() {
	RegisterComponent(constants.ComponentLight, func(data map[string]interface{}) (RuntimeComponent, error) {
		lTypeStr, _ := data[constants.FieldType].(string)
		lType := LightType(lTypeStr)

		color := mathx.Vector3{}
		if raw, ok := data[constants.FieldColor].(map[string]any); ok {
			color, _ = mathx.FromMap3(raw)
		}

		intensity := 0.0
		if val, ok := data[constants.FieldIntensity].(float64); ok && val >= 0 {
			intensity = val
		}
		rangeRaw := 0.0
		if val, ok := data[constants.FieldRange].(float64); ok && val >= 0 {
			rangeRaw = val
		}
		angle := 0.0
		if val, ok := data[constants.FieldAngle].(float64); ok && val >= 0 {
			angle = val
		}

		castShadows, _ := data[constants.FieldCastShadows].(bool)

		layer := uint32(0)
		if val, ok := data[constants.FieldLayer].(int); ok {
			layer = uint32(val)
		}

		mask := uint32(0xFFFFFFFF)
		if val, ok := data[constants.FieldMask].(int); ok {
			mask = uint32(val)
		}

		falloffStr, _ := data[constants.FieldFalloff].(string)
		falloff := FalloffType(falloffStr)

		switch lType {
		case LightPoint, LightSpot, LightDirectional:
			// ok
		default:
			return nil, fmt.Errorf("invalid light type: %s", lType)
		}

		switch falloff {
		case FalloffNone, FalloffLinear, FalloffQuadratic:
			// ok
		default:
			return nil, fmt.Errorf("invalid falloff type: %s", falloff)
		}

		return &LightComponent{
			LightType:    lType,
			Color:        color,
			Intensity:    intensity,
			Range:        rangeRaw,
			Angle:        angle,
			CastsShadows: castShadows,
			Layer:        layer,
			Mask:         mask,
			Falloff:      falloff,
		}, nil
	})
}
