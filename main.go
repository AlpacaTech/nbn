package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	MAX_BALLS  = 100
	NET_SCALE  = 1
	BOT_SCALE  = .7
	BALL_SCALE = 5
)

var (
	game *Game

	bot, net *pixel.Sprite
	balls    []*Ball
	angle    *imdraw.IMDraw
)

func main() {
	pixelgl.Run(run)
}

func run() {
	game = NewGame("Nothing But Net", pixel.R(0, 0, 1280, 768), false, true)
	center := game.Win.Bounds().Center()

	loadObjects()
	pos := pixel.IM.Moved(center)

	netPos := pos.Scaled(center, NET_SCALE)
	netPos = netPos.Moved(pixel.V(center.X-game.Pictures[1].Bounds().Max.X*.5-5, -center.Y+game.Pictures[1].Bounds().Max.Y*.5+25))

	botPos := pos.Scaled(center, BOT_SCALE)
	botPos = botPos.ScaledXY(center, pixel.V(-1, 1))
	botPos = botPos.Moved(pixel.V(-center.X+game.Pictures[0].Bounds().Max.X*.5+5, -center.Y+game.Pictures[0].Bounds().Max.Y*.5+5))
	var botVel, a float64 = 0, 0

	for game.Open() {
		if game.Win.JustPressed(pixelgl.KeySpace) && len(balls) <= MAX_BALLS {
			ball := NewBall(a, 1.05, botPos.Project(bot.Frame().Center()))
			balls = append(balls, ball)
		}

		if game.Win.Pressed(pixelgl.KeyLeft) {
			if botVel > -6.3 {
				botVel -= .7
			}
		} else if game.Win.Pressed(pixelgl.KeyRight) {
			if botVel < 6.3 {
				botVel += .7
			}
		} else {
			if botVel >= .175 {
				botVel -= .175
			} else if botVel <= -.175 {
				botVel += .175
			}
		}

		if game.Win.Pressed(pixelgl.KeyUp) {
			a += .02
		} else if game.Win.Pressed(pixelgl.KeyDown) {
			a -= .02
		}

		botPos = botPos.Moved(pixel.V(float64(int(botVel)), 0))

		game.Win.Clear(colornames.Lightslategray)
		net.Draw(game.Win, netPos)
		drawAngle(botPos.Project(bot.Frame().Center()), a)

		for i, ball := range balls {
			max := game.Win.Bounds().Contains(ball.Pos.Add(pixel.V(-8, -8)))
			min := game.Win.Bounds().Contains(ball.Pos.Add(pixel.V(8, 8)))
			if max || min {
				ball.Draw(game.Win)
			} else {
				ab := i + 1
				if ab >= len(balls) {
					ab--
				}
				balls = append(balls[:i], balls[ab:]...)
			}
		}
		bot.Draw(game.Win, botPos)

		game.Loop()
	}
}

func loadObjects() {
	f, err := game.LoadPicture("./img/flywheel.png")
	if err != nil {
		panic(err)
	}

	n, err := game.LoadPicture("./img/net.png")
	if err != nil {
		panic(err)
	}

	bot = pixel.NewSprite(game.Pictures[f], game.Pictures[f].Bounds())
	net = pixel.NewSprite(game.Pictures[n], game.Pictures[n].Bounds())
	angle = imdraw.New(nil)
}

func drawAngle(botPos pixel.Vec, a float64) {
	angle.Clear()
	angle.Reset()

	// t := pixel.IM.Moved(pixel.V(100, -25))
	// angle.Push(t.Project(botPos), t.Moved(pixel.V(50, a*50)).Project(botPos))
	angle.Push(pixel.V(50, math.Abs(a*20)+400), pixel.V(50, 400))

	angle.Color = colornames.Darkred
	angle.SetColorMask(colornames.Seagreen)

	angle.Line(9)
	angle.Draw(game.Win)
}
