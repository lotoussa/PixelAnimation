package sprites

import (
	"path"

	"../camera"
	"github.com/faiface/pixel"
)

type Planet struct {
	Sprites  []*pixel.Sprite
	Matrices []pixel.Matrix
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
		pixel.IM.Moved(mouse),
	)
}
