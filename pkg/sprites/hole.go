package sprites

import (
	"path"

	"github.com/faiface/pixel"
)

type Hole struct {
	Sprites []*pixel.Sprite
}

func NewHole() Hole {
	hole := Hole{}

	for _, picPath := range []string{"holeA.png", "holeB.png", "holeC.png"} {
		pic, err := LoadPicture(path.Join("assets", picPath))
		if err != nil {
			panic(err)
		}
		hole.Sprites = append(hole.Sprites, pixel.NewSprite(pic, pic.Bounds()))
	}

	return hole
}
