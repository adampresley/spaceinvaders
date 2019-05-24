package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Invaders struct {
	window    *pixelgl.Window
	invaders  [4][10]*Invader
	direction float64
}

func NewInvaders(window *pixelgl.Window) *Invaders {
	var err error
	var newInvader *Invader
	var invaders [4][10]*Invader

	y := 500.0

	for row := 0; row < len(invaders); row++ {
		x := 180.0

		for col := 0; col < len(invaders[row]); col++ {
			color := rand.Intn(3)

			if newInvader, err = NewInvader(window, color); err != nil {
				panic(err)
			}

			newInvader.SetPosition(pixel.V(x, y))
			invaders[row][col] = newInvader

			x += 115.0
		}

		y += 120.0
	}

	return &Invaders{
		window:    window,
		invaders:  invaders,
		direction: 1,
	}
}

func (invaders *Invaders) Draw() {
	for row := 0; row < len(invaders.invaders); row++ {
		for col := 0; col < len(invaders.invaders[row]); col++ {
			invaders.invaders[row][col].Draw()
		}
	}
}

func (invaders *Invaders) Move(dt float64) {
	if invaders.direction == 1 {
		if invaders.invaders[0][9].IsRightEdge() {
			invaders.direction = -1
		}
	} else {
		if invaders.invaders[0][0].IsLeftEdge() {
			invaders.direction = 1
		}
	}

	for row := 0; row < len(invaders.invaders); row++ {
		for col := 0; col < len(invaders.invaders[row]); col++ {
			invaders.invaders[row][col].Move(invaders.direction, dt)
		}
	}
}
