package main

import (
	"log"
	"time"

	"github.com/timothy-ch-cheung/go-game-tween/assets"
	"github.com/timothy-ch-cheung/go-game-tween/config"
	"github.com/timothy-ch-cheung/go-game-tween/game"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	scale = 4
)

var currentTime = time.Now()
var delta float32

type Game struct {
	loader           *resource.Loader
	gameMap          *game.GameMap
	gameUI           *game.GameUI
	cam              *ebitenCamera.Camera
	cameraController *game.CameraController
}

func (game *Game) Update() error {
	newTime := time.Now()
	delta = float32(newTime.UnixMicro() - currentTime.UnixMicro())
	currentTime = newTime

	game.cameraController.Update(float32(delta))
	game.gameUI.Update(game.cameraController.IsCameraMoving, game.gameMap.CurrentMarker)
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.cam.Surface.Clear()
	game.gameMap.Draw(screen, game.cam, game.loader)
	game.cam.Blit(screen)
	game.gameUI.Draw(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc

	assets.RegisterImageResources(loader)
	assets.RegisterFontResources(loader)

	cam := ebitenCamera.NewCamera(config.ScreenWidth, config.ScreenHeight, 0, 0, 0, 1)

	gameMap := game.NewGameMap(*loader)

	width, height := gameMap.GetDimensions()
	cameraController := game.NewCameraController(cam, width, height)
	cam.SetPosition(cameraController.GetCameraPositionFromMarker(gameMap.GetCurrentMarker()))

	gameUI := game.CreateUI(loader, &game.Callbacks{
		Prev: func(args *widget.ButtonClickedEventArgs) {
			gameMap.CurrentMarker.Value.(*game.Marker).State = game.Idle
			prevMarkerElement := gameMap.CurrentMarker.Prev()
			prevMarker := prevMarkerElement.Value.(*game.Marker)
			prevMarker.State = game.Selected
			if prevMarkerElement.Prev() == nil {
				args.Button.GetWidget().Disabled = true
			}
			cameraController.InitiateMove(float64(prevMarker.PosX), float64(prevMarker.PosY))
			gameMap.CurrentMarker = prevMarkerElement
		},
		Next: func(args *widget.ButtonClickedEventArgs) {
			gameMap.CurrentMarker.Value.(*game.Marker).State = game.Idle
			nextMarkerElement := gameMap.CurrentMarker.Next()
			nextMarker := nextMarkerElement.Value.(*game.Marker)
			nextMarker.State = game.Selected
			if nextMarkerElement.Next() == nil {
				args.Button.GetWidget().Disabled = true
			}
			cameraController.InitiateMove(float64(nextMarker.PosX), float64(nextMarker.PosY))
			gameMap.CurrentMarker = nextMarkerElement
		},
	})

	ebiten.SetWindowSize(config.ScreenWidth*scale, config.ScreenHeight*scale)
	ebiten.SetWindowTitle("Map Tween Demo")

	game := &Game{
		loader:           loader,
		gameMap:          gameMap,
		gameUI:           gameUI,
		cam:              cam,
		cameraController: cameraController,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
