//go:generate fyne bundle -package fyne -o bundled.go frame_mobile.svg
//go:generate fyne bundle -package fyne -o bundled.go -append frame_mobile_landscape.svg
//go:generate fyne bundle -package fyne -o bundled.go -append ../Icon.png

package fyne

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Theme = (*gameTheme)(nil)

type gameTheme struct {
	fyne.Theme
}

func newGameTheme() fyne.Theme {
	return &gameTheme{Theme: theme.DefaultTheme()}
}

func (t *gameTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameBackground {
		return &color.Gray{Y: 0xbd}
	}

	return t.Theme.Color(n, v)
}

type frame struct {
	output *canvas.Image

	power, tl, tr, bl, br *canvas.Circle
	b1, b2, b3, b4        *canvas.Rectangle
}

func newFrame(i *canvas.Image) *frame {
	return &frame{output: i}
}

func (f *frame) Layout(_ []fyne.CanvasObject, s fyne.Size) {
	scale := s.Width / 130.0

	f.output.Resize(fyne.NewSize(80*scale, 72*scale))
	f.output.Move(fyne.NewPos(25*scale, 14*scale))

	f.power.Resize(fyne.NewSize(4*scale, 4*scale))
	f.power.Move(fyne.NewPos(9.25*scale, 36.5*scale))

	f.tl.Resize(fyne.NewSize(8*scale, 8*scale))
	f.tl.Move(fyne.NewPos(2*scale, 2*scale))
	f.tr.Resize(fyne.NewSize(8*scale, 8*scale))
	f.tr.Move(fyne.NewPos(120*scale, 2*scale))
	f.bl.Resize(fyne.NewSize(8*scale, 8*scale))
	f.bl.Move(fyne.NewPos(2*scale, 90*scale))
	f.br.Resize(fyne.NewSize(40*scale, 40*scale))
	f.br.Move(fyne.NewPos(88*scale, 58*scale))

	f.b1.Move(fyne.NewPos(6*scale, 2*scale))
	f.b1.Resize(fyne.NewSize(118*scale, 8*scale))
	f.b2.Move(fyne.NewPos(2*scale, 6*scale))
	f.b2.Resize(fyne.NewSize(126*scale, 72*scale))
	f.b3.Move(fyne.NewPos(2*scale, 74*scale))
	f.b3.Resize(fyne.NewSize(106*scale, 20*scale))
	f.b4.Move(fyne.NewPos(6*scale, 90*scale))
	f.b4.Resize(fyne.NewSize(102*scale, 8*scale))
}

func (f *frame) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(130, 100)
}

func (f *frame) makeUI() fyne.CanvasObject {
	f.power = canvas.NewCircle(color.NRGBA{R: 0xd9, G: 0x10, B: 0x10, A: 0xff})

	frameCol := &color.NRGBA{R: 0x64, G: 0x66, B: 0x76, A: 0xff}
	f.tl = canvas.NewCircle(frameCol)
	f.tr = canvas.NewCircle(frameCol)
	f.bl = canvas.NewCircle(frameCol)
	f.br = canvas.NewCircle(frameCol)
	f.b1 = canvas.NewRectangle(frameCol)
	f.b2 = canvas.NewRectangle(frameCol)
	f.b3 = canvas.NewRectangle(frameCol)
	f.b4 = canvas.NewRectangle(frameCol)
	return container.New(f, f.tl, f.tr, f.bl, f.br, f.b1, f.b2, f.b3, f.b4, f.output, f.power)
}