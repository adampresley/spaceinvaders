package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
Bullet is the ammo a player uses
*/
type Bullet struct {
	window  *pixelgl.Window
	sprite  *pixel.Sprite
	pos     pixel.Vec
	width   float64
	height  float64
	topEdge float64
	dead    bool
}

/*
NewBullet initializes a new bullet. It is setup to start at the player's position
*/
func NewBullet(window *pixelgl.Window, playerPos pixel.Vec, playerHeight float64) *Bullet {
	sprite := getBulletSprite()

	return &Bullet{
		window:  window,
		sprite:  sprite,
		pos:     pixel.V(playerPos.X, playerPos.Y+(playerHeight/2)),
		width:   sprite.Frame().W(),
		height:  sprite.Frame().H(),
		topEdge: window.Bounds().H() + (sprite.Frame().H() / 2),
		dead:    true,
	}
}

/*
Draw renders this bullet onto the window
*/
func (b *Bullet) Draw() {
	if !b.dead {
		b.sprite.Draw(b.window, pixel.IM.Moved(b.pos))
	}
}

/*
IsAlive returns true if the bullet is alive
*/
func (b *Bullet) IsAlive() bool {
	return !b.dead
}

/*
IsDead returns true if the bullet is dead
*/
func (b *Bullet) IsDead() bool {
	return b.dead
}

/*
IsTopEdge returns true if the bullet has gone off the top of the window
*/
func (b *Bullet) IsTopEdge() bool {
	return b.pos.Y >= b.topEdge
}

/*
Kill marks the bullet as dead
*/
func (b *Bullet) Kill() {
	b.dead = true
}

/*
Move advances the bullet toward the top
*/
func (b *Bullet) Move(dt float64) {
	if !b.dead {
		move := 200.0 * dt
		y := b.pos.Y + move
		b.pos.Y = pixel.Clamp(y, 0, b.topEdge)
	}
}

/*
Reset positions the bullet to the player
*/
func (b *Bullet) Reset(playerPos pixel.Vec, playerHeight float64) {
	b.pos = pixel.V(playerPos.X, playerPos.Y+(playerHeight/2))
	fmt.Printf("Bullet pos: %f,%f\n", b.pos.X, b.pos.Y)
}

/*
Resurrect sets the bullet as not dead
*/
func (b *Bullet) Resurrect() {
	b.dead = false
}
