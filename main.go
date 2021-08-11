package main

import (
	_ "image/png"

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

	i := 0

	rect := pixel.R(0, 0, 10, 10)

	for !win.Closed() {

		if i % 10 == 0 {
			win.Clear(colornames.Darkslategrey)

			hole.DrawObject(
				i % 3,
				win,
				pixel.IM.
					Moved(win.Bounds().Center()),
			)

			planet.DrawObject(
				0,
				win,
				pixel.IM.
					Rotated(pixel.ZV, float64(i / 2)).
					Moved(pixel.ZV.Add(pixel.V(float64(i), float64(i)))),
			)

			rect = rect.
		}

		i++

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
