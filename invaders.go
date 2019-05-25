package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

/*
Invaders manages a set of invader structs and their movement
*/
type Invaders struct {
	window    *pixelgl.Window
	invaders  [4][10]*Invader
	direction float64
}

/*
NewInvaders creates a new struct. It is initialied with a set of
invaders, each positioned in rows and columns
*/
func NewInvaders(window *pixelgl.Window) *Invaders {
	var err error
	var newInvader *Invader
	var invaders [4][10]*Invader

	y := window.Bounds().H() - 25

	for row := 0; row < len(invaders); row++ {
		x := 180.0

		for col := 0; col < len(invaders[row]); col++ {
			color := rand.Intn(3)

			if newInvader, err = NewInvader(window, color); err != nil {
				panic(err)
			}

			newInvader.SetPosition(pixel.V(x, y))
			invaders[row][col] = newInvader

			x += 70.0
		}

		y -= 50.0
	}

	return &Invaders{
		window:    window,
		invaders:  invaders,
		direction: 1,
	}
}

/*
Draw renders all invaders onto the window
*/
func (invaders *Invaders) Draw() {
	for row := 0; row < len(invaders.invaders); row++ {
		for col := 0; col < len(invaders.invaders[row]); col++ {
			invaders.invaders[row][col].Draw()
		}
	}
}

/*
Move advances all invaders. When the group hits a window edge
they are moved down, and the direction reversed
*/
func (invaders *Invaders) Move(dt float64) {
	if invaders.direction == 1 {
		if invaders.invaders[0][9].IsRightEdge() {
			invaders.direction = -1
			invaders.PushDown()
		}
	} else {
		if invaders.invaders[0][0].IsLeftEdge() {
			invaders.direction = 1
			invaders.PushDown()
		}
	}

	for row := 0; row < len(invaders.invaders); row++ {
		for col := 0; col < len(invaders.invaders[row]); col++ {
			invaders.invaders[row][col].Move(invaders.direction, dt)
		}
	}
}

/*
PushDown moves all invaders down a row
*/
func (invaders *Invaders) PushDown() {
	for row := 0; row < len(invaders.invaders); row++ {
		for col := 0; col < len(invaders.invaders[row]); col++ {
			invaders.invaders[row][col].PushDown()
		}
	}
}
