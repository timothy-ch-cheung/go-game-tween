package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
)

type MiniMap struct {
	Cam *ebitenCamera.Camera
}

func (miniMap *MiniMap) Draw(screen *ebiten.Image, loader *resource.Loader) {
	sprite := loader.LoadImage(assets.ImgMiniMap).Data
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(config.ScreenWidth-sprite.Bounds().Dx()), 0)
	screen.DrawImage(sprite, op)
}
