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
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var (
	window *pixelgl.Window

	dt      float64
	dtList  [100]float64
	dtIndex int
	dtSum   float64

	invaders *Invaders
	player   *Player
	atlas    *text.Atlas
	posText  *text.Text
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

	config.Monitor = pixelgl.PrimaryMonitor()
	if window, err = pixelgl.NewWindow(config); err != nil {
		panic(err)
	}

	window.SetCursorVisible(false)

	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	posText = text.New(pixel.V(0, 0), atlas)

	levelBackground, err := loadPicture("./assets/stars.png")

	if err != nil {
		panic(err)
	}

	levelBackgroundSprite := pixel.NewSprite(levelBackground, levelBackground.Bounds())
	invaders = NewInvaders(window)
	player = NewPlayer(window)

	lastTick := time.Now()

	for !window.Closed() {
		checkForQuit()

		dt = time.Since(lastTick).Seconds()
		lastTick = time.Now()

		levelBackgroundSprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
		invaders.Move(dt)
		checkForPlayerMovement()

		invaders.Draw()
		player.Draw()

		playerPos := player.GetPosition()
		posText.Clear()
		fmt.Fprintf(posText, "X=%f, Y=%f", playerPos.X, playerPos.Y)
		posText.Draw(window, pixel.IM.Moved(pixel.V(10, 10)))

		window.Update()
	}
}

func checkForQuit() {
	if window.JustPressed(pixelgl.KeyQ) {
		window.SetClosed(true)
	}
}

func checkForPlayerMovement() {
	if window.Pressed(pixelgl.KeyLeft) {
		player.MoveLeft(dt)
	}

	if window.Pressed(pixelgl.KeyRight) {
		player.MoveRight(dt)
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
