package camera

import (
	"fmt"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	Pos         pixel.Vec
	Speed       float64
	Zoom        float64
	InverseZoom float64
	ZoomSpeed   float64
	Frames      int
	Second      <-chan time.Time
	Cam         pixel.Matrix
}

func NewCamera() Camera {
	return Camera{
		Pos:         pixel.ZV,
		Speed:       0.72,
		Zoom:        1.0,
		InverseZoom: 1.0,
		ZoomSpeed:   1.01,
		Frames:      0,
		Second:      time.Tick(time.Second),
	}
}

func (c *Camera) Move(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		c.Pos.X -= c.Speed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		c.Pos.X += c.Speed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		c.Pos.Y -= c.Speed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		c.Pos.Y += c.Speed * dt
	}
	c.zoom(win)
	c.reset(c, win)
}

func (c *Camera) zoom(win *pixelgl.Window) {
	c.Zoom *= math.Pow(c.ZoomSpeed, win.MouseScroll().Y)
	c.InverseZoom /= math.Pow(c.ZoomSpeed, win.MouseScroll().Y)
}

func (c Camera) reset(currentCam *Camera, win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyR) {
		*currentCam = NewCamera()
	}
}

func (c *Camera) PrintFps(win *pixelgl.Window, config pixelgl.WindowConfig) {
	c.Frames++
	select {
	case <-c.Second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", config.Title, c.Frames))
		c.Frames = 0
	default:
	}
}
