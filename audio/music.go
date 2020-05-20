package audio

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"os"
)

func NewMusic(ctx *audio.Context, f string) (*audio.Player, error) {
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

	return audioPlayer, err
}
