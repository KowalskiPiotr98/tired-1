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
			r: byte(i % 255),
			g: byte(i % 255),
			b: byte(i % 255),
			a: byte(i % 255),
		}
	}
	return &Grid{
		pixels: array,
		Width:  width,
		Height: height,
	}
}

func (g *Grid) GetPixelArray() []byte {
	pixelArray := make([]byte, g.Width*g.Height*4)
	i := 0
	for w := 0; w < g.Width; w++ {
		for h := 0; h < g.Height; h++ {
			pixel := g.pixels[w*h+h]
			pixelArray[i] = pixel.r
			pixelArray[i+1] = pixel.g
			pixelArray[i+2] = pixel.b
			pixelArray[i+3] = pixel.a
			i = i + 4
		}
	}
	return pixelArray
}
