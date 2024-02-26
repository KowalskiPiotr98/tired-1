package objects

import (
	"tired-1/drawing"
	"tired-1/utils"
)

type BallWallDirection int8

const (
	Top BallWallDirection = iota
	Right
	Bottom
	Left
)

type Ball struct {
	x, y, radius float64
	// movement vertex
	movX, movY float64
	colour     *drawing.Pixel
}

func NewBall(x, y, radius float64, colour *drawing.Pixel) *Ball {
	return &Ball{
		x:      x,
		y:      y,
		radius: radius,
		movX:   2,
		movY:   2,
		colour: colour,
	}
}

func (b *Ball) Move() {
	b.x += b.movX
	b.y += b.movY
}

func (b *Ball) Bounce(direction BallWallDirection) {
	//todo
}

func (b *Ball) Draw(grid *drawing.Grid) {
	for x := utils.Round(b.x - b.radius); x <= utils.Round(b.x+b.radius); x++ {
		for y := utils.Round(b.y - b.radius); y <= utils.Round(b.y+b.radius); y++ {
			xDist := utils.Round(b.x) - x
			yDist := utils.Round(b.y) - y
			dist := xDist*xDist + yDist*yDist
			if float64(dist) <= b.radius*b.radius {
				grid.SetPixel(x, y, b.colour)
			}
		}
	}
}
