package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	window        *pixelgl.Window
	background    *pixel.Sprite
	invaders      *Invaders
	player        *Player
	bulletManager *BulletManager
	atlas         *text.Atlas
	tempText      *text.Text
	fpsText       *text.Text
}

func NewGame(window *pixelgl.Window) *Game {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	backgroundPicture, _ := loadPicture("/assets/stars.png")

	invaders := NewInvaders(window)
	player := NewPlayer(window)
	bulletManager := NewBulletManager(window, player.GetPosition(), player.GetHeight())

	return &Game{
		window:        window,
		background:    pixel.NewSprite(backgroundPicture, backgroundPicture.Bounds()),
		invaders:      invaders,
		player:        player,
		bulletManager: bulletManager,
		atlas:         atlas,
		tempText:      text.New(pixel.V(0.0, 0.0), atlas),
		fpsText:       text.New(pixel.V(0.0, 0.0), atlas),
	}
}

/*
CheckForQuit looks to see if the player has chosen to quit
*/
func (g *Game) CheckForQuit() {
	if g.window.JustPressed(pixelgl.KeyQ) {
		g.window.SetClosed(true)
	}
}

/*
CheckForPlayerMovement will move the player when left and right are pressed
*/
func (g *Game) CheckForPlayerMovement(dt float64) {
	if g.window.Pressed(pixelgl.KeyLeft) {
		g.player.MoveLeft(dt)
	}

	if g.window.Pressed(pixelgl.KeyRight) {
		g.player.MoveRight(dt)
	}
}

/*
CheckForPlayerShooting will shoot bullets when the spacebar is pressed
*/
func (g *Game) CheckForPlayerShooting(dt float64) {
	if g.player.IsShooting() {
		g.bulletManager.Shoot(g.player.GetPosition(), g.player.GetHeight())
	}
}

/*
Draw renders the background, invaders, player, and bullets to the window
*/
func (g *Game) Draw() {
	g.background.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
	g.invaders.Draw()
	g.player.Draw()
	g.bulletManager.Draw()

	//g.drawPlayerPosition()
	g.drawFPS()
}

func (g *Game) drawFPS() {
	g.fpsText.Clear()
	fmt.Fprintf(g.fpsText, "FPS: %d", fps)
	g.fpsText.Draw(g.window, pixel.IM)
}

func (g *Game) drawPlayerPosition() {
	g.tempText.Clear()
	playerPos := g.player.GetPosition()
	fmt.Fprintf(g.tempText, "X = %0.1f Y = %0.1f", playerPos.X, playerPos.Y)

	g.tempText.Draw(window, pixel.IM)
}

func (g *Game) KillHitInvaders(hits []BulletHitVector) {
	for _, h := range hits {
		g.invaders.Kill(h.Row, h.Col)
	}
}

/*
MoveBullets moves all active bullets
*/
func (g *Game) MoveBullets(dt float64) {
	hits := g.bulletManager.Move(dt, g.player.GetPosition(), g.player.GetHeight(), g.invaders)
	g.KillHitInvaders(hits)
}

/*
MoveInvaders moves all active invaders
*/
func (g *Game) MoveInvaders(dt float64) {
	g.invaders.Move(dt)
}
