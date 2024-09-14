package main

import (
	"gogame/cavescene"
	"gogame/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Dark Cave")

	scene := cavescene.NewCaveScene()

	game := game.NewGame([]game.Scene{&scene})
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
