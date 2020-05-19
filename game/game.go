package game

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/sergiosegrera/jam/player"
)

var (
	BLACK = color.RGBA{0x00, 0x00, 0x00, 0xff}
	WHITE = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

type GameObject interface {
	Update() error
	Draw(*ebiten.Image)
}

type Game struct {
	player GameObject
}

func New() (*Game, error) {
	p, err := player.New()
	if err != nil {
		return nil, err
	}
	return &Game{
		player: p,
	}, err
}

func (g *Game) Update(image *ebiten.Image) error {
	// TODO: Escape for quick debugging, remove in production
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Escape called, game exited")
	}
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BLACK)
	g.player.Draw(screen)
	// TODO: TPS counter, remove in production
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outseideHeight int) (int, int) {
	return 128, 128
}
