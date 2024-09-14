package player

import (
	"encoding/binary"
	"gogame/game"
	"gogame/gamemath"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	PosX float64
	PosY float64
	Spr  *ebiten.Image
}

func CreatePlayer() Player {
	spr, _, err := ebitenutil.NewImageFromFile("frog.png")

	if err != nil {
		panic(err)
	}

	return Player{
		PosX: 0,
		PosY: 0,
		Spr:  spr,
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
	for _, entity := range *(*g.GetScene()).GetEntities() {
		if !entity.HasTag("solid") {
			continue
		}

		var buf [16]byte
		binary.BigEndian.PutUint64(buf[0:8], math.Float64bits(x))
		binary.BigEndian.PutUint64(buf[8:16], math.Float64bits(y))
		collides, err := entity.SendSignal(0, buf[:])
		if err != nil {
			panic(err)
		}

		if len(collides) != 0 {
			return true
		}
	}

	return false
}

func (player *Player) Update(g *game.Game) {
	player.movement(g)
}

func (player *Player) Render(screen *ebiten.Image) {
	if player.Spr == nil {
		log.Fatal("Player was expected to have a sprite")
	}

	sprSize := player.Spr.Bounds().Size()

	pos := ebiten.GeoM{}
	pos.Translate(player.PosX-float64(sprSize.X/2), player.PosY-float64(sprSize.Y)/2)
	screen.DrawImage(player.Spr, &ebiten.DrawImageOptions{
		GeoM: pos,
	})
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
