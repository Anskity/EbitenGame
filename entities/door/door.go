package door

import (
	"gogame/game"
	"gogame/gamemath"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var sprClosed *ebiten.Image

func init() {
	sprClosed = ebiten.NewImage(16, 16)
	clr := color.RGBA{}
	clr.G = 255
	vector.DrawFilledRect(sprClosed, 0, 0, 16, 16, clr, false)
}

type Door struct {
	X      int
	Y      int
	Opened bool
}

func NewDoor(x int, y int) Door {
	return Door{
		X:      x,
		Y:      y,
		Opened: false,
	}
}

func (d *Door) Update(g *game.Game) {

}
func (d *Door) GetImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(d.X), float64(d.Y))

	return sprClosed, &opts
}
func (d *Door) HasTag(tag string) bool {
	return tag == "door" || tag == "solid"
}
func (d *Door) SendSignal(_ int, _ []byte) ([]byte, error) {
	bin := gamemath.FloatsIntoBytes([]float64{
		float64(d.X),
		float64(d.Y),
		16,
		16,
	})
	return bin, nil
}
