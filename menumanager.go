package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	MENU_ITEM_NEW_GAME int = 1
	MENU_ITEM_QUIT     int = 2
)

/*
MenuManager handles rendering and accepting input for the menu
*/
type MenuManager struct {
	assetManager     *AssetManager
	window           *pixelgl.Window
	menuNewGameAsset *pixel.Sprite
	menuQuitAsset    *pixel.Sprite
	currentMenuItem  int
}

/*
NewMenuManager initializes the menu manager
*/
func NewMenuManager(window *pixelgl.Window, assetManager *AssetManager) *MenuManager {
	return &MenuManager{
		assetManager:     assetManager,
		window:           window,
		menuNewGameAsset: assetManager.LoadMenuNewGameAsset(),
		menuQuitAsset:    assetManager.LoadMenuQuitAsset(),
		currentMenuItem:  MENU_ITEM_NEW_GAME,
	}
}

/*
CheckForMenuMovement will change selected menu items
*/
func (m *MenuManager) CheckForMenuMovement() {
	if m.window.JustPressed(pixelgl.KeyUp) && m.currentMenuItem == MENU_ITEM_QUIT {
		m.currentMenuItem = MENU_ITEM_NEW_GAME
	}

	if m.window.JustPressed(pixelgl.KeyDown) && m.currentMenuItem == MENU_ITEM_NEW_GAME {
		m.currentMenuItem = MENU_ITEM_QUIT
	}
}

/*
Draw renders the currently selected menu to the window
*/
func (m *MenuManager) Draw() {
	center := pixel.IM.Moved(window.Bounds().Center())

	if m.currentMenuItem == MENU_ITEM_NEW_GAME {
		m.menuNewGameAsset.Draw(m.window, center)
	} else if m.currentMenuItem == MENU_ITEM_QUIT {
		m.menuQuitAsset.Draw(m.window, center)
	}
}

/*
PressedEnter returns true if the user pressed Enter
*/
func (m *MenuManager) PressedEnter() (bool, int) {
	if m.window.JustPressed(pixelgl.KeyEnter) {
		return true, m.currentMenuItem
	}

	return false, 0
}
