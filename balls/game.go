package balls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tired-1/drawing"
)

type Game struct {
	grid *drawing.Grid
}

func NewGame(width int, height int) *Game {
	return &Game{
		grid: drawing.NewGrid(width, height),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.grid.GetPixelArray())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.grid.Width, g.grid.Height
}
