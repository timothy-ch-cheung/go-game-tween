package game

import (
	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type GameMap struct {
	sprite *ebiten.Image
}

func (gameMap *GameMap) Width() int {
	return gameMap.sprite.Bounds().Dx()
}

func (gameMap *GameMap) Height() int {
	return gameMap.sprite.Bounds().Dy()
}

func (gameMap *GameMap) Draw(screen *ebiten.Image, cam *ebitenCamera.Camera) {
	op := &ebiten.DrawImageOptions{}
	cam.Surface.DrawImage(gameMap.sprite, cam.GetTranslation(op, 0, 0))
}

func (gameMap *GameMap) GetInitialPos() (float64, float64) {
	intialX := float64(gameMap.sprite.Bounds().Dx()) - config.ScreenWidth/2
	intialY := float64(gameMap.sprite.Bounds().Dy()) - config.ScreenHeight/2
	return intialX, intialY
}

func NewGameMap(loader resource.Loader) *GameMap {
	sprite := loader.LoadImage(assets.ImgMap).Data

	return &GameMap{
		sprite: sprite,
	}
}
