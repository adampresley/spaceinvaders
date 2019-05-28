//go:generate esc -o ./assets.go -pkg main -ignore "DS_Store|LICENSE|(.*?).go$|(.*?).md|(.*?).svg" ./assets
package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	window *pixelgl.Window
	game   *Game

	dt     float64
	fps    int
	frames int
	second <-chan time.Time
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var err error
	rand.Seed(time.Now().UnixNano())

	config := pixelgl.WindowConfig{
		Title:  "Space Invaders",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}

	if window, err = pixelgl.NewWindow(config); err != nil {
		panic(err)
	}

	window.SetCursorVisible(false)

	game = NewGame(window)

	lastTick := time.Now()
	fps = 0
	frames = 0
	second = time.Tick(time.Second)

	for !window.Closed() {
		dt = time.Since(lastTick).Seconds()
		lastTick = time.Now()

		game.Update(dt)
		game.Draw()
		window.Update()

		frames++

		select {
		case <-second:
			fps = frames
			frames = 0
		default:
		}
	}
}
