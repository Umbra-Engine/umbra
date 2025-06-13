package main

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/logger"
	"github.com/Umbra-Engine/umbra/engine/scene"
)

func main() {
	s, err := scene.LoadScene("fixtures/valid.scene")
	if err != nil {
		logger.Error("%s", err.Error())
	}

	logger.Info("Scene loaded successfully: %s", s.Name)
	for _, obj := range s.Objects {
		fmt.Printf("- GameObject: %s at %v\n", obj.Name, obj.Position)
		logger.Info("- GameObject: %s at %v\n", obj.Name, obj.Position)
	}
}
