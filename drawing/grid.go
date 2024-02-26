package drawing

type Grid struct {
	pixels []*Pixel

	Width  int
	Height int
}

func NewGrid(width int, height int) *Grid {
	array := make([]*Pixel, width*height)
	for i := 0; i < len(array); i++ {
		array[i] = &Pixel{
			r: 0,
			g: 255,
			b: 0,
			a: 255,
		}
	}
	return &Grid{
		pixels: array,
		Width:  width,
		Height: height,
	}
}

func (g *Grid) Fill(pixel *Pixel) {
	for x := 0; x < g.Height; x++ {
		for y := 0; y < g.Width; y++ {
			g.SetPixel(x, y, pixel)
		}
	}
}

func (g *Grid) SetPixel(x int, y int, pixel *Pixel) {
	if x < 0 || y < 0 || x >= g.Height || y >= g.Width {
		// clamp pixels when trying to draw out of bounds
		return
	}
	g.pixels[g.Height*y+x] = pixel
}

func (g *Grid) GetPixelArray() []byte {
	pixelArray := make([]byte, g.Width*g.Height*4)
	for i, pixel := range g.pixels {
		pixelArray[i*4] = pixel.r
		pixelArray[i*4+1] = pixel.g
		pixelArray[i*4+2] = pixel.b
		pixelArray[i*4+3] = pixel.a
	}
	return pixelArray
}
