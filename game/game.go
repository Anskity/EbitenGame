package game

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Entity interface {
	Update(*Game)
	GetImage() (*ebiten.Image, *ebiten.DrawImageOptions)
	HasTag(string) bool
	SendSignal(int, []byte) ([]byte, error)
}

type Scene interface {
	Render(*ebiten.Image)
	Update(g *Game)
	GetEntitiesByTag(string) []Entity
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
	(*g.currentScene).Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	(*g.currentScene).Render(screen)

	txt := ""
	txt += strconv.FormatFloat(ebiten.ActualFPS(), 'f', -1, 64)
	txt += " - "
	txt += strconv.FormatFloat(ebiten.ActualTPS(), 'f', -1, 64)
	ebitenutil.DebugPrint(screen, txt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func (g *Game) GetScene() *Scene {
	return g.currentScene
}

func (g *Game) GetEntitiesByTag(tag string) []Entity {
	return (*g.GetScene()).GetEntitiesByTag(tag)
}
