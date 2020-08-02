package game

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"deluna.pt/luna/sudokunix/internal/utils"
	"strconv"
)

type GameGrid struct {
	Grid   *gtk.Grid
	Nodes  [N_OF_LINES][N_OF_LINES]Node
	Window GameWindow // These are popup windows. There should only be one active at any given time (eg. pick number)
	CandidateMode bool // Alternates between write and candidate mode
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
	candidates [N_OF_LINES]bool
	candidateMode bool
}

func (g *GameGrid) SetCandidateMode() {
	g.CandidateMode = !g.CandidateMode
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

			nodes[x][y] = Node{ lab, 0,x, y,false, false, evBox, 0, [N_OF_LINES]bool{}, false }
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
	newPuzzle := GenerateNewPuzzle(diff)

	i := 0

	for x := 0; x < N_OF_LINES; x++ {
		for y := 0; y < N_OF_LINES; y++ {
			g.Nodes[x][y].DisableCandidateMode()

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
	g.Window.window.Close()

	node.UnsetWrong()

	// Check if in candidate mode
	if g.CandidateMode {
		// Set the value
		node.SetNodeValue("")

		node.EnableCandidateMode()

		if val == "" { // Candidates need to be cleared
			node.ResetCandidates()
		} else {
			fixedVal, _ := strconv.Atoi(val)

			node.ToggleCandidate(fixedVal)

			node.SetCandidatesLabel()
		}


	} else {
		node.DisableCandidateMode()

		node.SetNodeValue(val)

		// Empty value cannot be wrong
		if val == "" {
			return
		}

		// Verify if node is wrong
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
}

func (n *Node) SetNodeValue(val string) {
	n.Label.SetText(val)

	if val == "" {
		n.Value = 0
	} else {
		n.Value, _ = strconv.Atoi(val)
	}
}

func (n *Node) EnableCandidateMode() {
	if !n.candidateMode {
		// Restart everything
		n.candidateMode = true
		n.candidates = [N_OF_LINES]bool{}
	}

	ctx, _ := n.Label.GetStyleContext()
	ctx.AddClass("gamegrid-node--candidatemode")
}

func (n *Node) DisableCandidateMode() {
	// Restart everything
	n.candidateMode = false
	n.candidates = [N_OF_LINES]bool{}

	ctx, _ := n.Label.GetStyleContext()
	ctx.RemoveClass("gamegrid-node--candidatemode")
}

func (n *Node) ToggleCandidate(val int) {
	n.candidates[val - 1] = !n.candidates[val - 1]
}

func (n *Node) ResetCandidates() {
	n.candidates = [N_OF_LINES]bool{}
}

func (n *Node) SetCandidatesLabel() {
	output := ""

	k := 0
	for i := 0; i < N_OF_LINES; i ++ {
		if n.candidates[i] {
			if k == 3 {
				k = 0
				output += "\n"
			}

			output += " " + strconv.Itoa(i + 1)
			k++
		}
	}

	n.Label.SetText(output)
}
