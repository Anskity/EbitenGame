package solid

import (
	"gogame/game"
	"gogame/gamemath"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// var spr *ebiten.Image
//
// func init() {
// 	spr_temp, _, err := ebitenutil.NewImageFromFile("./solid.png")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	spr = spr_temp
// }

type Solid struct {
	x int
	y int
}

func NewSolid(x, y int) Solid {
	return Solid{x, y}
}

func (_ *Solid) Update(_ *game.Game) {}

func (self *Solid) GetImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	canvas := ebiten.NewImage(16, 16)
	vector.DrawFilledRect(canvas, 0, 0, 16, 16, color.White, false)

	geom := ebiten.GeoM{}
	geom.Translate(float64(self.x), float64(self.y))
	opts := ebiten.DrawImageOptions{}
	opts.GeoM = geom

	return canvas, &opts
}

func (self *Solid) HasTag(tag string) bool {
	return tag == "solid"
}

func (s *Solid) SendSignal(_ int, _ []byte) ([]byte, error) {
	buf := gamemath.FloatsIntoBytes([]float64{
		float64(s.x),
		float64(s.y),
		16,
		16,
	})

	return buf, nil
}
