package main

import (
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

	dt      float64
	dtList  [100]float64
	dtIndex int
	dtSum   float64
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var err error

	rand.Seed(time.Now().UnixNano())

	config := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1920, 1080),
		VSync:  true,
	}

	//config.Monitor = pixelgl.PrimaryMonitor()
	if window, err = pixelgl.NewWindow(config); err != nil {
		panic(err)
	}

	window.SetCursorVisible(false)

	levelBackground, err := loadPicture("./assets/stars.png")

	if err != nil {
		panic(err)
	}

	levelBackgroundSprite := pixel.NewSprite(levelBackground, levelBackground.Bounds())
	game = NewGame(window)

	lastTick := time.Now()

	for !window.Closed() {
		dt = time.Since(lastTick).Seconds()
		lastTick = time.Now()

		game.CheckForQuit()

		levelBackgroundSprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
		game.MoveInvaders(dt)
		game.CheckForPlayerMovement(dt)

		game.Draw()
		window.Update()
	}
}

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

func calculateFPS() float64 {
	dtSum -= dtList[dtIndex]
	dtSum += dt
	dtList[dtIndex] = dt

	dtIndex++

	if dtIndex == 100 {
		dtIndex = 0
	}

	return dtSum / 100.0
}
