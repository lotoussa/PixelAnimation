package main

// import with "github.com/lotoussa/PixelAnimation/pkg"

import (
	_ "image/png"
	"time"

	"../PixelAnimationGo/pkg/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	config := pixelgl.WindowConfig{
		Title:  "Pixel Animation",
		Bounds: pixel.R(0, 0, 768, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	hole := sprites.NewHole()
	planet := sprites.InitPlanet()

	last := time.Now()
	dynamicDt := 0.0

	for !win.Closed() {

		win.Clear(colornames.Darkslateblue)

		dt := time.Since(last).Seconds()
		last = time.Now()

		dynamicDt += 3 * dt

		hole.Sprites[int(dynamicDt*2)%3].Draw(
			win,
			pixel.IM.Moved(win.Bounds().Center()),
		)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			planet.AddPlanet(win.MousePosition())
		}

		for i, p := range planet.Sprites {
			p.Draw(win, planet.Matrices[i])
			//Rotated(pixel.ZV, dynamicDt).
			//Moved(pixel.ZV.Add(pixel.V(dynamicDt*10, dynamicDt*10))),

		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
