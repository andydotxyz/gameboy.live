package fyne

import (
	"fmt"
	"image"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"

	"github.com/andydotxyz/fynegameboy/util"
)

type LCD struct {
	pixels *[160][144][3]uint8
	screen *image.RGBA

	frame, output fyne.CanvasObject

	up, down, left, right fyne.CanvasObject
	start, sel, a, b      fyne.CanvasObject

	inputStatus *byte
	interrupt   bool
	title       string
}

func (lcd *LCD) Init(pixels *[160][144][3]uint8, title string) {
	lcd.pixels = pixels
	lcd.title = title
	log.Println("[Display] Initialize Fyne GUI display")
}

func (lcd *LCD) InitStatus(statusPointer *byte) {
	lcd.inputStatus = statusPointer
}

func (lcd *LCD) UpdateInput() bool {
	if lcd.interrupt {
		lcd.interrupt = false

		return true
	}

	return false
}

func (lcd *LCD) NewInput(b []byte) {
}

func (lcd *LCD) draw(w, h int) image.Image {
	i := 0
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			r, g, b := lcd.pixels[x][y][0], lcd.pixels[x][y][1], lcd.pixels[x][y][2]

			if r == 0xFF && g == 0xFF && b == 0xFF {
				lcd.screen.Pix[i] = 0x9b
				lcd.screen.Pix[i+1] = 0xbc
				lcd.screen.Pix[i+2] = 0x0f
			} else if r == 0xCC && g == 0xCC && b == 0xCC {
				lcd.screen.Pix[i] = 0x8b
				lcd.screen.Pix[i+1] = 0xac
				lcd.screen.Pix[i+2] = 0x0f
			} else if r == 0x77 && g == 0x77 && b == 0x77 {
				lcd.screen.Pix[i] = 0x30
				lcd.screen.Pix[i+1] = 0x62
				lcd.screen.Pix[i+2] = 0x30
			} else {
				lcd.screen.Pix[i] = 0x0f
				lcd.screen.Pix[i+1] = 0x38
				lcd.screen.Pix[i+2] = 0x0f
			}
			lcd.screen.Pix[i+3] = 0xff

			i += 4
		}
	}

	return lcd.screen
}

// Mapping from keys to GB index.
// Reference :https://github.com/Humpheh/goboy/blob/master/pkg/gbio/iopixel/pixels.go
var keyMap = map[fyne.KeyName]byte{
	// A button
	fyne.KeyX: 4,
	// B button
	fyne.KeyZ: 5,
	// SELECT button
	fyne.KeyBackspace: 6,
	// START button
	fyne.KeyReturn: 7,
	// RIGHT button
	fyne.KeyRight: 0,
	// LEFT button
	fyne.KeyLeft: 1,
	// UP button
	fyne.KeyUp: 2,
	// DOWN button
	fyne.KeyDown: 3,
}

func (lcd *LCD) downCode(num uint) {
	statusCopy := *lcd.inputStatus

	statusCopy = util.ClearBit(statusCopy, num)
	lcd.interrupt = true

	*lcd.inputStatus = statusCopy
}

func (lcd *LCD) buttonDown(ev *fyne.KeyEvent) {
	if offset, ok := keyMap[ev.Name]; ok {
		lcd.downCode(uint(offset))
	}
}

func (lcd *LCD) upCode(num uint) {
	statusCopy := *lcd.inputStatus

	statusCopy = util.SetBit(statusCopy, num)
	lcd.interrupt = true

	*lcd.inputStatus = statusCopy
}

func (lcd *LCD) buttonUp(ev *fyne.KeyEvent) {
	if offset, ok := keyMap[ev.Name]; ok {
		lcd.upCode(uint(offset))
	}
}

func (lcd *LCD) MinSize([]fyne.CanvasObject) fyne.Size {
	if fyne.CurrentDevice().IsMobile() {
		return fyne.NewSize(520, 800)
	}

	return fyne.NewSize(520, 400)
}

func (lcd *LCD) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	lcd.frame.Resize(size)

	xScale := float32(size.Width) / 520.0
	yScale := float32(size.Height) / 400.0
	if fyne.CurrentDevice().IsMobile() {
		yScale = float32(size.Height) / 2 / 400.0
	}

	abSize := fyne.NewSize(int(70*xScale), int(70*yScale))
	lcd.a.Resize(abSize)
	lcd.a.Move(fyne.NewPos(int(425*xScale), int(516*yScale)))
	lcd.b.Resize(abSize)
	lcd.b.Move(fyne.NewPos(int(328*xScale), int(565*yScale)))

	startSize := fyne.NewSize(int(90*xScale), int(20*yScale))
	lcd.start.Resize(startSize)
	lcd.start.Move(fyne.NewPos(int(230*xScale), int(725*yScale)))
	lcd.sel.Resize(startSize)
	lcd.sel.Move(fyne.NewPos(int(135*xScale), int(725*yScale)))

	dSize := fyne.NewSize(int(50*xScale), int(50*yScale))
	lcd.up.Resize(dSize)
	lcd.down.Resize(dSize)
	lcd.left.Resize(dSize)
	lcd.right.Resize(dSize)

	lcd.up.Move(fyne.NewPos(int(68*xScale), int(505*yScale)))
	lcd.down.Move(fyne.NewPos(int(68*xScale), int(605*yScale)))
	lcd.left.Move(fyne.NewPos(int(18*xScale), int(555*yScale)))
	lcd.right.Move(fyne.NewPos(int(118*xScale), int(555*yScale)))

	lcd.output.Resize(fyne.NewSize(int(320*xScale), int(296*yScale)))
	lcd.output.Move(fyne.NewPos(int(100*xScale), int(54*yScale)))
}

func (lcd *LCD) Run(drawSignal chan bool, onQuit func()) {
	a := app.New()
	a.SetIcon(resourceIconPng)
	win := a.NewWindow(fmt.Sprintf("GameBoy - %s", lcd.title))

	lcd.screen = image.NewRGBA(image.Rect(0, 0, 160, 144))
	lcd.output = canvas.NewRaster(lcd.draw)
	go func() {
		for {
			// drawSignal was sent by the emulator
			<-drawSignal

			canvas.Refresh(lcd.output)
		}
	}()

	if a.Driver().Device().IsMobile() {
		lcd.frame = canvas.NewImageFromResource(resourceFramemobileSvg)
	} else {
		lcd.frame = canvas.NewImageFromResource(resourceFrameSvg)
	}
	lcd.up = newGameButton(lcd, 2)
	lcd.down = newGameButton(lcd, 3)
	lcd.left = newGameButton(lcd, 1)
	lcd.right = newGameButton(lcd, 0)

	lcd.a = newGameButton(lcd, 4)
	lcd.b = newGameButton(lcd, 5)
	lcd.start = newGameButton(lcd, 7)
	lcd.sel = newGameButton(lcd, 6)

	if !a.Driver().Device().IsMobile() {
		lcd.a.Hide()
		lcd.b.Hide()
		lcd.start.Hide()
		lcd.sel.Hide()

		lcd.up.Hide()
		lcd.down.Hide()
		lcd.left.Hide()
		lcd.right.Hide()
	}

	content := fyne.NewContainerWithLayout(lcd, lcd.output, lcd.frame,
		lcd.a, lcd.b, lcd.start, lcd.sel, lcd.up, lcd.down, lcd.left, lcd.right)

	win.SetPadded(false)
	win.SetContent(content)
	if deskCanvas, ok := win.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(lcd.buttonDown)
		deskCanvas.SetOnKeyUp(lcd.buttonUp)
	}
	win.SetOnClosed(func() {
		onQuit()
		a.Quit()
	})
	win.ShowAndRun()
}
