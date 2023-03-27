package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	image        *ebiten.Image
	imageOptions *ebiten.DrawImageOptions
	x            float64
	y            float64
}

func NewPlayer() (*Player, error) {
	img, _, err := ebitenutil.NewImageFromFile("./assets/player.png")
	if err != nil {
		return nil, err
	}

	return &Player{
		image:        img,
		imageOptions: &ebiten.DrawImageOptions{},
	}, err
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.imageOptions.GeoM.Reset()
	p.imageOptions.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, p.imageOptions)
}

func (p *Player) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.x += 8
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.x -= 8
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.y -= 8
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.y += 8
	}
	return nil
}
