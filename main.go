package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tired-1/balls"
)

func main() {
	if err := ebiten.RunGame(balls.NewGame(800, 800)); err != nil {
		log.Fatal(err)
	}
}

func initWindow() {
	const height, width, title = 800, 800, "I'm sorry"
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
}
