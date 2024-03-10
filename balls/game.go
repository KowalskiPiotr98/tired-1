package balls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"os"
	"tired-1/drawing"
	"tired-1/objects"
)

type Game struct {
	grid  *drawing.Grid
	balls []*objects.Ball
	tiles []*objects.Tile
}

func NewGame(width int, height int) *Game {
	objectSize := 50
	grid := drawing.NewGrid(width, height)
	balls := []*objects.Ball{
		objects.NewBall(float64(width)/4, float64(height)/2, float64(objectSize)/2, drawing.PixelRed, drawing.PixelBlue),
		objects.NewBall(float64(width)/4*3, float64(height)/2, float64(objectSize)/2, drawing.PixelBlue, drawing.PixelRed),
	}
	tiles := make([]*objects.Tile, 0)
	for i := 0; i <= height; i += objectSize {
		for j := 0; j <= width; j += objectSize {
			tiles = append(tiles, objects.NewTile(float64(i), float64(j), float64(objectSize), drawing.PixelRed))
		}
	}
	return &Game{
		grid:  grid,
		balls: balls,
		tiles: tiles,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	for _, ball := range g.balls {
		ball.Move()
	}
	for _, ball := range g.balls {
		objects.BounceWall(ball, g.grid)
	}
	for _, tile := range g.tiles {
		for _, ball := range g.balls {
			tile.HandleBallTouch(ball)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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
