package main

import (
	"github.com/gotk3/gotk3/gtk"
	. "internal/game"
	"internal/utils"
	"log"
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
	newGameBtn.SetHAlign(gtk.ALIGN_CENTER)
	newGameBtn.Connect("clicked", startNewGame)
	utils.AddStyleClassAndProvider(&newGameBtn.Widget, styleProvider, "btn")

	uiGrid.Attach(newGameBtn, 1, 11, 2, 1)

	// Add candidate mode checkbox
	candModeChkbox, _ := gtk.CheckButtonNewWithLabel("Candidate Mode")
	candModeChkbox.SetHAlign(gtk.ALIGN_CENTER)
	candModeChkbox.Connect("clicked", func() { gameGrid.SetCandidateMode(); })
	utils.AddStyleClassAndProvider(&candModeChkbox.Widget, styleProvider, "btn")

	uiGrid.Attach(candModeChkbox, 5, 11, 2, 1)

	win.Add(uiGrid)

	// Set the default window size.
	win.SetDefaultSize(600, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
