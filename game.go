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

	// GAME_MODE_WON means the player won!
	GAME_MODE_WON int = 3

	// GAME_MODE_LOST means the player lost!
	GAME_MODE_LOST int = 4
)

/*
Game controls all high level aspects of the game
*/
type Game struct {
	assetManager  *AssetManager
	atlas         *text.Atlas
	background    *GameBackground
	bulletManager *BulletManager
	fpsText       *text.Text
	gameMode      int
	invaders      *Invaders
	menuManager   *MenuManager
	player        *Player
	tempText      *text.Text
	window        *pixelgl.Window
	youLoseAsset  *pixel.Sprite
	youWinAsset   *pixel.Sprite
}

/*
NewGame intializes the game
*/
func NewGame(window *pixelgl.Window) *Game {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	assetManager := NewAssetManager()
	invaders := NewInvaders(window, assetManager)
	player := NewPlayer(window, assetManager)
	bulletManager := NewBulletManager(window, assetManager, player.GetPosition(), player.GetHeight())

	return &Game{
		assetManager:  assetManager,
		atlas:         atlas,
		background:    NewGameBackground(window, assetManager),
		bulletManager: bulletManager,
		fpsText:       text.New(pixel.V(0.0, 0.0), atlas),
		gameMode:      GAME_MODE_MENU,
		invaders:      invaders,
		menuManager:   NewMenuManager(window, assetManager),
		player:        player,
		tempText:      text.New(pixel.V(0.0, 0.0), atlas),
		window:        window,
		youLoseAsset:  assetManager.LoadYouLoseAsset(),
		youWinAsset:   assetManager.LoadYouWinAsset(),
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
	} else if g.gameMode == GAME_MODE_WON {
		g.drawPlayerWon()
	} else if g.gameMode == GAME_MODE_LOST {
		g.drawPlayerLost()
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
	g.drawFPS()
}

func (g *Game) drawPlayerPosition() {
	g.tempText.Clear()
	playerPos := g.player.GetPosition()
	fmt.Fprintf(g.tempText, "X = %0.1f Y = %0.1f", playerPos.X, playerPos.Y)

	g.tempText.Draw(window, pixel.IM)
}

func (g *Game) drawPlayerLost() {
	g.background.Draw()
	g.youLoseAsset.Draw(g.window, pixel.IM.Moved(g.window.Bounds().Center()))
	g.drawFPS()
}

func (g *Game) drawPlayerWon() {
	g.background.Draw()
	g.youWinAsset.Draw(g.window, pixel.IM.Moved(g.window.Bounds().Center()))
	g.drawFPS()
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
	} else if g.gameMode == GAME_MODE_WON {
		g.updatePlayerWon(dt)
	} else if g.gameMode == GAME_MODE_LOST {
		g.updatePlayerLost(dt)
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
			g.invaders.Reset()
			g.bulletManager.Reset(g.player.GetPosition(), g.player.GetHeight())
			g.gameMode = GAME_MODE_PLAYING
		}
	}
}

func (g *Game) updatePlayerLost(dt float64) {
	if g.window.JustPressed(pixelgl.KeyEnter) {
		g.gameMode = GAME_MODE_MENU
	}

	g.CheckForQuit()
	g.background.Update(dt)
}

func (g *Game) updatePlayerWon(dt float64) {
	if g.window.JustPressed(pixelgl.KeyEnter) {
		g.gameMode = GAME_MODE_MENU
	}

	g.CheckForQuit()
	g.background.Update(dt)
}

func (g *Game) updatePlaying(dt float64) {
	g.CheckForQuit()
	g.background.Update(dt)
	g.MoveBullets(dt)
	g.MoveInvaders(dt)
	g.CheckForPlayerMovement(dt)
	g.CheckForPlayerShooting(dt)

	if g.invaders.GetNumInvadersLeft() <= 0 {
		g.gameMode = GAME_MODE_WON
	}

	if g.invaders.HaveReachedBottom(g.player) {
		g.gameMode = GAME_MODE_LOST
	}
}
