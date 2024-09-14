package cavescene

import (
	"gogame/game"
	"gogame/player"
	"gogame/solid"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var shader *ebiten.Shader

func init() {
	blob, err := os.ReadFile("./lightshader.kage")
	if err != nil {
		panic(err)
	}

	shader, err = ebiten.NewShader(blob)
	if err != nil {
		panic(err)
	}
}

type CaveScene struct {
	entities   []game.Entity
	mainPlayer *player.Player
}

func NewCaveScene() CaveScene {
	scene := CaveScene{}

	for i := 0; i < 10; i++ {
		solidObj := solid.NewSolid(i*16, 32)
		scene.PushEntity(&solidObj)
	}

	p := player.CreatePlayer()
	p.PosX = 64
	p.PosY = 64

	scene.PushEntity(&p)
	scene.mainPlayer = &p

	return scene
}

func (s *CaveScene) GetEntities() *[]game.Entity {
	return &s.entities
}

func (scene *CaveScene) Render(screen *ebiten.Image) {
	brown := color.RGBA{}
	brown.R = 255

	screen.Fill(brown)
	for _, entity := range scene.entities {
		entity.Render(screen)
	}

	size := screen.Bounds().Size()

	opts := ebiten.DrawRectShaderOptions{}
	opts.Uniforms = map[string]any{
		"LightPos": [2]float64{
			scene.mainPlayer.PosX,
			scene.mainPlayer.PosY,
		},
		"Size": [2]float64{
			float64(size.X),
			float64(size.Y),
		},
	}
	screen.DrawRectShader(size.X, size.Y, shader, &opts)
}

func (s *CaveScene) PushEntity(entity game.Entity) {
	s.entities = append(s.entities, entity)
}
