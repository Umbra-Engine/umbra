package components

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/constants"
)

type AudioSourceComponent struct {
	Clip        string
	Volume      float64
	Pitch       float64
	Loop        bool
	PlayOnStart bool
	IsMuted     bool
	Spatial     bool
	MinDistance float64
	MaxDistance float64
	Rolloff     string // "linear", "log", "none"
	Priority    int
	OutputGroup string
}

func (as *AudioSourceComponent) Type() string { return constants.ComponentAudioSource }

func init() {
	RegisterComponent(constants.ComponentAudioSource, func(data map[string]interface{}) (RuntimeComponent, error) {
		clip, ok := data[constants.FieldClip].(string)
		if !ok || clip == "" {
			return nil, fmt.Errorf("missing or invalid sound clip.")
		}

		volume := 1.0
		if val, ok := data[constants.FieldVolume].(float64); ok && val >= 0 && val <= 1.0 {
			volume = val
		}

		pitch := 1.0
		if val, ok := data[constants.FieldPitch].(float64); ok {
			pitch = val
		}

		loop := false
		if val, ok := data[constants.FieldLoop].(bool); ok {
			loop = val
		}

		playOnStart := false
		if val, ok := data[constants.FieldPlayOnStart].(bool); ok {
			playOnStart = val
		}

		isMuted := false
		if val, ok := data[constants.FieldIsMuted].(bool); ok {
			isMuted = val
		}

		spatial := false
		if val, ok := data[constants.FieldSpatial].(bool); ok {
			spatial = val
		}

		minDist := 0.0
		if val, ok := data[constants.FieldMinDistance].(float64); ok && val >= 0 {
			minDist = val
		}

		maxDist := 1000.0
		if val, ok := data[constants.FieldMaxDistance].(float64); ok && val >= 0 {
			maxDist = val
		}

		rolloff := "none"
		if val, ok := data[constants.FieldRolloff].(string); ok {
			rolloff = val
		}

		priority := 1
		if val, ok := data[constants.FieldPriority].(int); ok && val >= 0 {
			priority = val
		}

		outputGroup, _ := data[constants.FieldOutputGroup].(string)

		validRolloffs := map[string]bool{
			"none": true, "linear": true, "log": true,
		}
		if !validRolloffs[rolloff] {
			return nil, fmt.Errorf("invalid rolloff type: %s", rolloff)
		}

		return &AudioSourceComponent{
			Clip:        clip,
			Volume:      volume,
			Pitch:       pitch,
			Loop:        loop,
			PlayOnStart: playOnStart,
			IsMuted:     isMuted,
			Spatial:     spatial,
			MinDistance: minDist,
			MaxDistance: maxDist,
			Rolloff:     rolloff,
			Priority:    priority,
			OutputGroup: outputGroup,
		}, nil
	})
}
