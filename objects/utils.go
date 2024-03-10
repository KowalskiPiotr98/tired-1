package objects

import (
	"math"
	"math/rand/v2"
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
	if ball.x+ball.radius >= tile.x-tile.side/2 && math.Abs(ball.y-tile.y) <= ball.radius && ball.movX > 0 {
		return touchRight
	}
	if ball.x-ball.radius <= tile.x+tile.side/2 && math.Abs(ball.y-tile.y) <= ball.radius && ball.movX < 0 {
		return touchLeft
	}
	if ball.y-ball.radius <= tile.y+tile.side/2 && math.Abs(ball.x-tile.x) <= ball.radius && ball.movY < 0 {
		return touchTop
	}
	if ball.y+ball.radius >= tile.y-tile.side/2 && math.Abs(ball.x-tile.x) <= ball.radius && ball.movY > 0 {
		return touchBottom
	}
	return touchNone
}

func BounceWall(ball *Ball, grid *drawing.Grid) {
	if ball.x-ball.radius <= 0 {
		ball.Bounce(touchLeft)
	}
	if ball.x+ball.radius >= float64(grid.Height) {
		ball.Bounce(touchRight)
	}
	if ball.y-ball.radius <= 0 {
		ball.Bounce(touchTop)
	}
	if ball.y+ball.radius >= float64(grid.Width) {
		ball.Bounce(touchBottom)
	}
}

func getRandomMovementVector(length float64) (movX, movY float64) {
	movX = rand.Float64() * length
	movY = math.Sqrt(length*length - movX*movX)
	return
}
