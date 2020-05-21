package utils

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
)

func FileToImageSlice(f string, w, h int) ([]*ebiten.Image, error) {
	sheet, _, err := ebitenutil.NewImageFromFile(f, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	var images []*ebiten.Image
	sw, sh := sheet.Size()
	for x := 0; x < sw/w; x++ {
		for y := 0; y < sh/h; y++ {
			images = append(images, sheet.SubImage(image.Rect(x*w, y*h, x*w+w, y*h+h)).(*ebiten.Image))
		}
	}

	return images, err
}
