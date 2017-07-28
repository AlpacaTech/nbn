package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Ball struct {
	Shape   *imdraw.IMDraw
	Acc     float64
	Vel     float64
	Start   pixel.Vec
	Pos     pixel.Vec
	Rsource rand.Source
	Rgen    *rand.Rand
}

func NewBall(vel, acc float64, pos pixel.Vec) (b *Ball) {
	b = new(Ball)

	b.Rsource = rand.NewSource(int64(time.Now().Second()))
	b.Rgen = rand.New(b.Rsource)

	b.Shape = imdraw.New(nil)
	b.Vel = vel - .75 + b.Rgen.Float64()
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
	t := B.Pos.X - B.Start.X
	x := B.Start.X + t + B.Vel
	y := B.Start.Y + .9*math.Pow(t, 1.09)
	B.Pos = pixel.V(x+1.007, y)
	B.Start = pixel.V(B.Start.X, B.Start.Y-(0.03*t))
}
