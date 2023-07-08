package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

type ExtraForm struct {
	window *gtk.Window
}

func newExtraForm(parent *gtk.ApplicationWindow) *ExtraForm {
	e := &ExtraForm{}

	e.window = getObject("extraWindow").(*gtk.Window)

	e.window.SetTitle("extra form")
	e.window.SetModal(true)
	e.window.SetTransientFor(parent)
	e.window.SetPosition(gtk.WIN_POS_CENTER_ON_PARENT)
	e.window.HideOnDelete()

	button := getObject("extraWindowCloseButton").(*gtk.Button)
	button.Connect("clicked", e.window.Hide)

	return e
}

func (e *ExtraForm) open() {
	e.window.ShowAll()
}
