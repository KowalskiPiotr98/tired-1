package objects

import (
	"tired-1/drawing"
	"tired-1/utils"
)

type Tile struct {
	x, y, side float64
	colour     *drawing.Pixel
}

func NewTile(x, y, side float64, colour *drawing.Pixel) *Tile {
	return &Tile{
		x:      x,
		y:      y,
		side:   side,
		colour: colour,
	}
}

func (t *Tile) Draw(grid *drawing.Grid) {
	for x := t.x - t.side/2; x < t.x+t.side/2; x++ {
		for y := t.y - t.side/2; y < t.y+t.side/2; y++ {
			grid.SetPixel(utils.Round(x), utils.Round(y), t.colour)
		}
	}
}

func (t *Tile) HandleBallTouch(ball *Ball) {
	if *ball.touchColour == *t.colour {
		// ignore same touchColour balls
		return
	}

	direction := areTouching(ball, t)
	if direction == 0 {
		return
	}

	// steal ball touchColour
	t.colour = ball.touchColour
	ball.Bounce(direction)
}
