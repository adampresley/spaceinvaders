package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"time"

	_ "image/png"

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

	fmt.Printf("Window size: %0.1fx%0.1f\n", window.Bounds().W(), window.Bounds().H())

	lastTick := time.Now()
	fps = 0
	frames = 0
	second = time.Tick(time.Second)

	for !window.Closed() {
		dt = time.Since(lastTick).Seconds()
		lastTick = time.Now()

		game.CheckForQuit()
		game.MoveInvaders(dt)
		game.CheckForPlayerMovement(dt)

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

/*
loadPicture loads an image file into a picture struct
*/
func loadPicture(path string) (pixel.Picture, error) {
	var err error
	var file *os.File
	var img image.Image

	if file, err = os.Open(path); err != nil {
		return nil, err
	}

	defer file.Close()

	if img, _, err = image.Decode(file); err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
