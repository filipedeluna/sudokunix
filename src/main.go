package main

import (
	. "game/gamegrid"
	. "game/gamewindow"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"utils/gtkutils"
)

type GameStatus struct {
	Won bool
	Active bool // controls if user can insert/remove numbers
	gameWindow GameWindow // These are popup windows. There should only be one active at any given time (eg. pick number)
}

func main() {
	// Initialize the game status
	gameStatus := GameStatus{ false, false, GameWindow { } }

	// GTK Init
	win, styleProvider, err := gtkutils.GtkInit()
	if err != nil {
		log.Fatal("Failed to initialize GTK.")
	}

	// Create a new gamegrid widget to arrange child widgets
	uiGrid, _ := gtk.GridNew()
	uiGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	uiGridCtx, _ := uiGrid.GetStyleContext()
	uiGridCtx.AddClass("grid")
	uiGridCtx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// Add game gamegrid
	gameGrid := DrawGrid(styleProvider)
	uiGrid.Attach(gameGrid.Grid,1, 1, 9, 9)

	// Add new game button
	newGameBtn, _ := gtk.ButtonNewWithLabel("New Game")
	diff := 1 // temporary
	startNewGame := func() {
		gameGrid.CreateNewPuzzle(diff, &gameStatus.gameWindow) // window sent as ref
		gameStatus.Active = true
	}
	newGameBtn.Connect("clicked", startNewGame)
	newGameBtnCtx, _ := newGameBtn.GetStyleContext()
	newGameBtnCtx.AddClass("btn")
	newGameBtnCtx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	uiGrid.Attach(newGameBtn, 0, 11, 3, 1)
	newGameBtn.SetHExpand(true)
	newGameBtn.SetVExpand(true)

	// Add exit button
	exitBtn, _ := gtk.ButtonNewWithLabel("Exit")
	gracefulExit := func() { os.Exit(0) }
	exitBtn.Connect("clicked", gracefulExit)
	exitBtnCtx, _ := exitBtn.GetStyleContext()
	exitBtnCtx.AddClass("btn")
	exitBtnCtx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	uiGrid.Attach(exitBtn, 8, 11, 3, 1)
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
