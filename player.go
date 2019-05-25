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
	sprite := getShipSprite()
	pos := window.Bounds().Center()
	pos.Y = 16

	return &Player{
		window:    window,
		sprite:    sprite,
		pos:       pos,
		width:     sprite.Frame().W(),
		height:    sprite.Frame().H(),
		leftEdge:  sprite.Frame().W() / 2,
		rightEdge: window.Bounds().W() - (sprite.Frame().W() / 2),
		dead:      false,
	}
}

/*
Draw renders this ship onto the window
*/
func (p *Player) Draw() {
	if !p.dead {
		p.sprite.Draw(p.window, pixel.IM.Moved(p.pos))
	}
}

/*
GetHeight returns the ship's height
*/
func (p *Player) GetHeight() float64 {
	return p.height
}

/*
GetPosition retrieves the player's position vector
*/
func (p *Player) GetPosition() pixel.Vec {
	return p.pos
}

/*
IsLeftEdge returns true if the player is on the left edge of the window
*/
func (p *Player) IsLeftEdge() bool {
	return p.pos.X <= p.leftEdge
}

/*
IsRightEdge returns true if the player is on the right edge of the window
*/
func (p *Player) IsRightEdge() bool {
	return p.pos.X >= p.rightEdge
}

/*
IsShooting returns true if the player is pressing Space
*/
func (p *Player) IsShooting() bool {
	return p.window.Pressed(pixelgl.KeySpace)
}

/*
MoveLeft moves the player left, repspecting the left edge boundary
*/
func (p *Player) MoveLeft(dt float64) {
	move := -300.0
	x := p.pos.X + (move * dt)
	p.pos.X = pixel.Clamp(x, p.leftEdge, p.rightEdge)
}

/*
MoveRight moves the player right, respecting the right edge boundary
*/
func (p *Player) MoveRight(dt float64) {
	move := 300.0
	x := p.pos.X + (move * dt)
	p.pos.X = pixel.Clamp(x, p.leftEdge, p.rightEdge)
}
