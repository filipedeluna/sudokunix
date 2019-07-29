package gtkutils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// GTK extensions
func GtkInit() (*gtk.Window, *gtk.CssProvider, error) {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a style provider
	styleProvider, _ := gtk.CssProviderNew();
	styleProvider.LoadFromPath("src/style.css")

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Sudoku")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return win, styleProvider, err
}

func NewWindow(title string) (*gtk.Window, *gtk.CssProvider) {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	win.SetTitle(title)

	// Create a style provider
	styleProvider, _ := gtk.CssProviderNew();
	styleProvider.LoadFromPath("src/style.css")

	// Set the default window size.
	win.SetDefaultSize(200, 150)

	return win, styleProvider
}