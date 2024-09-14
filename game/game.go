package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Update(*Game)
	Render(*ebiten.Image)
	HasTag(string) bool
	SendSignal(int, []byte) ([]byte, error)
}

type Scene interface {
	GetEntities() *[]Entity
	Render(*ebiten.Image)
}

type Game struct {
	count        int
	scenes       []Scene
	currentScene *Scene
}

func NewGame(scenes []Scene) Game {
	return Game{
		count:        0,
		scenes:       scenes,
		currentScene: &scenes[0],
	}
}

func (g *Game) Update() error {
	g.count += 1
	for _, entity := range *(*g.currentScene).GetEntities() {
		entity.Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	(*g.currentScene).Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) GetScene() *Scene {
	return g.currentScene
}
