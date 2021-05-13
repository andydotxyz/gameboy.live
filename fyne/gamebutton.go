package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type gameButton struct {
	widget.BaseWidget
	buttonCode uint
	lcd        *LCD
	bg         fyne.CanvasObject
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
	return &gameButtonRenderer{objects: []fyne.CanvasObject{g.bg}}
}

func newGameButton(lcd *LCD, code uint, obj fyne.CanvasObject) *gameButton {
	b := &gameButton{lcd: lcd, buttonCode: code, bg: obj}
	b.ExtendBaseWidget(b)

	return b
}

type gameButtonRenderer struct {
	objects []fyne.CanvasObject
}

func (r *gameButtonRenderer) Layout(s fyne.Size) {
	r.objects[0].Resize(s)
}

func (r *gameButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (r *gameButtonRenderer) Refresh() {
}

func (r *gameButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *gameButtonRenderer) Destroy() {
}
