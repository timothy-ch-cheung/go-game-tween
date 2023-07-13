package game

import (
	"container/list"
	"image/color"

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

var textColour = &widget.ButtonTextColor{
	Idle:     color.Black,
	Disabled: color.Black,
}

var textPadding = widget.Insets{
	Top:    5,
	Bottom: 5,
	Left:   5,
	Right:  5,
}

func createBtnImg(loader *resource.Loader) *widget.ButtonImage {
	return &widget.ButtonImage{
		Idle:     assets.NineSliceImage(loader.LoadImage(assets.ImgBtnIdle).Data, 54, 14),
		Hover:    assets.NineSliceImage(loader.LoadImage(assets.ImgBtnHover).Data, 54, 14),
		Pressed:  assets.NineSliceImage(loader.LoadImage(assets.ImgBtnPressed).Data, 54, 14),
		Disabled: assets.NineSliceImage(loader.LoadImage(assets.ImgBtnDisabled).Data, 54, 14),
	}
}

func CreateUI(loader *resource.Loader, callbacks *Callbacks) *GameUI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(widget.Insets{
			Bottom: 5,
			Left:   5,
			Right:  5,
		}))),
	)
	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Spacing(2), widget.RowLayoutOpts.Direction(widget.DirectionVertical))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			VerticalPosition:   widget.AnchorLayoutPositionEnd,
			HorizontalPosition: widget.AnchorLayoutPositionStart,
		})),
	)
	rootContainer.AddChild(btnContainer)
	btnImg := createBtnImg(loader)
	nextBtn := widget.NewButton(
		widget.ButtonOpts.Image(btnImg),
		widget.ButtonOpts.Text("Next", loader.LoadFont(assets.FontDefault).Face, textColour),
		widget.ButtonOpts.TextPadding(textPadding),
		widget.ButtonOpts.ClickedHandler(callbacks.Next),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true, MaxHeight: 20})),
	)
	btnContainer.AddChild(nextBtn)
	prevBtn := widget.NewButton(
		widget.ButtonOpts.Image(btnImg),
		widget.ButtonOpts.Text("Prev", loader.LoadFont(assets.FontDefault).Face, textColour),
		widget.ButtonOpts.TextPadding(textPadding),
		widget.ButtonOpts.ClickedHandler(callbacks.Prev),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true, MaxHeight: 20})),
	)
	prevBtn.GetWidget().Disabled = true
	btnContainer.AddChild(prevBtn)

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
