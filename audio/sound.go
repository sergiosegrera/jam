package audio

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"os"
)

type Sound struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

func NewSound(ctx *audio.Context, f string) (*Sound, error) {
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

	return &Sound{
		audioContext: ctx,
		audioPlayer:  audioPlayer,
	}, err
}

func (s *Sound) Play() {
	s.audioPlayer.Rewind()
	s.audioPlayer.Play()
}

func (s *Sound) IsPlaying() bool {
	return s.audioPlayer.IsPlaying()
}
