package gui

import (
	_ "embed"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type MainForm struct {
	// Normally not empty
}

//go:embed assets/main.glade
var gladeFile string

var builder *gtk.Builder

// StartApplication : Opens the MainForm window
func (m MainForm) StartApplication(app *gtk.Application) {
	var err error

	builder, err = gtk.BuilderNewFromString(gladeFile)
	if err != nil {
		panic(err)
	}

	win := getObject("mainWindow").(*gtk.ApplicationWindow)
	win.SetApplication(app)
	win.SetTitle("main window")
	win.SetSizeRequest(320, 200)
	win.Connect("delete-event", win.Close)

	openFormButton := getObject("openFormButton").(*gtk.Button)
	openFormButton.Connect(
		"clicked", func() {
			e := newExtraForm(win)
			e.open()
		},
	)

	win.ShowAll()
}

func getObject(name string) glib.IObject {
	obj, err := builder.GetObject(name)
	if err != nil {
		panic(err)
	}
	return obj
}
