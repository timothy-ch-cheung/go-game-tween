package main

import (
	"log"

	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
	"github.com/timothy-ch-cheung/go-game-tween/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
)

type Game struct {
	loader  *resource.Loader
	gameMap *game.GameMap
}

func (g *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.gameMap.Draw(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc

	assets.RegisterImageResources(loader)

	gameMap := game.NewGameMap(*loader)

	ebiten.SetWindowSize(config.ScreenWidth*4, config.ScreenHeight*4)
	ebiten.SetWindowTitle("Map Tween Demo")
	if err := ebiten.RunGame(&Game{loader: loader, gameMap: gameMap}); err != nil {
		log.Fatal(err)
	}
}
