package main

// import with "github.com/lotoussa/PixelAnimation/pkg"

import (
	"fmt"
	_ "image/png"
	"time"

	"./pkg/camera"
	"./pkg/sprite"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func gameLoop(
	win *pixelgl.Window,
	cam *camera.Camera,
	hole *sprite.Hole,
	planet *sprite.Planet,
	config pixelgl.WindowConfig,
	basicTxt *text.Text,
) error {
	last := time.Now()
	dynamicDt := 0.0

	for !win.Closed() {

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
			planet.AddPlanet(win.MousePosition(), *cam)
		}

		planet.DrawBatch(win, dynamicDt)

		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1.45))

		cam.Move(win, dt*500)

		win.Update()

		cam.PrintFps(win, config)
	}
	return nil
}

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

	_, err = fmt.Fprintf(basicTxt, "Camera movement: Arrows\n"+
		"Camera zoom: Mouse scroll\n"+
		"Camera reset: Key R\n"+
		"Place Objects: Left click\n")
	if err != nil {
		panic(err)
	}

	cam := camera.NewCamera()

	hole, err := sprite.NewHole()
	if err != nil {
		panic(err)
	}

	planet, err := sprite.InitPlanet()
	if err != nil {
		panic(err)
	}

	err = gameLoop(win, cam, hole, planet, config, basicTxt)
	if err != nil {
		panic(err)
	}
}

func main() {
	pixelgl.Run(run)
}
