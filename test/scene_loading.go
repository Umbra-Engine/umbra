package main

import (
	"fmt"
	"github.com/Umbra-Engine/umbra/engine/scene"
)

func main() {
	s, err := scene.LoadScene("fixtures/missing_name.scene")
	if err != nil {
		panic(err)
	}

	fmt.Println("Scene loaded successfully:", s.Name)
	for _, obj := range s.Objects {
		fmt.Printf("- GameObject: %s at %v\n", obj.Name, obj.Position)
	}
}
