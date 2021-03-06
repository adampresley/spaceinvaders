package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	MAX_BULLETS int           = 10
	SHOOT_DELAY time.Duration = time.Millisecond * 650
)

/*
BulletManager manages bullets
*/
type BulletManager struct {
	assetManager  *AssetManager
	window        *pixelgl.Window
	bullets       []*Bullet
	currentBullet int
	nextShotTime  time.Time
}

/*
BulletHitVector tells us which invader was hit
*/
type BulletHitVector struct {
	Row int
	Col int
}

/*
NewBulletManager makes a new manager to handle tracking bullets, collisions, etc
*/
func NewBulletManager(window *pixelgl.Window, assetManager *AssetManager, playerPos pixel.Vec, playerHeight float64) *BulletManager {
	bullets := make([]*Bullet, MAX_BULLETS)

	for i := 0; i < MAX_BULLETS; i++ {
		bullets[i] = NewBullet(window, assetManager, playerPos, playerHeight)
	}

	return &BulletManager{
		assetManager:  assetManager,
		window:        window,
		bullets:       bullets,
		currentBullet: -1,
		nextShotTime:  time.Now(),
	}
}

/*
Draw renders all bullets onto the window
*/
func (bm *BulletManager) Draw() {
	for i := 0; i < MAX_BULLETS; i++ {
		bm.bullets[i].Draw()
	}
}

/*
Move moves all active bullets
*/
func (bm *BulletManager) Move(dt float64, playerPos pixel.Vec, playerHeight float64, invaders *Invaders) []BulletHitVector {
	hits := make([]BulletHitVector, 0, MAX_BULLETS)

	for i := 0; i < MAX_BULLETS; i++ {
		bm.bullets[i].Move(dt)

		isHit, row, col := bm.bullets[i].HitInvader(invaders)

		if isHit {
			bm.bullets[i].Kill()
			bm.bullets[i].Reset(playerPos, playerHeight)
			hits = append(hits, BulletHitVector{Row: row, Col: col})
		}

		if bm.bullets[i].IsTopEdge() {
			bm.bullets[i].Kill()
			bm.bullets[i].Reset(playerPos, playerHeight)
		}
	}

	return hits
}

/*
Reset puts all the bullets back
*/
func (bm *BulletManager) Reset(playerPos pixel.Vec, playerHeight float64) {
	bm.currentBullet = -1
	bm.nextShotTime = time.Now()

	for i := 0; i < MAX_BULLETS; i++ {
		bm.bullets[i].Reset(playerPos, playerHeight)
	}
}

/*
Shoot fires a bullet
*/
func (bm *BulletManager) Shoot(playerPos pixel.Vec, playHeight float64) {
	if time.Now().Before(bm.nextShotTime) {
		return
	}

	if bm.currentBullet >= MAX_BULLETS-1 {
		bm.currentBullet = -1
	}

	if bm.bullets[bm.currentBullet+1].IsAlive() {
		return
	}

	bm.currentBullet++

	bm.bullets[bm.currentBullet].Reset(playerPos, playHeight)
	bm.bullets[bm.currentBullet].Resurrect()

	bm.nextShotTime = time.Now().Add(SHOOT_DELAY)
}
