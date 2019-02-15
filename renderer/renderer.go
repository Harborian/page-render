package renderer

import (
	"io"
)

// Renderer interface
type Renderer interface {
	Render(url string) (io.Reader, error)
}
