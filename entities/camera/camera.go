package camera

import (
	"gogame/entities/player"
	"gogame/game"
	"gogame/gamemath"

	"github.com/hajimehoshi/ebiten/v2"
)

const SUBPIXELS int = 4

type Camera struct {
	targetPlayer *player.Player
	entities     *[]game.Entity
	X            float64
	Y            float64
	Width        int
	Height       int

	imgBuf *ebiten.Image
}

func NewCamera() Camera {
	w := 160 * 2
	h := 90 * 2

	return Camera{
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
		imgBuf: ebiten.NewImage(w*SUBPIXELS, h*SUBPIXELS),
	}
}

func (c *Camera) Update(targetPlayer *player.Player) {
	if targetPlayer == nil {
		panic("Got nil targetPlayer")
	}

	c.X = targetPlayer.PosX
	c.Y = targetPlayer.PosY
}

func (c *Camera) GetOffset(roomWidth, roomHeight float64) (float64, float64) {
	xoffset := gamemath.Clamp(c.X-float64(c.Width)/2, 0, roomWidth-float64(c.Width))
	yoffset := gamemath.Clamp(c.Y-float64(c.Height)/2, 0, roomHeight-float64(c.Height))

	return xoffset, yoffset
}

func (c *Camera) GetImage(entities *[]game.Entity, roomWidth int, roomHeight int) (*ebiten.Image, *ebiten.DrawImageOptions) {
	xoffset, yoffset := c.GetOffset(float64(roomWidth), float64(roomHeight))

	c.imgBuf.Clear()
	bufopts := ebiten.DrawImageOptions{}
	bufopts.GeoM.Scale(1/float64(SUBPIXELS), 1/float64(SUBPIXELS))

	for _, entity := range *entities {
		img, opts := entity.GetImage()
		opts.GeoM.Translate(-xoffset, -yoffset)
		opts.GeoM.Scale(float64(SUBPIXELS), float64(SUBPIXELS))

		c.imgBuf.DrawImage(img, opts)
	}

	return c.imgBuf, &bufopts
}
