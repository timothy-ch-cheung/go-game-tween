package game

import (
	"math"

	ebitenCamera "github.com/melonfunction/ebiten-camera"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/timothy-ch-cheung/go-game-tween/config"
)

const (
	moveTime = 500000
)

type cameraBounds struct {
	xMin float64
	xMax float64
	yMin float64
	yMax float64
}

type CameraController struct {
	camera         *ebitenCamera.Camera
	bounds         cameraBounds
	IsCameraMoving bool
	xTween         *gween.Tween
	isFinishedX    bool
	yTween         *gween.Tween
	isFinishedY    bool
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
		camera:         camera,
		bounds:         bounds,
		IsCameraMoving: false,
	}
}

func (cameraController *CameraController) SetCameraPosition(marker *Marker) {
	cameraController.camera.SetPosition(cameraController.GetCameraPositionFromMarker(marker))
}

func (cameraController *CameraController) GetCameraPositionFromCoord(targetX float64, targetY float64) (float64, float64) {
	x := math.Min(math.Max(targetX, cameraController.bounds.xMin), cameraController.bounds.xMax)
	y := math.Min(math.Max(targetY, cameraController.bounds.yMin), cameraController.bounds.yMax)
	return x, y
}

func (cameraController *CameraController) GetCameraPositionFromMarker(marker *Marker) (float64, float64) {
	return cameraController.GetCameraPositionFromCoord(float64(marker.PosX), float64(marker.PosY))
}

func (cameraController *CameraController) Update(delta float32) {
	if cameraController.IsCameraMoving {
		var newX, newY = cameraController.camera.X, cameraController.camera.Y
		if !cameraController.isFinishedX {
			x, isFinishedX := cameraController.xTween.Update(delta)
			newX = float64(x)
			cameraController.isFinishedX = isFinishedX
		}
		if !cameraController.isFinishedY {
			y, isFinishedY := cameraController.yTween.Update(delta)
			newY = float64(y)
			cameraController.isFinishedY = isFinishedY
		}

		cameraController.camera.SetPosition(float64(newX), float64(newY))
		if cameraController.isFinishedX && cameraController.isFinishedY {
			cameraController.IsCameraMoving = false
			cameraController.isFinishedX = false
			cameraController.isFinishedY = false
		}
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func (cameraController *CameraController) InitiateMove(targetX float64, targetY float64) {
	x, y := cameraController.GetCameraPositionFromCoord(targetX, targetY)
	if almostEqual(x, cameraController.camera.X) && almostEqual(y, cameraController.camera.Y) {
		return
	}
	cameraController.IsCameraMoving = true
	cameraController.xTween = gween.New(float32(cameraController.camera.X), float32(x), moveTime, ease.InOutCubic)
	cameraController.yTween = gween.New(float32(cameraController.camera.Y), float32(y), moveTime, ease.InOutCubic)
}
