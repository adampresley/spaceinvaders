package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	window *pixelgl.Window
	sprite *pixel.Sprite
	pos    pixel.Vec
	width  float64
	height float64
	dead   bool
}

func NewPlayer(window *pixelgl.Window) *Player {
	var err error
	var img pixel.Picture

	if img, err = loadPicture("./assets/ship.png"); err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())
	pos := window.Bounds().Center()
	pos.Y = 0

	return &Player{
		window: window,
		sprite: sprite,
		pos:    pos,
		width:  img.Bounds().W(),
		height: img.Bounds().H(),
		dead:   false,
	}
}

func (player *Player) Draw() {
	if !player.dead {
		player.sprite.Draw(player.window, pixel.IM.Moved(player.pos))
	}
}

func (player *Player) MoveLeft(dt float64) {
	move := -260.0
	x := player.pos.X + (move * dt)

	player.pos.X = x
}

func (player *Player) MoveRight(dt float64) {
	move := 260.0
	x := player.pos.X + (move * dt)

	player.pos.X = x
}
