package main

// import with "github.com/lotoussa/PixelAnimation/pkg"

import (
	_ "image/png"
	"time"

	"../PixelAnimationGo/pkg"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	config := pixelgl.WindowConfig{
		Title:  "Pixel Animation",
		Bounds: pixel.R(0, 0, 768, 768),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	hole := pkg.NewObject([]string{"holeA.png", "holeB.png", "holeC.png"})
	planet := pkg.NewObject([]string{"planetA.png"})

	last := time.Now()
	dynamicDt := 0.0

	for !win.Closed() {

		win.Clear(colornames.Darkslategrey)

		dt := time.Since(last).Seconds()
		last = time.Now()

		dynamicDt += 3 * dt

		hole.DrawObject(
			int(dynamicDt * 2) % 3,
			win,
			pixel.IM.
				Moved(win.Bounds().Center()),
		)

		planet.DrawObject(
			0,
			win,
			pixel.IM.
				Rotated(pixel.ZV, dynamicDt).
				Moved(pixel.ZV.Add(pixel.V(dynamicDt * 10, dynamicDt * 10))),
		)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
