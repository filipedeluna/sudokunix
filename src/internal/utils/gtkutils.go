package utils

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

const CSS_FOLDER string = "src/assets/css/style.css"

// GTK extensions
func GtkInit() (*gtk.Window, *gtk.CssProvider, error) {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a style provider
	styleProvider, _ := gtk.CssProviderNew();
	styleProvider.LoadFromPath(CSS_FOLDER)

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
	styleProvider.LoadFromPath(CSS_FOLDER)

	// Set the default window size.
	win.SetDefaultSize(200, 150)

	return win, styleProvider
}

func AddStyleClassAndProvider(actionable *gtk.Widget, styleProvider gtk.IStyleProvider, class string) {
	// Add CSS classes to node
	ctx, _ := actionable.GetStyleContext()
	ctx.AddClass(class)
	ctx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func CreateLabel(text string) (*gtk.Label, error) {
	lab, err := gtk.LabelNew(text)
	lab.SetJustify(gtk.JUSTIFY_CENTER)
	lab.SetHExpand(true)
	lab.SetVExpand(true)

	return lab, err
}