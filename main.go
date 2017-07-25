package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	game    *Game
	gophers []*pixel.Sprite
)

func main() {
	pixelgl.Run(run)
}

func run() {
	game = NewGame("Stay Alive", pixel.R(0, 0, 1280, 768), false)
	center := game.Win.Bounds().Center()

	loadGophers()

	animate := time.Now().Add(time.Millisecond * 250)
	for blink := false; game.Open(); {
		game.Loop()
		game.Win.Clear(colornames.Cornflowerblue)

		pos := pixel.IM.Moved(center).Scaled(center, 5)
		if blink {
			gophers[1].Draw(game.Win, pos)

			if time.Now().After(animate) {
				blink = !blink
				animate = time.Now().Add(time.Millisecond * 250)
			}
		} else {
			gophers[0].Draw(game.Win, pos)

			if time.Now().After(animate) {
				blink = !blink
				animate = time.Now().Add(time.Millisecond * 50)
			}
		}
	}
}

func loadGophers() {
	spritesheet, err := game.LoadPicture("./img/gopher.png")
	if err != nil {
		panic(err)
	}

	var l, h float64 = 12, 14
	for i := 0.0; i < 27; i++ {
		gophers = append(gophers, pixel.NewSprite(game.Pictures[spritesheet], pixel.R(12*i, 0, 12*i+l, h)))
	}
}
