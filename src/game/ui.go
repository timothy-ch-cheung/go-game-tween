package game

import (
	"container/list"

	"github.com/timothy-ch-cheung/go-game-tween/assets"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	resource "github.com/quasilyte/ebitengine-resource"
)

type Callbacks struct {
	Next func(args *widget.ButtonClickedEventArgs)
	Prev func(args *widget.ButtonClickedEventArgs)
}

type Buttons struct {
	prev *widget.Button
	next *widget.Button
}

type GameUI struct {
	ui               *ebitenui.UI
	interfaceEnabled bool
	isCameraMoving   bool
	buttons          *Buttons
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

func CreateUI(loader *resource.Loader, callbacks *Callbacks) *GameUI {
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
	prevBtn.GetWidget().Disabled = true
	btnContainer.AddChild(prevBtn)
	nextBtn := widget.NewButton(
		widget.ButtonOpts.Image(createNextBtnImg(loader)),
		widget.ButtonOpts.ClickedHandler(callbacks.Next),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true, MaxHeight: 20})),
	)
	btnContainer.AddChild(nextBtn)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	buttons := &Buttons{
		prev: prevBtn,
		next: nextBtn,
	}
	return &GameUI{
		ui:               ui,
		interfaceEnabled: true,
		isCameraMoving:   false,
		buttons:          buttons,
	}
}

func (gameUI *GameUI) Update(isCameraMoving bool, currentMarker *list.Element) {
	gameUI.ui.Update()
	if isCameraMoving == gameUI.isCameraMoving {
		return
	}

	gameUI.isCameraMoving = isCameraMoving
	if isCameraMoving {
		gameUI.buttons.prev.GetWidget().Disabled = true
		gameUI.buttons.next.GetWidget().Disabled = true
	} else {
		gameUI.buttons.prev.GetWidget().Disabled = currentMarker.Prev() == nil
		gameUI.buttons.next.GetWidget().Disabled = currentMarker.Next() == nil
	}
}

func (gameUI *GameUI) Draw(screen *ebiten.Image) {
	gameUI.ui.Draw(screen)
}

func (gameUI *GameUI) IsInterfaceEnabled() bool {
	return gameUI.interfaceEnabled
}
