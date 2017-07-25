package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	Title    string
	Win      *pixelgl.Window
	Pictures []pixel.Picture
}

func NewGame(title string, bounds pixel.Rect, resizable bool) *Game {
	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    bounds,
		Resizable: resizable,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &Game{
		Title: title,
		Win:   win,
	}
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

/*
func (g *Game) TransitionWindowBounds(bounds pixel.Rect) {
	b := g.Win.Bounds()
	an int = (math.Abs(b.Max.X - bounds.Max.X) + math.Abs(b.Max.Y - Bounds.Max.Y)) / 2
	for g.Win.Bounds().Size().X {}
}
*/

func (g *Game) Open() bool {
	return !g.Win.Closed()
}

func (g *Game) Loop() {
	g.Win.Update()
}
