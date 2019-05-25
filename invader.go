package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
Invader represents a single alien invader
*/
type Invader struct {
	window    *pixelgl.Window
	sprite    *pixel.Sprite
	color     int
	pos       pixel.Vec
	width     float64
	height    float64
	leftEdge  float64
	rightEdge float64
	dead      bool
}

/*
NewInvader makes a new invader struct. It is initialized with one of
three colors: blue(1), green(2), or red(3).
*/
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

		leftEdge:  img.Bounds().W() / 2,
		rightEdge: window.Bounds().W() - (img.Bounds().W() / 2),
	}

	return result, nil
}

/*
Draw renders this invader onto the window
*/
func (invader *Invader) Draw() {
	if !invader.dead {
		invader.sprite.Draw(invader.window, pixel.IM.Moved(invader.pos))
	}
}

/*
IsLeftEdge returns true if the invader is on the left edge of the window
*/
func (invader *Invader) IsLeftEdge() bool {
	return invader.pos.X <= invader.leftEdge
}

/*
IsRightEdge returns true if the invader is on the right edge of the window
*/
func (invader *Invader) IsRightEdge() bool {
	return invader.pos.X >= invader.rightEdge
}

/*
Move advances the invader to the left or right. Direction
is either 1 for right, or -1 for left.
*/
func (invader *Invader) Move(direction float64, dt float64) {
	move := 140.0 * direction
	x := invader.pos.X + (move * dt)
	invader.pos.X = pixel.Clamp(x, invader.leftEdge, invader.rightEdge)
}

/*
PushDown moves the invader down a row
*/
func (invader *Invader) PushDown() {
	invader.pos.Y -= 20.0
}

/*
SetPosition sets the invader's vector
*/
func (invader *Invader) SetPosition(pos pixel.Vec) {
	invader.pos = pos
}
