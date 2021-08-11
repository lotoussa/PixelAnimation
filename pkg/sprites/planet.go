package sprites

import (
	"path"

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

func (p *Planet) AddPlanet(mousePosition pixel.Vec) {
	p.Sprites = append(
		p.Sprites,
		pixel.NewSprite(p.pic, p.pic.Bounds()),
	)
	p.Matrices = append(
		p.Matrices,
		pixel.IM.Moved(mousePosition),
	)
}
