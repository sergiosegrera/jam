package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Player struct {
	image     *ebiten.Image
	x         float64
	y         float64
	moveTimer int
}

func NewPlayer() (*Player, error) {
	img, _, err := ebitenutil.NewImageFromFile("./assets/player.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return &Player{
		image: img,
	}, err
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, options)
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
