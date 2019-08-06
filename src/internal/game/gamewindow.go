package game

import (
	"github.com/gotk3/gotk3/gtk"
	"internal/utils"
	"strconv"
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

func (gw *GameWindow) Close() {
	gw.active = false
	gw.window.Close()
}

func NewNumberSelectWindow(g *GameGrid, node *Node) {
	win, styleProvider := utils.NewWindow("Choose number")

	numberGrid, _ := gtk.GridNew()
	numberGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	utils.AddStyleClassAndProvider(&numberGrid.Widget, styleProvider, "numbergrid")

	for x := 0; x < 9; x++ {
		evBox, _ := gtk.EventBoxNew()

		labelText := strconv.FormatInt(int64(x+1), 10)

		lab, _ := utils.CreateLabel(labelText)

		// Add CSS classes to node
		utils.AddStyleClassAndProvider(&lab.Widget, styleProvider, "numbergrid-node")

		// Create OnClickEvent
		evBox.Add(lab)
		evBox.Connect("button_press_event", func() { g.NumberSelect(labelText, node); g.Window.Close() })

		numberGrid.Attach(evBox, x % 3, x / 3 + 1, 1, 1)
	}

	evBox, _ := gtk.EventBoxNew()

	lab, _ := utils.CreateLabel("Clear")

	// Add CSS classes to node
	utils.AddStyleClassAndProvider(&lab.Widget, styleProvider, "numbergrid-node")

	// Create OnClickEvent
	evBox.Add(lab)
	evBox.Connect("button_press_event", func() { g.NumberSelect("", node); g.Window.Close() })

	numberGrid.Attach(evBox, 0, 4, 3, 1)

	win.Add(numberGrid);

	g.Window.Launch(win)
}

func NewDifficultySelectWindow(g *GameGrid) {
	win, styleProvider := utils.NewWindow("Choose difficulty")

	diffGrid, _ := gtk.GridNew()
	diffGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	utils.AddStyleClassAndProvider(&diffGrid.Widget, styleProvider, "numbergrid")

	for x := 1; x < 4; x++ {
		evBox, _ := gtk.EventBoxNew()

		labelText := "Easy"

		if x == 2 {
			labelText = "Medium"
		}

		if x == 3 {
			labelText = "Hard"
		}

		lab, _ := utils.CreateLabel(labelText)

		// Add CSS classes to node
		utils.AddStyleClassAndProvider(&lab.Widget, styleProvider, "numbergrid-node")


		// Create OnClickEvent
		evBox.Add(lab)
		diff := x // Save x value
		evBox.Connect("button_press_event", func() { g.CreateNewPuzzle(diff); g.Window.Close() })

		diffGrid.Attach(evBox, 0, x - 1, 1, 1)
	}

	win.Add(diffGrid);

	g.Window.Launch(win)
}

