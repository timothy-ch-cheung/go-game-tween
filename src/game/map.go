package game

import (
	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"

	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
)

type GameMap struct {
	sprite *ebiten.Image
	posX   int
	posY   int
}

func (gameMap *GameMap) Width() int {
	return gameMap.sprite.Bounds().Dx()
}

func (gameMap *GameMap) Height() int {
	return gameMap.sprite.Bounds().Dy()
}

func (gameMap *GameMap) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(gameMap.posX), float64(gameMap.posY))
	screen.DrawImage(gameMap.sprite, op)
}

func NewGameMap(loader resource.Loader) *GameMap {
	sprite := loader.LoadImage(assets.ImgMap).Data
	intialX := -float64(sprite.Bounds().Dx()) + config.ScreenWidth
	intialY := -float64(sprite.Bounds().Dy()) + config.ScreenHeight
	return &GameMap{
		sprite: sprite,
		posX:   int(intialX),
		posY:   int(intialY),
	}
}
