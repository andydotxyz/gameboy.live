//go:generate fyne bundle -package fyne -o bundled.go frame.svg
//go:generate fyne bundle -package fyne -o bundled.go -append frame_mobile.svg
//go:generate fyne bundle -package fyne -o bundled.go -append frame_mobile_landscape.svg
//go:generate fyne bundle -package fyne -o bundled.go -append ../Icon.png

package fyne

import (
	"image/color"

	"fyne.io/fyne/v2"
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
