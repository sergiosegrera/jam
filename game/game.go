package game

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	ebitenaudio "github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	player          GameObject
	slime           GameObject
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

	audioContext, err := ebitenaudio.NewContext(44100)
	if err != nil {
		return nil, err
	}

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
		player:          player,
		slime:           slime,
		keySound:        keySound,
		backgroundMusic: backgroundMusic,
	}, err
}

func (g *Game) Update(image *ebiten.Image) error {
	// TODO: Escape for quick debugging, remove in production
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Escape called, game exited")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) && !g.keySound.IsPlaying() {
		g.keySound.Rewind()
		g.keySound.Play()
	}

	g.player.Update()
	g.slime.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BLACK)
	g.player.Draw(screen)
	g.slime.Draw(screen)
	// TODO: TPS counter, remove in production
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outseideHeight int) (int, int) {
	return 128, 128
}
