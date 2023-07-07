package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

type ExtraForm struct {
	window *gtk.Window
}

func (e *ExtraForm) open(parent *gtk.ApplicationWindow) {
	e.window = getObject("extraWindow").(*gtk.Window)

	e.window.SetTitle("extra form")
	e.window.SetModal(true)
	e.window.SetTransientFor(parent)
	e.window.SetPosition(gtk.WIN_POS_CENTER_ON_PARENT)

	e.window.Connect("delete-event", func() {
		e.window.Destroy()
		e.window = nil
	})

	button := getObject("extraWindowCloseButton").(*gtk.Button)
	button.Connect("clicked", e.window.Hide)

	e.window.ShowAll()
}
