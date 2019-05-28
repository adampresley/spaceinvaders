package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	MAX_ROWS int = 4
	MAX_COLS int = 10
)

/*
Invaders manages a set of invader structs and their movement
*/
type Invaders struct {
	assetManager      *AssetManager
	direction         float64
	firstCol          int
	invaders          [MAX_ROWS][MAX_COLS]*Invader
	lastCol           int
	lowestInvaderRect pixel.Rect
	numInvadersLeft   int
	window            *pixelgl.Window
}

/*
NewInvaders creates a new struct. It is initialied with a set of
invaders, each positioned in rows and columns
*/
func NewInvaders(window *pixelgl.Window, assetManager *AssetManager) *Invaders {
	var err error
	var newInvader *Invader
	var invaders [MAX_ROWS][MAX_COLS]*Invader

	y := window.Bounds().H() - 25

	for row := 0; row < MAX_ROWS; row++ {
		x := 180.0

		for col := 0; col < MAX_COLS; col++ {
			color := rand.Intn(3)

			if newInvader, err = NewInvader(window, color, assetManager); err != nil {
				panic(err)
			}

			newInvader.SetPosition(pixel.V(x, y))
			invaders[row][col] = newInvader

			x += 70.0
		}

		y -= 50.0
	}

	return &Invaders{
		assetManager:      assetManager,
		direction:         1,
		firstCol:          0,
		invaders:          invaders,
		lastCol:           MAX_COLS - 1,
		lowestInvaderRect: invaders[MAX_ROWS-1][0].GetRect(),
		numInvadersLeft:   MAX_ROWS * MAX_COLS,
		window:            window,
	}
}

/*
Draw renders all invaders onto the window
*/
func (invaders *Invaders) Draw() {
	for row := 0; row < MAX_ROWS; row++ {
		for col := 0; col < MAX_COLS; col++ {
			invaders.invaders[row][col].Draw()
		}
	}
}

/*
GetInvaders returns the invaders slice
*/
func (invaders *Invaders) GetInvaders() [MAX_ROWS][MAX_COLS]*Invader {
	return invaders.invaders
}

/*
GetNumInvadersLeft returns the number of invaders left alive
*/
func (invaders *Invaders) GetNumInvadersLeft() int {
	return invaders.numInvadersLeft
}

func (invaders *Invaders) HaveReachedBottom(player *Player) bool {
	playerRect := player.GetRect()

	if invaders.lowestInvaderRect.Min.Y <= playerRect.Min.Y {
		return true
	}

	return false
}

/*
Kill marks an invader as dead
*/
func (invaders *Invaders) Kill(row, col int) {
	invaders.invaders[row][col].dead = true
	invaders.numInvadersLeft--
	invaders.recalculateFirstAndLastColumn()
}

/*
Move advances all invaders. When the group hits a window edge
they are moved down, and the direction reversed
*/
func (invaders *Invaders) Move(dt float64) {
	if invaders.direction == 1 {
		if invaders.invaders[0][invaders.lastCol].IsRightEdge() {
			invaders.direction = -1
			invaders.PushDown()
		}
	} else {
		if invaders.invaders[0][invaders.firstCol].IsLeftEdge() {
			invaders.direction = 1
			invaders.PushDown()
		}
	}

	for row := 0; row < MAX_ROWS; row++ {
		for col := 0; col < MAX_COLS; col++ {
			invaders.invaders[row][col].Move(invaders.direction, dt)
		}
	}
}

/*
PushDown moves all invaders down a row
*/
func (invaders *Invaders) PushDown() {
	lowestRow := 0
	lowestCol := 0

	for row := 0; row < MAX_ROWS; row++ {
		for col := 0; col < MAX_COLS; col++ {
			invaders.invaders[row][col].PushDown()

			if invaders.invaders[row][col].IsAlive() {
				if row > lowestRow {
					lowestRow = row
				}

				if col > lowestCol {
					lowestCol = col
				}
			}
		}
	}

	invader := invaders.invaders[lowestRow][lowestCol]
	invaders.lowestInvaderRect = invader.GetRect()
}

func (invaders *Invaders) recalculateFirstAndLastColumn() {
	lastCol := 0
	firstCol := MAX_COLS - 1

	var col int

	for row := 0; row < MAX_ROWS; row++ {
		for col = 0; col < MAX_COLS; col++ {
			if invaders.invaders[row][col].IsAlive() {
				if col > lastCol {
					lastCol = col
				}
			}
		}

		for col = MAX_COLS - 1; col > -1; col-- {
			if invaders.invaders[row][col].IsAlive() {
				if col < firstCol {
					firstCol = col
				}
			}
		}
	}

	invaders.lastCol = lastCol
	invaders.firstCol = firstCol
}

/*
Reset puts all the invaders back into their place and brings them
back to life
*/
func (invaders *Invaders) Reset() {
	y := window.Bounds().H() - 25

	for row := 0; row < MAX_ROWS; row++ {
		x := 180.0

		for col := 0; col < MAX_COLS; col++ {
			invaders.invaders[row][col].SetPosition(pixel.V(x, y))
			invaders.invaders[row][col].Resurrect()

			x += 70.0
		}

		y -= 50.0
	}

	invaders.numInvadersLeft = MAX_ROWS * MAX_COLS
	invaders.lastCol = MAX_COLS - 1
}
