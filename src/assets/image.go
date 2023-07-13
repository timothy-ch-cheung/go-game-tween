package assets

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
)

func NineSliceImage(img *ebiten.Image, centerWidth, centerHeight int) *image.NineSlice {
	width, height := img.Size()
	return image.NewNineSlice(img,
		[3]int{(width - centerWidth) / 2, centerWidth, width - (width-centerWidth)/2 - centerWidth},
		[3]int{(height - centerHeight) / 2, centerHeight, height - (height-centerHeight)/2 - centerHeight})
}
