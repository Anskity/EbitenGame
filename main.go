package main

import (
	"gogame/game"
	"gogame/scenes/cavescene"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280/2, 720/2)
	ebiten.SetWindowTitle("Dark Cave")

	scene := cavescene.NewCaveScene()

	game := game.NewGame([]game.Scene{&scene})
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
