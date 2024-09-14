package gamemath

import (
	"math"
)

type Int interface {
	int | int8 | int16 | int32 | int64
}
type Float interface {
	float32 | float64
}

type Number interface {
	Int | Float
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func PointDirection(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(y2-y1, x2-x1)
}

func PointInRectangle[T Number](px, py, rx, ry, w, h T) bool {
	x1 := rx
	y1 := ry
	x2 := x1 + w
	y2 := y1 + h

	return (px > x1 && px < x2) && (py > y1 && py < y2)
}

func RectangleInRectangle[T Number](rx1, ry1, rw1, rh1, rx2, ry2, rw2, rh2 T) bool {
	ax1 := rx1
	ay1 := ry1
	ax2 := ax1 + rw1
	ay2 := ay1 + rh1

	bx1 := rx2
	by1 := ry2
	bx2 := bx1 + rw2
	by2 := by1 + rh2

	meetsX := (ax1 > bx1 && ax1 < bx2) || (ax2 > bx1 && ax2 < bx2)
	meetsY := (ay1 > by1 && ay1 < by2) || (ay2 > by1 && ay2 < by2)

	return meetsX && meetsY
}

func Sign[T Number](val T) T {
	if val < 0 {
		return -1
	} else if val > 0 {
		return 1
	} else {
		return 0
	}
}
