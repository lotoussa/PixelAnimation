package sprites

import (
	"path"

	"../camera"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Planet struct {
	Sprites  []*pixel.Sprite
	Matrices []pixel.Matrix
	Batch    *pixel.Batch
	pic      pixel.Picture
}

func InitPlanet() Planet {
	planet := Planet{}

	pic, err := LoadPicture(path.Join("assets", "planetA.png"))
	if err != nil {
		panic(err)
	}

	planet.Sprites = nil
	planet.Matrices = nil
	planet.Batch = pixel.NewBatch(&pixel.TrianglesData{}, pic)
	planet.pic = pic

	return planet
}

func (p *Planet) AddPlanet(mousePosition pixel.Vec, cam camera.Camera) {
	p.Sprites = append(
		p.Sprites,
		pixel.NewSprite(p.pic, p.pic.Bounds()),
	)
	mouse := cam.Cam.Unproject(mousePosition)
	p.Matrices = append(
		p.Matrices,
		pixel.IM.Scaled(pixel.ZV, cam.InverseCamZoom).Moved(mouse),
	)
}

func (p *Planet) DrawBatch(win *pixelgl.Window, dynamicDt float64) {
	p.Batch.Clear()
	for i, planet := range p.Sprites {
		planet.Draw(p.Batch, p.Matrices[i].
			Rotated(win.Bounds().Center(), dynamicDt))
		//Moved(pixel.ZV.Add(pixel.V(dynamicDt*10, dynamicDt*10))),

	}
	p.Batch.Draw(win)
}
