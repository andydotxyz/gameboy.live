#!/bin/sh

cd `dirname $0`
fyne bundle -package fyne frame.svg > bundled.go
fyne bundle -append frame_mobile.svg >> bundled.go
fyne bundle -append ../Icon.png >> bundled.go

