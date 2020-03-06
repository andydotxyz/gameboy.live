package fyne

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/mobile"
	"fyne.io/fyne/widget"
)

type gameButton struct {
	widget.BaseWidget
	buttonCode uint
	lcd        *LCD
}

func (g *gameButton) TouchDown(*mobile.TouchEvent) {
	g.lcd.downCode(g.buttonCode)
}

func (g *gameButton) TouchUp(*mobile.TouchEvent) {
	g.lcd.upCode(g.buttonCode)
}

func (g *gameButton) TouchCancel(*mobile.TouchEvent) {
	g.lcd.upCode(g.buttonCode)
}

func (g *gameButton) CreateRenderer() fyne.WidgetRenderer {
	return &gameButtonRenderer{}
}

func newGameButton(lcd *LCD, code uint) *gameButton {
	b := &gameButton{lcd: lcd, buttonCode: code}
	b.ExtendBaseWidget(b)

	return b
}

type gameButtonRenderer struct {
}

func (r *gameButtonRenderer) Layout(fyne.Size) {
}

func (r *gameButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (r *gameButtonRenderer) Refresh() {
}

func (r *gameButtonRenderer) Objects() []fyne.CanvasObject {
	return nil
}

func (r *gameButtonRenderer) Destroy() {
}

func (r *gameButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}
