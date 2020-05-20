package audio

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"os"
)

func NewSound(ctx *audio.Context, f string) (*audio.Player, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	stream, err := wav.Decode(ctx, file)
	if err != nil {
		return nil, err
	}

	audioPlayer, err := audio.NewPlayer(ctx, stream)
	if err != nil {
		return nil, err
	}

	return audioPlayer, err
}
