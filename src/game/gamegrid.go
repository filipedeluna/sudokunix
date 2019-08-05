package game

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
	"utils"
)

type GameGrid struct {
	Grid   *gtk.Grid
	Nodes  [N_OF_LINES][N_OF_LINES]Node
	Window GameWindow // These are popup windows. There should only be one active at any given time (eg. pick number)
}

type Node struct {
	Label *gtk.Label
	Value int
	X int
	Y int
	isWrong bool
	isActive bool
	EventBox *gtk.EventBox
	Signal glib.SignalHandle
}

func DrawGrid(styleProvider *gtk.CssProvider) GameGrid {
	// create grid value matrix
	nodes := [N_OF_LINES][N_OF_LINES]Node{}

	// Create a new gamegrid widget to arrange child widgets
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	utils.AddStyleClassAndProvider(&grid.Widget, styleProvider, "gamegrid")

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			evBox, _:= gtk.EventBoxNew()

			// Generate values for label
			lab, _ := utils.CreateLabel("")

			// Add CSS classes to node
			ctx, _ := lab.GetStyleContext()
			ctx.AddClass("gamegrid-node")
			addGridNodeBorders(ctx, x, y)
			ctx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

			evBox.Add(lab)

			grid.Attach(evBox, x, y, 1, 1)

			nodes[x][y] = Node{ lab, 0,x, y,false, false, evBox, 0 }
		}
	}

	gameGrid := GameGrid {
		Grid:   grid,
		Nodes:  nodes,
		Window: GameWindow{},
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

func (g *GameGrid) CreateNewPuzzle(diff int) {
	newPuzzle := GenerateNewPuzzle(diff);

	i := 0

	for x := 0; x < N_OF_LINES; x++ {
		for y := 0; y < N_OF_LINES; y++ {
			if newPuzzle[i] != '0' {
				g.Nodes[x][y].Label.SetText(string(newPuzzle[i]))
				g.Nodes[x][y].Value, _ = strconv.Atoi(string(newPuzzle[i]))
				g.Nodes[x][y].SetInactive()
			} else {
				g.Nodes[x][y].Label.SetText("")
				g.Nodes[x][y].Value = 0
				g.Nodes[x][y].SetActive(g)
			}

			i++
		}
	}
}

func (n *Node) SetActive(grid *GameGrid) {
	if n.Signal != 0 {
		n.EventBox.HandlerDisconnect(n.Signal)
		n.Signal = 0
	}

	n.Signal, _ = n.EventBox.Connect("button_press_event", func() { NewNumberSelectWindow(grid, n) })

	ctx, _ := n.Label.GetStyleContext()
	ctx.RemoveClass("gamegrid-node--inactive")

	n.isActive = true
}

func (n *Node) SetInactive() {
	if n.Signal != 0 {
		n.EventBox.HandlerDisconnect(n.Signal)
		n.Signal = 0
	}

	ctx, _ := n.Label.GetStyleContext()
	ctx.AddClass("gamegrid-node--inactive")

	n.isActive = false
}

func (n *Node) SetWrong() {
	ctx, _ := n.Label.GetStyleContext()
	ctx.AddClass("gamegrid-node--wrong")

	n.isWrong = true
}

func (n *Node) UnsetWrong() {
	ctx, _ := n.Label.GetStyleContext()
	ctx.RemoveClass("gamegrid-node--wrong")

	n.isWrong = false
}

func (g *GameGrid) NumberSelect(val string, node *Node) {
	// Close the window
	g.Window.window.Close()

	// Set the value
	node.SetNodeValue(val)

	// Check if node is wrong
	node.UnsetWrong()

	if val == "" {
		node.isWrong = false
		return // Empty value cannot be wrong
	}

	wrong := g.VerifyNode(node)
	node.isWrong = wrong

	if wrong {
		node.SetWrong()
	} else {
		// Verify if all are correct and filled
		wrong = g.VerifyAllNodes()

		// Game is won
		if !wrong {
			g.SetAllNodesAsInactive()
		}
	}
}

func (n *Node) SetNodeValue(val string) {
	n.Label.SetText(val)

	if val == "" {
		n.Value = 0
	} else {
		n.Value, _ = strconv.Atoi(val)
	}
}