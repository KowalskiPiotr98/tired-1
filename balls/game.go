package balls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tired-1/drawing"
	"tired-1/objects"
)

type Game struct {
	grid  *drawing.Grid
	balls []*objects.Ball
	tiles []*objects.Tile
}

func NewGame(width int, height int) *Game {
	balls := []*objects.Ball{
		objects.NewBall(100, 100, 10, drawing.PixelRed),
	}
	tiles := []*objects.Tile{
		objects.NewTile(200, 200, 20, drawing.PixelRed),
	}
	return &Game{
		grid:  drawing.NewGrid(width, height),
		balls: balls,
		tiles: tiles,
	}
}

func (g *Game) Update() error {
	for _, ball := range g.balls {
		ball.Move()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Fill(drawing.PixelGreen)
	for _, tile := range g.tiles {
		tile.Draw(g.grid)
	}
	for _, ball := range g.balls {
		ball.Draw(g.grid)
	}
	screen.WritePixels(g.grid.GetPixelArray())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.grid.Width, g.grid.Height
}
