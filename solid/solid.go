package solid

import (
	"encoding/binary"
	"errors"
	"gogame/game"
	"gogame/gamemath"
	"image/color"
	"math"

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

func (self *Solid) Render(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(self.x), float32(self.y), 16, 16, color.White, false)
}

func (self *Solid) HasTag(tag string) bool {
	return tag == "solid"
}

func (self *Solid) SendSignal(_ int, posBin []byte) ([]byte, error) {
	if len(posBin) != 16 {
		return []byte{}, errors.New("Position data should be 16 bytes long")
	}

	frogW := float64(13)
	frogH := float64(8)

	xBin := binary.BigEndian.Uint64(posBin[0:8])
	yBin := binary.BigEndian.Uint64(posBin[8:16])
	x := math.Float64frombits(xBin) - frogW/2
	y := math.Float64frombits(yBin) - frogH/2

	if gamemath.RectangleInRectangle(x, y, frogW, frogH, float64(self.x), float64(self.y), 16, 16) {
		return []byte{1}, nil
	} else {
		return []byte{}, nil
	}
}
