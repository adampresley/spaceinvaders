package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	// GAME_MODE_MENU is sitting at the menu
	GAME_MODE_MENU int = 1

	// GAME_MODE_PLAYING means we are playing the game
	GAME_MODE_PLAYING int = 2
)

/*
Game controls all high level aspects of the game
*/
type Game struct {
	window        *pixelgl.Window
	gameMode      int
	background    *GameBackground
	invaders      *Invaders
	player        *Player
	bulletManager *BulletManager
	atlas         *text.Atlas
	tempText      *text.Text
	fpsText       *text.Text
	menuManager   *MenuManager
}

/*
NewGame intializes the game
*/
func NewGame(window *pixelgl.Window) *Game {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	invaders := NewInvaders(window)
	player := NewPlayer(window)
	bulletManager := NewBulletManager(window, player.GetPosition(), player.GetHeight())

	return &Game{
		window:        window,
		gameMode:      GAME_MODE_MENU,
		background:    NewGameBackground(window),
		invaders:      invaders,
		player:        player,
		bulletManager: bulletManager,
		atlas:         atlas,
		tempText:      text.New(pixel.V(0.0, 0.0), atlas),
		fpsText:       text.New(pixel.V(0.0, 0.0), atlas),
		menuManager:   NewMenuManager(window),
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
	if g.gameMode == GAME_MODE_MENU {
		g.drawMenu()
	} else if g.gameMode == GAME_MODE_PLAYING {
		g.drawGame()
	}
}

func (g *Game) drawFPS() {
	g.fpsText.Clear()
	fmt.Fprintf(g.fpsText, "FPS: %d", fps)
	g.fpsText.Draw(g.window, pixel.IM)
}

func (g *Game) drawGame() {
	g.background.Draw()
	g.invaders.Draw()
	g.player.Draw()
	g.bulletManager.Draw()

	g.drawFPS()
}

func (g *Game) drawMenu() {
	g.menuManager.Draw()
}

func (g *Game) drawPlayerPosition() {
	g.tempText.Clear()
	playerPos := g.player.GetPosition()
	fmt.Fprintf(g.tempText, "X = %0.1f Y = %0.1f", playerPos.X, playerPos.Y)

	g.tempText.Draw(window, pixel.IM)
}

/*
GetGameMode returns the current game mode
*/
func (g *Game) GetGameMode() int {
	return g.gameMode
}

/*
KillHitInvaders marks any invaders who have been hit by a bullet as "dead"
*/
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

/*
Update does all the updating, such as bullets, ships, and invaders
*/
func (g *Game) Update(dt float64) {
	if g.gameMode == GAME_MODE_MENU {
		g.updateMenu(dt)
	} else {
		g.updatePlaying(dt)
	}
}

func (g *Game) updateMenu(dt float64) {
	var pressedEnter bool
	var selection int

	g.menuManager.CheckForMenuMovement()

	if pressedEnter, selection = g.menuManager.PressedEnter(); pressedEnter {
		if selection == MENU_ITEM_QUIT {
			g.window.SetClosed(true)
			return
		}

		if selection == MENU_ITEM_NEW_GAME {
			g.gameMode = GAME_MODE_PLAYING
		}
	}
}

func (g *Game) updatePlaying(dt float64) {
	g.CheckForQuit()
	g.background.Update(dt)
	g.MoveBullets(dt)
	g.MoveInvaders(dt)
	g.CheckForPlayerMovement(dt)
	g.CheckForPlayerShooting(dt)
}
