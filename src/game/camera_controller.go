package game

import (
	"math"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	"github.com/tanema/gween"
	"github.com/timothy-ch-cheung/go-game-tween/config"
)

type cameraBounds struct {
	xMin float64
	xMax float64
	yMin float64
	yMax float64
}

type CameraController struct {
	camera         *ebitenCamera.Camera
	isCameraMoving bool
	tween          gween.Tween
	bounds         cameraBounds
	posX           float64
	posY           float64
}

func NewCameraController(camera *ebitenCamera.Camera, mapWidth float64, mapHeight float64) *CameraController {
	xBorder := float64(config.ScreenWidth / 2)
	yBorder := float64(config.ScreenHeight / 2)
	bounds := cameraBounds{
		xMin: xBorder,
		xMax: mapWidth - xBorder,
		yMin: yBorder,
		yMax: mapHeight - yBorder,
	}

	return &CameraController{
		camera: camera,
		bounds: bounds,
	}
}

func (cameraController *CameraController) GetCameraPosition(marker *Marker) (float64, float64) {
	x := math.Min(math.Max(float64(marker.posX), cameraController.bounds.xMin), cameraController.bounds.xMax)
	y := math.Min(math.Max(float64(marker.posY), cameraController.bounds.yMin), cameraController.bounds.yMax)
	return x, y
}
