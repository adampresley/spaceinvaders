package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	window   *pixelgl.Window
	invaders *Invaders
	player   *Player
	atlas    *text.Atlas
	posText  *text.Text
}

func NewGame(window *pixelgl.Window) *Game {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	return &Game{
		window:   window,
		invaders: NewInvaders(window),
		player:   NewPlayer(window),
		atlas:    atlas,
		posText:  text.New(pixel.V(0, 0), atlas),
	}
}

func (g *Game) CheckForQuit() {
	if g.window.JustPressed(pixelgl.KeyQ) {
		g.window.SetClosed(true)
	}
}

func (g *Game) CheckForPlayerMovement(dt float64) {
	if g.window.Pressed(pixelgl.KeyLeft) {
		g.player.MoveLeft(dt)
	}

	if g.window.Pressed(pixelgl.KeyRight) {
		g.player.MoveRight(dt)
	}
}

func (g *Game) Draw() {
	g.invaders.Draw()
	g.player.Draw()
}

func (g *Game) MoveInvaders(dt float64) {
	g.invaders.Move(dt)
}