package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/sergiosegrera/jam/utils"
)

type Slime struct {
	images []*ebiten.Image
	x      float64
	y      float64
	frame  int
	time   int64
}

func NewSlime(x, y float64) (*Slime, error) {
	images, err := utils.FileToImageSlice("./assets/slime-sheet.png", 8, 8)
	if err != nil {
		return nil, err
	}

	return &Slime{
		images: images,
		x:      x,
		y:      y,
	}, err
}

func NewSlimeWithImage(images []*ebiten.Image, x, y float64) (*Slime, error) {
	// More performant
	return &Slime{
		images: images,
		x:      x,
		y:      y,
	}, nil
}

func (s *Slime) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.images[s.frame], options)
}

func (s *Slime) Update() error {
	if s.time > 60 {
		s.frame++
		if s.frame > len(s.images)-1 {
			s.frame = 0
		}
		s.time = 0
	}
	s.time++
	return nil
}
