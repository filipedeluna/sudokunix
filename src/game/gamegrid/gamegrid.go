package gamegrid

import (
	. "game/gamelogic"
	. "game/gamewindow"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type GameGrid struct {
	Grid  *gtk.Grid
	Nodes [N_OF_LINES][N_OF_LINES]Node
}

type Node struct {
	Label *gtk.Label
	EventBox *gtk.EventBox
	Signal glib.SignalHandle
}

func DrawGrid(styleProvider *gtk.CssProvider) GameGrid {
	// create grid value matrix
	nodes := [N_OF_LINES][N_OF_LINES]Node{}

	// Create a new gamegrid widget to arrange child widgets
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	gridCtx, _ := grid.GetStyleContext()
	gridCtx.AddClass("gamegrid")
	gridCtx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	totalValsInserted := 0
	totalValsInsertedPerSquare := [9]int{}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			evBox, _:= gtk.EventBoxNew()

			// Generate values for label
			lab, _ := gtk.LabelNew("")
			lab.SetJustify(gtk.JUSTIFY_CENTER)
			lab.SetHExpand(true)
			lab.SetVExpand(true)

			totalValsInserted++;
			totalValsInsertedPerSquare[y % 3 + x / 3 + 1]++

			// Add CSS classes to node
			ctx, _ := lab.GetStyleContext()
			ctx.AddClass("gamegrid-node")
			addGridNodeBorders(ctx, x, y)
			ctx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

			evBox.Add(lab)

			grid.Attach(evBox, x, y, 1, 1)

			nodes[x][y] = Node{ lab, evBox, 0 }
		}
	}

	gameGrid := GameGrid {
		Grid:  grid,
		Nodes: nodes,
	}

	return gameGrid
}

func addGridNodeBorders(ctx *gtk.StyleContext, x int, y int) {
	if y == 2 || y == 5 {
		ctx.AddClass("gamegrid-node--bottomborder")
	}

	if x == 3 || x == 6 {
		ctx.AddClass("gamegrid-node--leftborder")
	}
}

func (g *GameGrid) CreateNewPuzzle(diff int, window *GameWindow) {
	newPuzzle := GenerateNewPuzzle(diff);

	i := 0

	for x := 0; x < N_OF_LINES; x++ {
		for y := 0; y < N_OF_LINES; y++ {
			if newPuzzle[i] != '0' {
				g.Nodes[x][y].Label.SetText(string(newPuzzle[i]))
				g.Nodes[x][y].SetInactive()
			} else {
				g.Nodes[x][y].Label.SetText("")
				g.Nodes[x][y].SetActive(window)
			}

			i++
		}
	}
}

func (n *Node) SetActive(window *GameWindow) {
	if n.Signal != 0 {
		n.EventBox.HandlerDisconnect(n.Signal)
		n.Signal = 0
	}

	n.Signal, _ = n.EventBox.Connect("button_press_event", func() { LaunchNumberSelectWindow(n.Label, window) })

	ctx, _ := n.Label.GetStyleContext()
	ctx.RemoveClass("gamegrid-node--inactive")
}

func (n *Node) SetInactive() {
	if n.Signal != 0 {
		n.EventBox.HandlerDisconnect(n.Signal)
		n.Signal = 0
	}

	ctx, _ := n.Label.GetStyleContext()
	ctx.AddClass("gamegrid-node--inactive")
}

func (n Node) SetWrong() {
	ctx, _ := n.Label.GetStyleContext()
	ctx.AddClass("gamegrid-node--wrong")
}

func (n Node) UnsetWrong() {
	ctx, _ := n.Label.GetStyleContext()
	ctx.RemoveClass("gamegrid-node--wrong")
}