package game

import (
	"github.com/gotk3/gotk3/gtk"
)

type GameWindow struct {
	active bool
	window *gtk.Window
}

func (gw *GameWindow) Launch(newWindow *gtk.Window) {
	if gw.active {
		gw.window.Close()
		gw.active = false
	}

	gw.window = newWindow
	gw.window.ShowAll()
	gw.active = true
}





