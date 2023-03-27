package game

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenaudio "github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sergiosegrera/jam/audio"
	"github.com/sergiosegrera/jam/entities"
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
	entities        []GameObject
	keySound        *ebitenaudio.Player
	backgroundMusic *ebitenaudio.Player
}

func New() (*Game, error) {
	player, err := entities.NewPlayer()
	if err != nil {
		return nil, err
	}

	slime, err := entities.NewSlime(32, 32)
	if err != nil {
		return nil, err
	}

	audioContext := ebitenaudio.NewContext(44100)

	keySound, err := audio.NewSound(audioContext, "./assets/sounds/keypress.wav")
	if err != nil {
		return nil, err
	}

	backgroundMusic, err := audio.NewMusic(audioContext, "./assets/music/ouverture1812.mp3")
	if err != nil {
		return nil, err
	}
	backgroundMusic.Play()

	return &Game{
		entities: []GameObject{
			player,
			slime,
		},
		keySound:        keySound,
		backgroundMusic: backgroundMusic,
	}, err
}

func (g *Game) Update() error {
	// TODO: Escape for quick debugging, remove in production
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("escape called, game exited")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) && !g.keySound.IsPlaying() {
		g.keySound.Rewind()
		g.keySound.Play()
	}

	for _, entity := range g.entities {
		entity.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BLACK)

	for _, entity := range g.entities {
		entity.Draw(screen)
	}

	// TODO: TPS counter, remove in production
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f, FPS: %0.2f", ebiten.ActualTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 128, 128
}
