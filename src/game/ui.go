package game

import (
	"github.com/timothy-ch-cheung/go-game-tween/assets"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"

	resource "github.com/quasilyte/ebitengine-resource"
)

type Callbacks struct {
	Next func(args *widget.ButtonClickedEventArgs)
	Prev func(args *widget.ButtonClickedEventArgs)
}

func createPrevBtnImg(loader *resource.Loader) *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:     assets.NineSliceImage(loader.LoadImage(assets.ImgPrevBtnIdle).Data, 54, 14),
		Hover:    assets.NineSliceImage(loader.LoadImage(assets.ImgPrevBtnHover).Data, 54, 14),
		Pressed:  assets.NineSliceImage(loader.LoadImage(assets.ImgPrevBtnPressed).Data, 54, 14),
		Disabled: assets.NineSliceImage(loader.LoadImage(assets.ImgPrevBtnDisabled).Data, 54, 14),
	}
}

func createNextBtnImg(loader *resource.Loader) *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:     assets.NineSliceImage(loader.LoadImage(assets.ImgNextBtnIdle).Data, 54, 14),
		Hover:    assets.NineSliceImage(loader.LoadImage(assets.ImgNextBtnHover).Data, 54, 14),
		Pressed:  assets.NineSliceImage(loader.LoadImage(assets.ImgNextBtnPressed).Data, 54, 14),
		Disabled: assets.NineSliceImage(loader.LoadImage(assets.ImgNextBtnDisabled).Data, 54, 14),
	}
}

func CreateUI(loader *resource.Loader, callbacks *Callbacks) *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(widget.Insets{
			Bottom: 5,
			Left:   10,
			Right:  10,
		}))),
	)
	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Spacing(10))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			VerticalPosition:   widget.AnchorLayoutPositionEnd,
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
		})),
	)
	rootContainer.AddChild(btnContainer)
	prevBtn := widget.NewButton(
		widget.ButtonOpts.Image(createPrevBtnImg(loader)),
		widget.ButtonOpts.ClickedHandler(callbacks.Prev),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true, MaxHeight: 20})),
	)
	btnContainer.AddChild(prevBtn)
	nextBtn := widget.NewButton(
		widget.ButtonOpts.Image(createNextBtnImg(loader)),
		widget.ButtonOpts.ClickedHandler(callbacks.Next),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true, MaxHeight: 20})),
	)
	btnContainer.AddChild(nextBtn)

	return &ebitenui.UI{
		Container: rootContainer,
	}
}
