package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
Bullet is the ammo a player uses
*/
type Bullet struct {
	window    *pixelgl.Window
	sprite    *pixel.Sprite
	pos       pixel.Vec
	width     float64
	height    float64
	leftEdge  float64
	rightEdge float64
}

//func NewBullet(window *pixelgl.Window, )
