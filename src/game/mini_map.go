package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
)

type MiniMap struct {
	mapWidth  float64
	mapHeight float64
	cam       *ebitenCamera.Camera
	posX      float64
	posY      float64
	xScale    float64
	yScale    float64
}

func getMiniMapViewCoords(miniMap *MiniMap, viewSprite *ebiten.Image) (float64, float64) {
	xOffset := miniMap.cam.X / miniMap.xScale
	yOffset := miniMap.cam.Y / miniMap.yScale
	x := miniMap.posX + xOffset - float64(viewSprite.Bounds().Dx()/2)
	y := miniMap.posY + yOffset - float64(viewSprite.Bounds().Dy()/2)
	return x, y
}

func (miniMap *MiniMap) Draw(screen *ebiten.Image, loader *resource.Loader) {
	sprite := loader.LoadImage(assets.ImgMiniMap).Data
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(miniMap.posX, miniMap.posY)
	screen.DrawImage(sprite, op)

	viewSprite := loader.LoadImage(assets.ImgMiniMapView).Data
	viewOp := &ebiten.DrawImageOptions{}
	viewOp.GeoM.Translate(getMiniMapViewCoords(miniMap, viewSprite))
	screen.DrawImage(viewSprite, viewOp)
}

func NewMiniMap(mapWidth float64, mapHeight float64, cam *ebitenCamera.Camera, loader *resource.Loader) *MiniMap {
	sprite := loader.LoadImage(assets.ImgMiniMap).Data.Bounds()
	posX := float64(config.ScreenWidth - sprite.Bounds().Dx())
	xScale := mapWidth / float64(sprite.Dx())
	yScale := mapHeight / float64(sprite.Dy())
	return &MiniMap{
		mapWidth:  mapWidth,
		mapHeight: mapHeight,
		cam:       cam,
		posX:      posX,
		posY:      0,
		xScale:    xScale,
		yScale:    yScale,
	}
}
