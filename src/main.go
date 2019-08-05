package main

import (
	. "game"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"utils"
)

func main() {
	// GTK Init
	win, styleProvider, err := utils.GtkInit()
	if err != nil {
		log.Fatal("Failed to initialize GTK.")
	}

	// Create a new gamegrid widget to arrange child widgets
	uiGrid, _ := gtk.GridNew()
	uiGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	utils.AddStyleClassAndProvider(&uiGrid.Widget, styleProvider, "grid")

	// Add game gamegrid
	gameGrid := DrawGrid(styleProvider)
	uiGrid.Attach(gameGrid.Grid,1, 1, 9, 9)

	// Add new game button
	newGameBtn, _ := gtk.ButtonNewWithLabel("New Game")
	startNewGame := func() { NewDifficultySelectWindow(&gameGrid) }
	newGameBtn.Connect("clicked", startNewGame)
	utils.AddStyleClassAndProvider(&newGameBtn.Widget, styleProvider, "btn")

	uiGrid.Attach(newGameBtn, 0, 11, 2, 1)
	newGameBtn.SetHExpand(true)
	newGameBtn.SetVExpand(true)

	// Add exit button
	exitBtn, _ := gtk.ButtonNewWithLabel("Exit")
	gracefulExit := func() { os.Exit(0) }
	exitBtn.Connect("clicked", gracefulExit)
	utils.AddStyleClassAndProvider(&exitBtn.Widget, styleProvider, "btn")


	uiGrid.Attach(exitBtn, 9, 11, 2, 1)
	exitBtn.SetHExpand(true)
	exitBtn.SetVExpand(true)

	win.Add(uiGrid)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
