package game

import (
	"container/list"
	"math"

	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type cameraBounds struct {
	xMin float64
	xMax float64
	yMin float64
	yMax float64
}

type GameMap struct {
	width         float64
	height        float64
	bounds        cameraBounds
	markers       *list.List
	CurrentMarker *list.Element
}

func (gameMap *GameMap) Draw(screen *ebiten.Image, cam *ebitenCamera.Camera, loader *resource.Loader) {
	op := &ebiten.DrawImageOptions{}
	cam.Surface.DrawImage(loader.LoadImage(assets.ImgMap).Data, cam.GetTranslation(op, 0, 0))

	for marker := gameMap.markers.Front(); marker != nil; marker = marker.Next() {
		marker.Value.(*Marker).Draw(screen, cam, loader)
	}
}

func NewGameMap(loader resource.Loader) *GameMap {
	sprite := loader.LoadImage(assets.ImgMap).Data
	width := float64(sprite.Bounds().Dx())
	height := float64(sprite.Bounds().Dy())

	xBorder := float64(config.ScreenWidth / 2)
	yBorder := float64(config.ScreenHeight / 2)
	bounds := cameraBounds{
		xMin: xBorder,
		xMax: width - xBorder,
		yMin: yBorder,
		yMax: height - yBorder,
	}

	markers := list.New()
	markers.PushBack(newMarker(420, 420, Selected))
	markers.PushBack(newMarker(400, 415, Locked))
	markers.PushBack(newMarker(375, 425, Locked))
	markers.PushBack(newMarker(340, 405, Locked))
	markers.PushBack(newMarker(320, 380, Locked))
	markers.PushBack(newMarker(330, 350, Locked))
	markers.PushBack(newMarker(300, 300, Locked))

	return &GameMap{
		width:         width,
		height:        height,
		bounds:        bounds,
		markers:       markers,
		CurrentMarker: markers.Front(),
	}
}

// 420, max 400, min  80
func (gameMap *GameMap) GetCameraPosition() (float64, float64) {
	marker := gameMap.CurrentMarker.Value.(*Marker)
	x := math.Min(math.Max(float64(marker.posX), gameMap.bounds.xMin), gameMap.bounds.xMax)
	y := math.Min(math.Max(float64(marker.posY), gameMap.bounds.yMin), gameMap.bounds.yMax)
	return x, y
}
