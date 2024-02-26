package objects

import "tired-1/drawing"

type Object interface {
	// Draw places the object on the [drawing.Grid].
	// Note that z-layering is kind of dumb: objects drawn later will be on top of the ones drawn earlier.
	Draw(grid *drawing.Grid)
}
