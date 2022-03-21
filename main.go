package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/sergiosegrera/jam/game"
)

func main() {
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle("sergiosegrera/jam")

	_, icon, _ := ebitenutil.NewImageFromFile("./assets/icon128.png")
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
