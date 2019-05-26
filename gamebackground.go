package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
GameBackground handles the scrolling background during the game
*/
type GameBackground struct {
	window              *pixelgl.Window
	starsTile1          *pixel.Sprite
	starsTile2          *pixel.Sprite
	closerStarsTile1    *pixel.Sprite
	closerStarsTile2    *pixel.Sprite
	starsTile1Pos       pixel.Vec
	starsTile2Pos       pixel.Vec
	closerStarsTile1Pos pixel.Vec
	closerStarsTile2Pos pixel.Vec
	starsHeight         float64

	t <-chan time.Time
}

/*
NewGameBackground creates a new background manager
*/
func NewGameBackground(window *pixelgl.Window) *GameBackground {
	var err error
	var pic1 pixel.Picture
	var pic2 pixel.Picture

	if pic1, err = loadPicture("/assets/stars.png"); err != nil {
		panic(err)
	}

	if pic2, err = loadPicture("/assets/closer-stars.png"); err != nil {
		panic(err)
	}

	centerVector := window.Bounds().Center()

	topVector := window.Bounds().Center()
	topVector = topVector.Add(pixel.V(0.0, 768.0))

	return &GameBackground{
		window:              window,
		starsTile1:          pixel.NewSprite(pic1, pic1.Bounds()),
		starsTile2:          pixel.NewSprite(pic1, pic1.Bounds()),
		closerStarsTile1:    pixel.NewSprite(pic2, pic2.Bounds()),
		closerStarsTile2:    pixel.NewSprite(pic2, pic2.Bounds()),
		starsTile1Pos:       topVector,
		starsTile2Pos:       centerVector,
		closerStarsTile1Pos: topVector,
		closerStarsTile2Pos: centerVector,
		starsHeight:         pic1.Bounds().Size().Y,

		t: time.Tick(time.Millisecond * 30),
	}
}

/*
Draw renders the stars to the window
*/
func (b *GameBackground) Draw() {
	b.starsTile1.Draw(b.window, pixel.IM.Moved(b.starsTile1Pos))
	b.starsTile2.Draw(b.window, pixel.IM.Moved(b.starsTile2Pos))
	b.closerStarsTile1.Draw(b.window, pixel.IM.Moved(b.closerStarsTile1Pos))
	b.closerStarsTile2.Draw(b.window, pixel.IM.Moved(b.closerStarsTile2Pos))
}

/*
Update scrolls the background
*/
func (b *GameBackground) Update(dt float64) {
	select {
	case <-b.t:
		if b.starsTile2Pos.Y <= -384 {
			b.starsTile2Pos = b.starsTile1Pos.Add(pixel.V(0.0, b.starsHeight))
		} else if b.starsTile1Pos.Y <= -384 {
			b.starsTile1Pos = b.starsTile2Pos.Add(pixel.V(0.0, b.starsHeight))
		}

		if b.closerStarsTile2Pos.Y <= -384 {
			b.closerStarsTile2Pos = b.closerStarsTile1Pos.Add(pixel.V(0.0, b.starsHeight))
		} else if b.closerStarsTile1Pos.Y <= -384 {
			b.closerStarsTile1Pos = b.closerStarsTile2Pos.Add(pixel.V(0.0, b.starsHeight))
		}

		starsMovement := 1.0
		b.starsTile1Pos = b.starsTile1Pos.Sub(pixel.V(0.0, starsMovement))
		b.starsTile2Pos = b.starsTile2Pos.Sub(pixel.V(0.0, starsMovement))

		closerStarsMovement := 2.0
		b.closerStarsTile1Pos = b.closerStarsTile1Pos.Sub(pixel.V(0.0, closerStarsMovement))
		b.closerStarsTile2Pos = b.closerStarsTile2Pos.Sub(pixel.V(0.0, closerStarsMovement))

	default:
	}
}
