// +build android

package app // import "golang.org/x/mobile/app"

import fynemobile "github.com/fyne-io/mobile/app"

func RunOnJVM(f func(vm, env, ctx uintptr) error) error {
	return fynemobile.RunOnJVM(f)
}
