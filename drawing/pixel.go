package drawing

var (
	PixelGreen = NewPixel(0, 255, 0, 255)
	PixelRed   = NewPixel(255, 0, 0, 255)
)

type Pixel struct {
	r byte
	g byte
	b byte
	a byte
}

func NewPixel(r byte, g byte, b byte, a byte) *Pixel {
	return &Pixel{
		r: r,
		g: g,
		b: b,
		a: a,
	}
}
