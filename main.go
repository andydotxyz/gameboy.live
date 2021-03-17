package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	fyneAPI "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"github.com/andydotxyz/fynegameboy/fyne"
	"github.com/andydotxyz/fynegameboy/gb"
)

var (
	h bool

	romPath string
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

func loadRom(romPath string) ([]byte, fyneAPI.URI) {
	data, err := ioutil.ReadFile(romPath)
	if err != nil {
		log.Println("ERROR", err)
		return nil, storage.NewURI("")
	}

	return data, storage.NewURI("file://" + romPath)
}

func newCore(d *fyne.LCD) *gb.Core {
	core := new(gb.Core)

	core.FPS = FPS
	core.Clock = 4194304
	core.Debug = Debug
	core.DisplayDriver = d
	core.Controller = d
	core.DrawSignal = make(chan bool)
	core.SpeedMultiple = 0
	core.ToggleSound = SoundOn

	return core
}

func startGUI() {
	d := fyne.NewDriver()

	uri := storage.NewURI(fyneAPI.CurrentApp().Preferences().String("RomURI"))

	var data []byte
	if romPath == "" && uri != nil && uri.String() != "" {
		read, err := storage.OpenFileFromURI(uri)
		log.Println("err", err)
		data, err = ioutil.ReadAll(read)
		log.Println("err", err, len(data))
	}

	var core *gb.Core
	start := func() {
		core = newCore(d)
		d.DrawSignal = core.DrawSignal

		if data != nil && len(data) > 0 {
			core.Init(data, uri)
		} else {
			core.Init(loadRom(romPath))
		}

		go core.Run()
	}
	start()

	d.Reset = func() {
		core.Exit = true
		start()
	}
	d.Open = func(r fyneAPI.URIReadCloser) {
		core.Exit = true
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			fyneAPI.LogError("Unable to load ROM", err)
			return
		}
		_ = r.Close()
		data = bytes
		start()
	}

	d.Run(core.DrawSignal, func() {
		core.SaveRAM()
	})
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	if len(flag.Args()) == 1 { // probably a ROM parameter
		romPath, _ = filepath.Abs(flag.Arg(0))
	}

	startGUI()
}
