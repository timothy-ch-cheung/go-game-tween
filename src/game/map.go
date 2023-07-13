package game

import (
	"container/list"

	"github.com/timothy-ch-cheung/go-game-tween/assets"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type GameMap struct {
	width  float64
	height float64

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
		markers:       markers,
		CurrentMarker: markers.Front(),
	}
}

func (gameMap *GameMap) GetDimensions() (float64, float64) {
	return gameMap.width, gameMap.height
}

func (gameMap *GameMap) GetCurrentMarker() *Marker {
	return gameMap.CurrentMarker.Value.(*Marker)
}
