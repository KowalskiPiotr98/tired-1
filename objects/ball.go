package objects

import (
	"math"
	"tired-1/drawing"
	"tired-1/utils"
)

type Ball struct {
	x, y, radius float64
	// movement vertex
	movX, movY  float64
	touchColour *drawing.Pixel
	drawColour  *drawing.Pixel
}

func NewBall(x, y, radius float64, touchColour *drawing.Pixel, drawColour *drawing.Pixel) *Ball {
	return &Ball{
		x:           x,
		y:           y,
		radius:      radius,
		movX:        2,
		movY:        2,
		touchColour: touchColour,
		drawColour:  drawColour,
	}
}

func (b *Ball) Move() {
	b.x += b.movX
	b.y += b.movY
}

func (b *Ball) Bounce(direction TouchDirection) {
	switch direction {
	case touchTop:
		b.movY = math.Abs(b.movY)
	case touchRight:
		b.movX = -math.Abs(b.movX)
	case touchBottom:
		b.movY = -math.Abs(b.movY)
	case touchLeft:
		b.movX = math.Abs(b.movX)
	case touchNone:
	default:
		return
	}
}

func (b *Ball) Draw(grid *drawing.Grid) {
	for x := utils.Round(b.x - b.radius); x <= utils.Round(b.x+b.radius); x++ {
		for y := utils.Round(b.y - b.radius); y <= utils.Round(b.y+b.radius); y++ {
			xDist := utils.Round(b.x) - x
			yDist := utils.Round(b.y) - y
			dist := xDist*xDist + yDist*yDist
			if float64(dist) <= b.radius*b.radius {
				grid.SetPixel(x, y, b.drawColour)
			}
		}
	}
}
