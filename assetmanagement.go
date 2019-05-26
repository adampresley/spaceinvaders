package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/faiface/pixel"
)

/*
loadPicture loads an image file into a picture struct
*/
func loadPicture(path string) (pixel.Picture, error) {
	var err error
	//var file *os.File
	var img image.Image
	var imageBytes []byte

	if imageBytes, err = FSByte(false, path); err != nil {
		return nil, err
	}

	// if file, err = os.Open(path); err != nil {
	// 	return nil, err
	// }

	// defer file.Close()

	reader := bytes.NewReader(imageBytes)

	//if img, _, err = image.Decode(file); err != nil {
	if img, _, err = image.Decode(reader); err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func loadSpritesheet() pixel.Picture {
	var err error
	var pic pixel.Picture

	if pic, err = loadPicture("/assets/spritesheet.png"); err != nil {
		panic(err)
	}

	return pic
}

func getBulletSprite() *pixel.Sprite {
	result := pixel.NewSprite(spritesheet, pixel.R(0, 0, 10, 38))
	return result
}

func getShipSprite() *pixel.Sprite {
	result := pixel.NewSprite(spritesheet, pixel.R(10, 0, 60, 38))
	return result
}

func getBlueInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(spritesheet, pixel.R(61, 0, 110, 38))
	return result
}

func getGreenInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(spritesheet, pixel.R(111, 0, 160, 38))
	return result
}

func getRedInvaderSprite() *pixel.Sprite {
	result := pixel.NewSprite(spritesheet, pixel.R(161, 0, 210, 38))
	return result
}

func loadMenuNewGameAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = loadPicture("/assets/menu-new-game.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}

func loadMenuQuitAsset() *pixel.Sprite {
	var err error
	var pic pixel.Picture

	if pic, err = loadPicture("/assets/menu-quit.png"); err != nil {
		panic(err)
	}

	result := pixel.NewSprite(pic, pic.Bounds())
	return result
}
