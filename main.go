package main

import (
	"flag"

	"github.com/andydotxyz/fynegameboy/driver"
	"github.com/andydotxyz/fynegameboy/fyne"
	"github.com/andydotxyz/fynegameboy/gb"
)

var (
	h bool

	ROMPath string
	SoundOn bool
	FPS     int
	Debug   bool
)

func init() {
	flag.BoolVar(&h, "h", false, "This help")
	flag.BoolVar(&SoundOn, "m", true, "Turn on sound in GUI mode")
	flag.BoolVar(&Debug, "d", false, "Use Debugger in GUI mode")
	flag.IntVar(&FPS, "f", 60, "Set the `FPS` in GUI mode")
}

func startGUI(screen driver.DisplayDriver, control driver.ControllerDriver) {
	core := new(gb.Core)
	core.FPS = FPS
	core.Clock = 4194304
	core.Debug = Debug
	core.DisplayDriver = screen
	core.Controller = control
	core.DrawSignal = make(chan bool)
	core.SpeedMultiple = 0
	core.ToggleSound = SoundOn
	core.Init(ROMPath)

	go core.Run()
	screen.Run(core.DrawSignal, func() {
		core.SaveRAM()
	})
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	driver := new(fyne.LCD)
	startGUI(driver, driver)
}
