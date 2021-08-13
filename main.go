package main

// import with "github.com/lotoussa/PixelAnimation/pkg"

import (
	"fmt"
	_ "image/png"
	"time"

	"../PixelAnimationGo/pkg/camera"
	"../PixelAnimationGo/pkg/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func run() {

	config := pixelgl.WindowConfig{
		Title:  "Pixel Animation",
		Bounds: pixel.R(0, 0, 1280, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(20, 20), basicAtlas)
	basicTxt.LineHeight = basicAtlas.LineHeight() * 1.3
	basicTxt.Color = colornames.Violet

	fmt.Fprintf(basicTxt, "Camera movement: Arrows\n"+
		"Camera zoom: Mouse scroll\n"+
		"Camera reset: Key R\n"+
		"Place Objects: Left click\n")

	last := time.Now()
	dynamicDt := 0.0

	cam := camera.NewCamera()

	hole := sprites.NewHole()
	planet := sprites.InitPlanet()

	for !win.Closed() {

		// try to move only the hole
		cam.Cam = pixel.IM.Scaled(
			win.Bounds().Center(), cam.Zoom).
			Moved(pixel.ZV.Sub(cam.Pos))
		win.SetMatrix(cam.Cam)

		win.Clear(colornames.Black)

		dt := time.Since(last).Seconds()
		last = time.Now()

		dynamicDt += 3 * dt

		hole.Sprites[int(dynamicDt*2)%3].Draw(
			win,
			pixel.IM.Moved(win.Bounds().Center()),
		)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			planet.AddPlanet(win.MousePosition(), cam)
		}

		planet.DrawBatch(win, dynamicDt)

		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1.45))

		cam.Move(win, dt*500)

		win.Update()

		cam.PrintFps(win, config)
	}
}

func main() {
	pixelgl.Run(run)
}
