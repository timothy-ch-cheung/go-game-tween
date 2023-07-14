package game

import (
	"container/list"
	"strconv"

	"github.com/timothy-ch-cheung/go-game-tween/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type GameMap struct {
	width  float64
	height float64

	Markers       *list.List
	CurrentMarker *list.Element
}

func centerText(text string, options *ebiten.DrawImageOptions) {
	switch len(text) {
	case 1:
		options.GeoM.Translate(1, -2)
	case 2:
		options.GeoM.Translate(-2, -2)
	}
}

func (gameMap *GameMap) Draw(screen *ebiten.Image, cam *ebitenCamera.Camera, loader *resource.Loader) {
	op := &ebiten.DrawImageOptions{}
	cam.Surface.DrawImage(loader.LoadImage(assets.ImgMap).Data, cam.GetTranslation(op, 0, 0))

	for i, markerElement := 1, gameMap.Markers.Front(); markerElement != nil; i, markerElement = i+1, markerElement.Next() {
		textOp := &ebiten.DrawImageOptions{}
		marker := markerElement.Value.(*Marker)
		num := strconv.Itoa(i)
		centerText(num, textOp)

		marker.Draw(screen, cam, loader)
		text.DrawWithOptions(
			cam.Surface,
			num,
			loader.LoadFont(assets.FontSmall).Face,
			cam.GetTranslation(textOp, float64(marker.PosX), float64(marker.PosY)),
		)
	}
}

func NewGameMap(loader resource.Loader) *GameMap {
	sprite := loader.LoadImage(assets.ImgMap).Data
	width := float64(sprite.Bounds().Dx())
	height := float64(sprite.Bounds().Dy())

	markers := list.New()
	markers.PushBack(newMarker(420, 420, Selected))
	markers.PushBack(newMarker(400, 415, Locked))
	markers.PushBack(newMarker(340, 405, Locked))
	markers.PushBack(newMarker(320, 380, Locked))
	markers.PushBack(newMarker(370, 320, Locked))
	markers.PushBack(newMarker(250, 260, Locked))
	markers.PushBack(newMarker(200, 360, Locked))
	markers.PushBack(newMarker(190, 200, Locked))
	markers.PushBack(newMarker(150, 170, Locked))
	markers.PushBack(newMarker(250, 50, Locked))
	markers.PushBack(newMarker(310, 40, Locked))
	markers.PushBack(newMarker(400, 60, Locked))

	return &GameMap{
		width:         width,
		height:        height,
		Markers:       markers,
		CurrentMarker: markers.Front(),
	}
}

func (gameMap *GameMap) GetDimensions() (float64, float64) {
	return gameMap.width, gameMap.height
}

func (gameMap *GameMap) GetCurrentMarker() *Marker {
	return gameMap.CurrentMarker.Value.(*Marker)
}

func (gameMap *GameMap) GetNextMarker() *Marker {
	next := gameMap.CurrentMarker.Next()
	if next == nil {
		return nil
	}
	return next.Value.(*Marker)
}

func (gameMap *GameMap) GetPrevMarker() *Marker {
	prev := gameMap.CurrentMarker.Prev()
	if prev == nil {
		return nil
	}
	return prev.Value.(*Marker)
}
