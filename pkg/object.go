package pkg

import (
	"image"
	"os"
	"path"

	"github.com/faiface/pixel"
)

type Object struct {
	Sprite []*pixel.Sprite
}

func NewObject(objPaths []string) Object {
	obj := Object{}

	for _, objPath := range objPaths {
		pic, err := obj.loadPicture(path.Join("assets", objPath))
		if err != nil {
			panic(err)
		}
		obj.Sprite = append(obj.Sprite, pixel.NewSprite(pic, pic.Bounds()))
	}

	return obj
}

func (o *Object) loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (o *Object) DrawObject(index int, t pixel.Target, m pixel.Matrix) {
	o.Sprite[index].Draw(t, m)
}
