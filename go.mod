module github.com/andydotxyz/fynegameboy

go 1.12

require (
	fyne.io/fyne/v2 v2.0.0
	github.com/faiface/beep v1.0.3-0.20200712202812-d836f29bdc50
)

replace golang.org/x/mobile => github.com/fyne-io/gomobile-bridge v0.0.2

replace github.com/go-gl/glfw/v3.3/glfw v0.0.0-20200625191551-73d3c3675aa3 => github.com/fyne-io/glfw/v3.3/glfw v0.0.0-20201123143003-f2279069162d
