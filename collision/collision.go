package collision

import "gogame/gamemath"

type CollisionBox struct {
	X float64
	Y float64
	W float64
	H float64
}

func (b *CollisionBox) CollidesWithBox(x, y, w, h float64) bool {
	return gamemath.RectangleInRectangle(b.X, b.Y, b.W, b.H, x, y, w, h)
}

func (b *CollisionBox) CollidesWithPoint(x, y float64) bool {
	return gamemath.PointInRectangle(x, y, b.X, b.Y, b.W, b.H)
}
