package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
Player is the ship a player uses to fight aliens!
*/
type Player struct {
	window    *pixelgl.Window
	sprite    *pixel.Sprite
	pos       pixel.Vec
	width     float64
	height    float64
	leftEdge  float64
	rightEdge float64
	dead      bool
}

/*
NewPlayer creates a new player struct. Is is initialized positioned
at the bottom middle of the window
*/
func NewPlayer(window *pixelgl.Window) *Player {
	var err error
	var img pixel.Picture

	if img, err = loadPicture("./assets/ship.png"); err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())
	pos := window.Bounds().Center()
	pos.Y = 16

	return &Player{
		window: window,
		sprite: sprite,
		pos:    pos,
		width:  img.Bounds().W(),
		height: img.Bounds().H(),
		dead:   false,

		leftEdge:  img.Bounds().W() / 2,
		rightEdge: window.Bounds().W() - (img.Bounds().W() / 2),
	}
}

/*
Draw renders this ship onto the window
*/
func (player *Player) Draw() {
	if !player.dead {
		player.sprite.Draw(player.window, pixel.IM.Moved(player.pos))
	}
}

/*
GetPosition retrieves the player's position vector
*/
func (player *Player) GetPosition() pixel.Vec {
	return player.pos
}

/*
IsLeftEdge returns true if the player is on the left edge of the window
*/
func (player *Player) IsLeftEdge() bool {
	return player.pos.X <= player.leftEdge
}

/*
IsRightEdge returns true if the player is on the right edge of the window
*/
func (player *Player) IsRightEdge() bool {
	return player.pos.X >= player.rightEdge
}

/*
MoveLeft moves the player left, repspecting the left edge boundary
*/
func (player *Player) MoveLeft(dt float64) {
	move := -300.0
	x := player.pos.X + (move * dt)
	player.pos.X = pixel.Clamp(x, player.leftEdge, player.rightEdge)
}

/*
MoveRight moves the player right, respecting the right edge boundary
*/
func (player *Player) MoveRight(dt float64) {
	move := 300.0
	x := player.pos.X + (move * dt)
	player.pos.X = pixel.Clamp(x, player.leftEdge, player.rightEdge)
}
