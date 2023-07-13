package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/timothy-ch-cheung/go-game-tween/assets"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type MarkerState int

const (
	Idle MarkerState = iota
	Selected
	Locked
)

type Marker struct {
	PosX  int
	PosY  int
	State MarkerState
}

func newMarker(x int, y int, state MarkerState) *Marker {
	return &Marker{
		PosX:  x,
		PosY:  y,
		State: state,
	}
}

func (marker *Marker) getSprite(loader *resource.Loader) *ebiten.Image {
	switch marker.State {
	case Idle:
		return loader.LoadImage(assets.ImgMarkerIdle).Data
	case Selected:
		return loader.LoadImage(assets.ImgMarkerSelected).Data
	case Locked:
		return loader.LoadImage(assets.ImgMarkerLocked).Data
	default:
		return nil
	}
}

func (marker *Marker) Draw(screen *ebiten.Image, cam *ebitenCamera.Camera, loader *resource.Loader) {
	op := &ebiten.DrawImageOptions{}
	cam.Surface.DrawImage(marker.getSprite(loader), cam.GetTranslation(op, float64(marker.PosX), float64(marker.PosY)))
}
