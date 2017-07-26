package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	MAX_BALLS  = 10
	NET_SCALE  = 1
	BOT_SCALE  = .7
	BALL_SCALE = 5
)

var (
	game *Game

	bot, net *pixel.Sprite
	balls    []Ball
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
	var botVel float64 = 0

	for game.Open() {
		game.Win.Clear(colornames.Lightslategray)

		if game.Win.JustPressed(pixelgl.KeySpace) {
			ball := Ball{
				Shape: imdraw.New(nil),
				Vel:   1,
				Acc:   2,
				Start: pixel.V(10, 10),
				Pos:   pixel.V(10, 10),
			}
			balls = append(balls, ball)
		}

		if game.Win.Pressed(pixelgl.KeyLeft) && bot.Frame().Min.X >= 0 {
			if botVel > -6.3 {
				botVel -= .7
			}
		} else if game.Win.Pressed(pixelgl.KeyRight) && bot.Frame().Max.X <= 1280 {
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

		botPos = botPos.Moved(pixel.V(float64(int(botVel)), 0))

		net.Draw(game.Win, netPos)
		bot.Draw(game.Win, botPos)

		for _, ball := range balls {
			ball.Draw(game.Win)
		}

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
}
