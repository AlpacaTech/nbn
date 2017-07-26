package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Ball struct {
	Shape *imdraw.IMDraw
	Acc   float64
	Vel   float64
	Start pixel.Vec
	Pos   pixel.Vec
}

func (B Ball) Draw(w *pixelgl.Window) {
	B.Pos = B.Calc()

	B.Shape.Reset()
	B.Shape.Push(B.Pos)
	B.Shape.Color = colornames.Yellow
	B.Shape.Circle(8, 0)

	B.Shape.Draw(w)
}

func (B Ball) Calc() pixel.Vec {
	// y := B.Acc*(B.Pos.X+B.Vel) + B.Start.Y
	// return pixel.V(B.Pos.X+B.Vel, y)
	return B.Pos.Add(pixel.V(1, 1))
}
