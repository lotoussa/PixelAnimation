package camera

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	CamPos       pixel.Vec
	CamSpeed     float64
	CamZoom      float64
	CamZoomSpeed float64
	Cam          pixel.Matrix
}

func NewCamera() Camera {
	return Camera{
		CamPos:       pixel.ZV,
		CamSpeed:     0.5,
		CamZoom:      1.0,
		CamZoomSpeed: 1.01,
	}
}

func (c *Camera) Move(win *pixelgl.Window, dynamicDt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		c.CamPos.X += c.CamSpeed * dynamicDt
	}
	if win.Pressed(pixelgl.KeyRight) {
		c.CamPos.X -= c.CamSpeed * dynamicDt
	}
	if win.Pressed(pixelgl.KeyDown) {
		c.CamPos.Y += c.CamSpeed * dynamicDt
	}
	if win.Pressed(pixelgl.KeyUp) {
		c.CamPos.Y -= c.CamSpeed * dynamicDt
	}
	c.zoom(win)
}

// deprecated
func (c *Camera) zoom(win *pixelgl.Window) {
	c.CamZoom *= math.Pow(c.CamZoomSpeed, win.MouseScroll().Y)
}
