package objects

import (
	"math"
	"tired-1/drawing"
)

type TouchDirection int8

const (
	touchNone TouchDirection = iota
	touchTop
	touchRight
	touchBottom
	touchLeft
)

func areTouching(ball *Ball, tile *Tile) TouchDirection {
	if (math.Abs(ball.x - tile.x)) > ball.radius*2 {
		return touchNone
	}
	if (math.Abs(ball.y - tile.y)) > ball.radius*2 {
		return touchNone
	}
	//todo
	return touchRight
}

func BounceWall(ball *Ball, grid *drawing.Grid) {
	if ball.x-ball.radius <= 0 {
		ball.Bounce(touchTop)
	}
	if ball.x+ball.radius >= float64(grid.Height) {
		ball.Bounce(touchBottom)
	}
	if ball.y-ball.radius <= 0 {
		ball.Bounce(touchLeft)
	}
	if ball.y+ball.radius >= float64(grid.Width) {
		ball.Bounce(touchRight)
	}
}
