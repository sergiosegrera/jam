package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/sergiosegrera/jam/game"
	"image"
	"log"
)

func main() {
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle("sergiosegrera/jam")

	_, icon, _ := ebitenutil.NewImageFromFile("./assets/icon.png", ebiten.FilterDefault)
	ebiten.SetWindowIcon([]image.Image{icon})

	game, err := game.New()
	if err != nil {
		log.Panic(err)
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Println(err)
	}
}
