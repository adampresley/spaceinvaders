package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/faiface/pixel"
)

type AssetManager struct {
	Batch *pixel.Batch
	spritesheet pixel.Picture
}

func NewAssetManager() *AssetManager {
	result := &AssetManager{}

	result.spritesheet = result.loadSpritesheet()
	result.Batch = result.createBatch(result.spritesheet)
	return result
}

func (am *AssetManager) createBatch(spritesheet pixel.Picture) *pixel.Batch {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	return batch
}
/*
loadPicture loads an image file into a picture struct
*/
func (am *AssetManager) loadPicture(path string) (pixel.Picture, error) {
	var err error
	var img image.Image
	var imageBytes []byte

	if imageBytes, err = FSByte(false, path); err != nil {
		return nil, err
	}

	reader := bytes.NewReader(imageBytes)

	if img, _, err = image.Decode(reader); err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func (am *AssetManager) loadSpritesheet() pixel.Picture {
	var err error
	var pic pixel.Picture

	if pic, err = am.loadPicture("/assets/spritesheet.png"); err != nil {
		panic(err)
	}

	return pic
}

func (am *AssetManager) GetBulletSprite() *pixel.Sprite {
	result := pixel.NewSprite(am.spritesheet, pixel.R(0, 0, 10, 38))
	return result
}

func (am *AssetManager) GetShipSprite() *pixel.Sprite {
	result := pixel.NewSprite(am.spritesheet, pixel.R(10, 0, 60, 38))
	return result
}

func (am *AssetManager) GetBlueInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(am.spritesheet, pixel.R(61, 0, 110, 38))
	return result
}

func (am *AssetManager) GetGreenInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(am.spritesheet, pixel.R(111, 0, 160, 38))
	return result
}

func (am *AssetManager) GetRedInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(am.spritesheet, pixel.R(161, 0, 210, 38))
	return result
}

func (am *AssetManager) LoadMenuNewGameAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = am.loadPicture("/assets/menu-new-game.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}

func (am *AssetManager) LoadMenuQuitAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = am.loadPicture("/assets/menu-quit.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}

func (am *AssetManager) LoadYouLoseAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = am.loadPicture("/assets/you-lose.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}

func (am *AssetManager) LoadYouWinAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = am.loadPicture("/assets/you-win.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}

func (am *AssetManager) GetStarsAsset() pixel.Picture {
	pic, err := am.loadPicture("/assets/stars.png")

	if err != nil {
		panic(err)
	}

	return pic
}

func (am *AssetManager) GetCloserStarsAsset() pixel.Picture {
	pic, err := am.loadPicture("/assets/closer-stars.png")

	if err != nil {
		panic(err)
	}

	return pic
}
