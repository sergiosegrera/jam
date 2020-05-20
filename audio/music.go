package audio

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"os"
	"time"
)

type Music struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

func NewMusic(ctx *audio.Context, f string) (*Music, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	stream, err := mp3.Decode(ctx, file)
	if err != nil {
		return nil, err
	}

	// TODO: Fix hard coded length
	loop := audio.NewInfiniteLoop(stream, 251*4*44100)

	audioPlayer, err := audio.NewPlayer(ctx, loop)
	if err != nil {
		return nil, err
	}

	audioPlayer.SetVolume(1)
	return &Music{
		audioContext: ctx,
		audioPlayer:  audioPlayer,
	}, err
}

func (m *Music) Play() {
	m.audioPlayer.Play()
}

func (m *Music) Current() time.Duration {
	return m.audioPlayer.Current()
}
