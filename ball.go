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

func NewBall(vel, acc float64, pos pixel.Vec) (b *Ball) {
	b = new(Ball)
	b.Shape = imdraw.New(nil)
	b.Vel = vel
	b.Acc = acc

	t := pixel.IM.Moved(pixel.V(100, -25))
	b.Pos = t.Project(pos)
	b.Start = b.Pos
	return
}

func (B *Ball) Draw(w *pixelgl.Window) {
	B.Shape.Clear()
	B.Shape.Reset()

	B.Calc()
	B.Shape.Push(B.Pos)

	B.Shape.Color = colornames.Lightyellow
	B.Shape.SetColorMask(colornames.Greenyellow)

	B.Shape.Circle(8, 0)
	B.Shape.Draw(w)
}

func (B *Ball) Calc() {
	x := B.Pos.X + B.Vel
	y := B.Start.Y + (B.Pos.X-B.Start.X)*B.Vel
	B.Pos = pixel.V(x*B.Acc, y*B.Acc)
}
