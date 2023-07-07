package main

import (
	"os"

	"github.com/hultan/testGtk/internal/gui"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	ApplicationId    = "se.softteam.gui"
	ApplicationFlags = glib.APPLICATION_FLAGS_NONE
)

func main() {
	application, err := gtk.ApplicationNew(ApplicationId, ApplicationFlags)
	if err != nil {
		panic("Failed to create GTK Application : " + err.Error())
	}

	mainForm := &gui.MainForm{}
	application.Connect("activate", mainForm.StartApplication)
	if err != nil {
		panic("Failed to connect Application.Activate event : " + err.Error())
	}

	os.Exit(application.Run(nil))
}
