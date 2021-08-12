package camera

import (
	"fmt"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	CamPos         pixel.Vec
	CamSpeed       float64
	CamZoom        float64
	InverseCamZoom float64
	CamZoomSpeed   float64
	Frames         int
	Second         <-chan time.Time
	Cam            pixel.Matrix
}

func NewCamera() Camera {
	return Camera{
		CamPos:         pixel.ZV,
		CamSpeed:       0.72,
		CamZoom:        1.0,
		InverseCamZoom: 1.0,
		CamZoomSpeed:   1.01,
		Frames:         0,
		Second:         time.Tick(time.Second),
	}
}

func (c *Camera) Move(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		c.CamPos.X -= c.CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		c.CamPos.X += c.CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		c.CamPos.Y -= c.CamSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		c.CamPos.Y += c.CamSpeed * dt
	}
	c.zoom(win)
	c.reset(c, win)
}

func (c *Camera) zoom(win *pixelgl.Window) {
	c.CamZoom *= math.Pow(c.CamZoomSpeed, win.MouseScroll().Y)
	c.InverseCamZoom /= math.Pow(c.CamZoomSpeed, win.MouseScroll().Y)
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
