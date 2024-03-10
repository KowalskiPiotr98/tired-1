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
	grid := drawing.NewGrid(width, height)
	balls := []*objects.Ball{
		objects.NewBall(100, 100, 10, drawing.PixelRed),
	}
	tiles := []*objects.Tile{
		objects.NewTile(0, 200, 20, drawing.PixelGreen),
		objects.NewTile(20, 200, 20, drawing.PixelGreen),
		objects.NewTile(40, 200, 20, drawing.PixelGreen),
		objects.NewTile(60, 200, 20, drawing.PixelGreen),
		objects.NewTile(80, 200, 20, drawing.PixelGreen),
		objects.NewTile(100, 200, 20, drawing.PixelGreen),
		objects.NewTile(120, 200, 20, drawing.PixelGreen),
		objects.NewTile(140, 200, 20, drawing.PixelGreen),
		objects.NewTile(160, 200, 20, drawing.PixelGreen),
		objects.NewTile(180, 200, 20, drawing.PixelGreen),
		objects.NewTile(200, 200, 20, drawing.PixelGreen),
		objects.NewTile(220, 200, 20, drawing.PixelGreen),
		objects.NewTile(240, 200, 20, drawing.PixelGreen),
		objects.NewTile(260, 200, 20, drawing.PixelGreen),
		objects.NewTile(280, 200, 20, drawing.PixelGreen),
		objects.NewTile(300, 200, 20, drawing.PixelGreen),
		objects.NewTile(320, 200, 20, drawing.PixelGreen),
		objects.NewTile(340, 200, 20, drawing.PixelGreen),
		objects.NewTile(360, 200, 20, drawing.PixelGreen),
		objects.NewTile(380, 200, 20, drawing.PixelGreen),
		objects.NewTile(400, 200, 20, drawing.PixelGreen),
		objects.NewTile(200, 100, 20, drawing.PixelGreen),
		objects.NewTile(200, 120, 20, drawing.PixelGreen),
		objects.NewTile(200, 140, 20, drawing.PixelGreen),
		objects.NewTile(200, 160, 20, drawing.PixelGreen),
		objects.NewTile(200, 180, 20, drawing.PixelGreen),
		objects.NewTile(200, 200, 20, drawing.PixelGreen),
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
	g.grid.Fill(drawing.PixelBlue)
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
