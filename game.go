package main

import (
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	Title    string
	Win      *pixelgl.Window
	Ticker   *time.Ticker
	Pictures []pixel.Picture
	Over     bool
}

func NewGame(title string, bounds pixel.Rect, resizable, smooth bool, length time.Duration) *Game {
	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    bounds,
		Resizable: resizable,
		VSync:     smooth,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(smooth)

	g := &Game{
		Title:  title,
		Win:    win,
		Ticker: time.NewTicker(time.Second / 60),
		Over:   false,
	}
	go g.Wait(length)
	return g
}

func (g *Game) LoadPicture(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, err
	}
	g.Pictures = append(g.Pictures, pixel.PictureDataFromImage(img))
	return len(g.Pictures) - 1, nil
}

func (g *Game) Open() bool {
	return !g.Win.Closed() && !g.Over
}

func (g *Game) Loop() {
	g.Win.Update()
	<-g.Ticker.C
}

func (g *Game) Wait(length time.Duration) {
	time.Sleep(length)
	g.Over = true
}
