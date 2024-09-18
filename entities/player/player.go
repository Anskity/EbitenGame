package player

import (
	"encoding/binary"
	"gogame/assert"
	"gogame/game"
	"gogame/gamemath"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var idleSpr *ebiten.Image

func init() {
	spr, _, err := ebitenutil.NewImageFromFile("Sprites/frog.png")
	idleSpr = spr

	if err != nil {
		panic(err)
	}
}

type Player struct {
	PosX float64
	PosY float64
}

func CreatePlayer() Player {
	return Player{
		PosX: 0,
		PosY: 0,
	}
}

func (player *Player) Move(hsp float64, vsp float64, g *game.Game) {
	if pointIsSolid(player.PosX+hsp, player.PosY, g) {
		if hsp > 0 {
			player.PosX = math.Floor(player.PosX)
		} else {
			player.PosX = math.Ceil(player.PosX)
		}
		for !pointIsSolid(player.PosX+gamemath.Sign(hsp), player.PosY, g) {
			player.PosX += gamemath.Sign(hsp)
		}
		hsp = 0
	}
	player.PosX += hsp

	if pointIsSolid(player.PosX, player.PosY+vsp, g) {
		if vsp > 0 {
			player.PosY = math.Floor(player.PosY)
		} else {
			player.PosY = math.Ceil(player.PosY)
		}
		for !pointIsSolid(player.PosX, player.PosY+gamemath.Sign(vsp), g) {
			player.PosY += gamemath.Sign(vsp)
		}
		vsp = 0
	}
	player.PosY += vsp
}

func pointIsSolid(x float64, y float64, g *game.Game) bool {
	const playerW float64 = 13
	const playerH float64 = 9
	for _, entity := range g.GetEntitiesByTag("solid") {
		buf, err := entity.SendSignal(0, []byte{})
		if err != nil {
			panic(err)
		}

		assert.AssertEq[uint](len(buf), 32)

		colX := math.Float64frombits(binary.BigEndian.Uint64(buf[0:8]))
		colY := math.Float64frombits(binary.BigEndian.Uint64(buf[8:16]))
		colW := math.Float64frombits(binary.BigEndian.Uint64(buf[16:24]))
		colH := math.Float64frombits(binary.BigEndian.Uint64(buf[24:32]))

		if gamemath.RectangleInRectangle(colX, colY, colW, colH, x-playerW/2, y-playerH/2, playerW, playerH) {
			return true
		}
	}

	return false
}

func (player *Player) Update(g *game.Game) {
	player.movement(g)
}

func (player *Player) GetImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	sprSize := idleSpr.Bounds().Size()

	pos := ebiten.GeoM{}
	pos.Translate(player.PosX-float64(sprSize.X/2), player.PosY-float64(sprSize.Y)/2)
	return idleSpr, &ebiten.DrawImageOptions{
		GeoM: pos,
	}
}

func (player *Player) movement(g *game.Game) {
	leftInput := ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	rightInput := ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	upInput := ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	downInput := ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)

	moveX := gamemath.BoolToInt(rightInput) - gamemath.BoolToInt(leftInput)
	moveY := gamemath.BoolToInt(downInput) - gamemath.BoolToInt(upInput)

	dir := gamemath.PointDirection(0, 0, float64(moveX), float64(moveY))
	hsp := math.Cos(dir)
	vsp := math.Sin(dir)

	if moveX != 0 || moveY != 0 {
		player.Move(hsp, vsp, g)
	}
}

func (_ *Player) HasTag(tag string) bool {
	return tag == "player"
}

func (_ *Player) SendSignal(_ int, _ []byte) ([]byte, error) {
	return []byte{}, nil
}
