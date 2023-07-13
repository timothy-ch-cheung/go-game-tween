package main

import (
	"log"

	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
	"github.com/timothy-ch-cheung/go-game-tween/game"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	scale = 4
)

type Game struct {
	loader  *resource.Loader
	gameMap *game.GameMap
	ui      *ebitenui.UI
	cam     *ebitenCamera.Camera
}

func (game *Game) Update() error {
	game.ui.Update()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.gameMap.Draw(screen, game.cam, game.loader)
	game.cam.Blit(screen)
	game.ui.Draw(screen)
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

	width, height := gameMap.GetDimensions()
	cameraController := game.NewCameraController(cam, width, height)
	cam.SetPosition(cameraController.GetCameraPosition(gameMap.GetCurrentMarker()))

	ui := game.CreateUI(loader, &game.Callbacks{
		Prev: func(args *widget.ButtonClickedEventArgs) {},
		Next: func(args *widget.ButtonClickedEventArgs) {},
	})

	ebiten.SetWindowSize(config.ScreenWidth*scale, config.ScreenHeight*scale)
	ebiten.SetWindowTitle("Map Tween Demo")

	game := &Game{
		loader:  loader,
		gameMap: gameMap,
		ui:      ui,
		cam:     cam,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
