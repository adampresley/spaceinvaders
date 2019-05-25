package main

import (
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
HitInvader return true, and the row/col of the invader hit when there is
a collision
*/
func (b *Bullet) HitInvader(invaders *Invaders) (bool, int, int) {
	ivs := invaders.GetInvaders()

	for row := 0; row < MAX_ROWS; row++ {
		for col := 0; col < MAX_COLS; col++ {
			rect := ivs[row][col].GetRect()
			bulletLeft := b.pos.X - (b.width / 2)
			bulletRight := b.pos.X + (b.width / 2)
			bulletTop := b.pos.Y + (b.height / 2)
			bulletBottom := b.pos.Y - (b.height / 2)

			if ivs[row][col].IsAlive() && bulletRight >= rect.Min.X && bulletLeft <= rect.Max.X && bulletTop >= rect.Min.Y && bulletBottom <= rect.Max.Y {
				return true, row, col
			}
		}
	}

	return false, 0, 0
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
}

/*
Resurrect sets the bullet as not dead
*/
func (b *Bullet) Resurrect() {
	b.dead = false
}
