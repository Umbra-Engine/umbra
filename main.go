package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: Render Scene
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Umbra Engine")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
