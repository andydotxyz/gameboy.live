#!/bin/sh

if [[ "$#" -ne "1" ]]; then
	echo "Script takes a single parameter, the rom to bundle"
	exit 1
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
NAME=`echo "$(basename $1)" | cut -d'.' -f1`

fyne bundle -package gb -name romResource $1 > $DIR/gb/rom.go
echo "Bundled $NAME"

