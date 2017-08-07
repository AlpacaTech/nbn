package main

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	MAX_BALLS  = 50
	NET_SCALE  = 1
	BOT_SCALE  = .7
	BALL_SCALE = 5
)

var (
	game *Game

	bot, net *pixel.Sprite
	balls    []*Ball
	angle    *imdraw.IMDraw

	score int = 0
)

func main() {
	pixelgl.Run(run)
	fmt.Println(score)
}

func run() {
	game = NewGame("Nothing But Net", pixel.R(0, 0, 1280, 768), true, true, time.Second*45)
	center := game.Win.Bounds().Center()

	loadObjects()

	pos := pixel.IM.Moved(center)

	netPos := pos.Scaled(center, NET_SCALE)
	netPos = netPos.Moved(pixel.V(center.X-game.Pictures[1].Bounds().Max.X*.5-5, -center.Y+game.Pictures[1].Bounds().Max.Y*.5))

	botPos := pos.Scaled(center, BOT_SCALE)
	botPos = botPos.ScaledXY(center, pixel.V(-1, 1))
	botPos = botPos.Moved(pixel.V(-center.X+game.Pictures[0].Bounds().Max.X*.5+5, -center.Y+game.Pictures[0].Bounds().Max.Y*.5-20))
	// botPos = botPos.Moved(pixel.V(bot.Frame().W(), bot.Frame().H()))
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
		}
		if game.Win.Pressed(pixelgl.KeyRight) {
			if botVel < 6.3 {
				botVel += .7
			}
		}
		if !game.Win.Pressed(pixelgl.KeyRight) && !game.Win.Pressed(pixelgl.KeyLeft) {
			if botVel >= .175 {
				botVel -= .175
			} else if botVel <= -.175 {
				botVel += .175
			}
		}

		if game.Win.Pressed(pixelgl.KeyUp) {
			a += .02
		}
		if game.Win.Pressed(pixelgl.KeyDown) {
			a -= .02
		}

		botPos = botPos.Moved(pixel.V(float64(int(botVel)), 0))
		in := game.Win.Bounds().Contains(botPos.Project(bot.Frame().Center()))
		// in = in && game.Win.Bounds().Contains(botPos.Project(bot.Frame().Center()).Add(pixel.V(bot.Frame().W()*.6, 0)))
		in = in && net.Frame().Intersect(bot.Frame()).Min.X == 0
		// in = in && !net.Frame().Contains(botPos.Project(bot.Frame().Max))

		if !in {
			botPos = botPos.Moved(pixel.V(float64(int(-botVel)), 0))
		}

		game.Win.Clear(colornames.Lightslategray)
		net.Draw(game.Win, netPos)
		drawAngle(botPos.Project(bot.Frame().Center()), a)

		for i, ball := range balls {
			upperScored := pixel.R(1160, 275, 1270, 355).Contains(ball.Pos) && ball.Pos.Y < netLine(ball.Pos.X, true)
			lowerScored := pixel.R(1040, 30, 1270, 280).Contains(ball.Pos) && ball.Pos.Y < netLine(ball.Pos.X, false)
			max := game.Win.Bounds().Contains(ball.Pos.Add(pixel.V(-8, -8)))
			min := game.Win.Bounds().Contains(ball.Pos.Add(pixel.V(8, 8)))
			if (max || min) && !upperScored && !lowerScored {
				ball.Draw(game.Win)
			} else {
				func() {
					defer func() {
						if e := recover(); e != nil {
							balls = balls
						}
					}()
					balls = append(balls[:i], balls[i+1:]...)
				}()
				if upperScored {
					score += 5
				} else if lowerScored {
					score += 1
				}
			}
		}

		bot.Draw(game.Win, botPos)

		game.Loop()
	}
}

func loadObjects() {
	f, err := game.LoadPicture("./resource/flywheel.png")
	if err != nil {
		panic(err)
	}

	n, err := game.LoadPicture("./resource/net.png")
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

	angle.Push(pixel.V(50, math.Abs(a*20)+400), pixel.V(50, 400))

	angle.Color = colornames.Darkred
	angle.SetColorMask(colornames.Darkred)

	angle.Line(9)
	angle.Draw(game.Win)
}

func netLine(x float64, upper bool) float64 {
	if upper {
		return 0.69565217391*x - 535.86956521
	}
	return 0.4347826086956521739130*x - 422.1739130434782608695
}
