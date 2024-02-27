package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tired-1/balls"
)

const (
	cellHalfSize   = 25
	fieldSideCount = 10
	playFieldSide  = fieldSideCount * cellHalfSize * 2
)

func main() {
	initWindow()
	if err := ebiten.RunGame(balls.NewGame(playFieldSide, playFieldSide)); err != nil {
		log.Fatal(err)
	}
}

func initWindow() {
	const height, width, title = playFieldSide, playFieldSide, "I'm sorry"
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
}
