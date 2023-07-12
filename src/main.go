package main

import (
	"log"

	"github.com/timothy-ch-cheung/go-game-tween/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	screenWidth  = 160
	screenHeight = 160
)

type Game struct {
	loader *resource.Loader
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mapImg := g.loader.LoadImage(assets.ImgMap).Data
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(mapImg.Bounds().Dx())+screenWidth, -float64(mapImg.Bounds().Dy())+screenHeight)
	screen.DrawImage(mapImg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc

	assets.RegisterImageResources(loader)

	println(loader.LoadImage(assets.ImgMap).DefaultFrameHeight)
	println(loader.LoadImage(assets.ImgMap).DefaultFrameWidth)

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("Map Tween Demo")
	if err := ebiten.RunGame(&Game{loader: loader}); err != nil {
		log.Fatal(err)
	}
}
