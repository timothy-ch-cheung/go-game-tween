package main

import (
	"log"

	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
	"github.com/timothy-ch-cheung/go-game-tween/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

type Game struct {
	loader  *resource.Loader
	gameMap *game.GameMap
	cam     *ebitenCamera.Camera
}

func (g *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.gameMap.Draw(screen, game.cam, game.loader)
	game.cam.Blit(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc

	assets.RegisterImageResources(loader)

	cam := ebitenCamera.NewCamera(config.ScreenWidth, config.ScreenHeight, 0, 0, 0, 1)

	gameMap := game.NewGameMap(*loader)
	cam.SetPosition(gameMap.GetInitialPos())

	ebiten.SetWindowSize(config.ScreenWidth*4, config.ScreenHeight*4)
	ebiten.SetWindowTitle("Map Tween Demo")
	if err := ebiten.RunGame(&Game{loader: loader, gameMap: gameMap, cam: cam}); err != nil {
		log.Fatal(err)
	}
}
