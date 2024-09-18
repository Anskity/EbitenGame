package cavescene

import (
	"gogame/entities/camera"
	"gogame/entities/door"
	"gogame/entities/player"
	"gogame/entities/solid"
	"gogame/game"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var shader *ebiten.Shader

func init() {
	blob, err := os.ReadFile("Shaders/lightshader.kage")
	if err != nil {
		panic(err)
	}

	shader, err = ebiten.NewShader(blob)
	if err != nil {
		panic(err)
	}
}

type CaveScene struct {
	mainPlayer player.Player
	solids     []solid.Solid
	doors      []door.Door
	cam        camera.Camera
}

func NewCaveScene() CaveScene {
	scene := CaveScene{}

	for i := 0; i < 10; i++ {
		solidObj := solid.NewSolid(i*16, 32)
		scene.solids = append(scene.solids, solidObj)
	}

	scene.mainPlayer = player.CreatePlayer()
	scene.mainPlayer.PosX = 64
	scene.mainPlayer.PosY = 64
	scene.cam = camera.NewCamera()

	scene.doors = append(scene.doors, door.NewDoor(0, 0))

	return scene
}

func (s *CaveScene) Update(g *game.Game) {
	s.mainPlayer.Update(g)
	s.cam.Update(&s.mainPlayer)

	for _, door := range s.doors {
		door.Update(g)
	}
}

func (scene *CaveScene) Render(screen *ebiten.Image) {
	brown := color.RGBA{}
	brown.R = 255
	brown.G = 50

	screen.Fill(brown)

	entities := []game.Entity{}
	for _, solid_entity := range scene.solids {
		entities = append(entities, &solid_entity)
	}
	for _, door_entity := range scene.doors {
		entities = append(entities, &door_entity)
	}
	entities = append(entities, &scene.mainPlayer)
	size := screen.Bounds().Size()

	cam_img, cam_opts := scene.cam.GetImage(&entities, 500, 500)
	cam_xscale := float64(size.X) / float64(scene.cam.Width)
	cam_yscale := float64(size.Y) / float64(scene.cam.Height)
	cam_opts.GeoM.Scale(cam_xscale, cam_yscale)
	screen.DrawImage(cam_img, cam_opts)

	camXoffset, camYoffset := scene.cam.GetOffset(500, 500)

	opts := ebiten.DrawRectShaderOptions{}
	opts.Uniforms = map[string]any{
		"LightPos": [2]float64{
			scene.mainPlayer.PosX*cam_xscale - camXoffset*cam_xscale,
			scene.mainPlayer.PosY*cam_yscale - camYoffset*cam_yscale,
		},
		"Size": [2]float64{
			float64(scene.cam.Width),
			float64(scene.cam.Height),
		},
	}
	screen.DrawRectShader(size.X, size.Y, shader, &opts)
}

func (s *CaveScene) GetEntitiesByTag(tag string) []game.Entity {
	entities := []game.Entity{}
	if s.mainPlayer.HasTag(tag) {
		entities = append(entities, &s.mainPlayer)
	}

	for _, solid := range s.solids {
		if solid.HasTag(tag) {
			entities = append(entities, &solid)
		}
	}

	for _, door := range s.doors {
		if door.HasTag(tag) {
			entities = append(entities, &door)
		}
	}

	return entities
}
