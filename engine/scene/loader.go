package scene

import (
	"gopkg.in/yaml.v3"
	"os"

	"github.com/Umbra-Engine/umbra/engine/logger"
)

func LoadScene(path string) (*Scene, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var scene Scene
	if err := yaml.Unmarshal(data, &scene); err != nil {
		return nil, err
	}

	// Validate the scene
	if errs := scene.Validate(); len(errs) > 0 {
		for _, e := range errs {
			logger.Error("Scene validation error: %s", e.Error())
		}
		logger.Fatal("scene validation failed.")
	}

	// Apply defaults to all objects
	for i := range scene.Objects {
		scene.Objects[i].InitDefaults()
	}

	return &scene, nil
}
