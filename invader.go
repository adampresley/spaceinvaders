package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Invader struct {
	window *pixelgl.Window
	sprite *pixel.Sprite
	color  int
	pos    pixel.Vec
	width  float64
	height float64
	dead   bool
}

func NewInvader(window *pixelgl.Window, color int) (*Invader, error) {
	var err error
	var img pixel.Picture

	imageToLoad := ""

	switch color {
	case 1:
		imageToLoad = "./assets/invader-blue.png"

	case 2:
		imageToLoad = "./assets/invader-green.png"

	default:
		imageToLoad = "./assets/invader-red.png"
	}

	if img, err = loadPicture(imageToLoad); err != nil {
		return nil, err
	}

	sprite := pixel.NewSprite(img, img.Bounds())

	result := &Invader{
		window: window,
		sprite: sprite,
		color:  color,
		pos:    pixel.V(0, 0),
		width:  img.Bounds().W(),
		height: img.Bounds().H(),
		dead:   false,
	}

	return result, nil
}

func (invader *Invader) IsLeftEdge() bool {
	return (invader.pos.X - (invader.width / 2)) <= 1
}

func (invader *Invader) IsRightEdge() bool {
	return (invader.pos.X + (invader.width / 2)) >= 990
}

func (invader *Invader) Draw() {
	if !invader.dead {
		invader.sprite.Draw(invader.window, pixel.IM.Moved(invader.pos))
	}
}

func (invader *Invader) Move(direction float64, dt float64) {
	move := 40.0 * direction
	x := invader.pos.X + (move * dt)

	invader.pos.X = x
}

func (invader *Invader) PushDown(lastTick time.Time) {

}

func (invader *Invader) SetPosition(pos pixel.Vec) {
	invader.pos = pos
}
