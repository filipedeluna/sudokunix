package gamelogic

import (
	"bufio"
	"github.com/gotk3/gotk3/gtk"
	"math/rand"
	"os"
	"strconv"
	"time"
	"utils/gtkutils"
)

const HARD_SEEDS_FILE string = "hard.txt"
const MED_SEEDS_FILE  string = "med.txt"
const EASY_SEEDS_FILE string = "easy.txt"

const N_OF_LINES int = 9

const ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LaunchNumberSelectWindow(gridLabel *gtk.Label) {
	win, styleProvider := gtkutils.NewWindow("Choose number")

	numberGrid, _ := gtk.GridNew()
	numberGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	numberGridCtx, _ := numberGrid.GetStyleContext()
	numberGridCtx.AddClass("numbergrid")
	numberGridCtx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	for x := 0; x < 9; x++ {
		evBox, _ := gtk.EventBoxNew()

		labelName := strconv.FormatInt(int64(x+1), 10)

		lab, _ := gtk.LabelNew(labelName)
		lab.SetJustify(gtk.JUSTIFY_CENTER)
		lab.SetHExpand(true)
		lab.SetVExpand(true)

		// Add CSS classes to node
		ctx, _ := lab.GetStyleContext()
		ctx.AddClass("numbergrid-node")
		ctx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

		// Create OnClickEvent
		evBox.Add(lab)
		evBox.Connect("button_press_event", func() { SetNodeValue(labelName, gridLabel, win) })

		numberGrid.Attach(evBox, x % 3, x / 3 + 1, 1, 1)
	}

	evBox, _ := gtk.EventBoxNew()

	lab, _ := gtk.LabelNew("Clear")
	lab.SetJustify(gtk.JUSTIFY_CENTER)
	lab.SetHExpand(true)
	lab.SetVExpand(true)

	// Add CSS classes to node
	ctx, _ := lab.GetStyleContext()
	ctx.AddClass("numbergrid-node")
	ctx.AddProvider(styleProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// Create OnClickEvent
	evBox.Add(lab)
	evBox.Connect("button_press_event", func() { SetNodeValue("", gridLabel, win) })

	numberGrid.Attach(evBox, 0, 4, 3, 1)


	win.Add(numberGrid);

	// Recursively show all widgets contained in this window.
	win.ShowAll()
}

func SetNodeValue(val string, lab *gtk.Label, win *gtk.Window) {
	lab.SetText(val)
	win.Close()
}



func GenerateNewPuzzle(diff int) string {
	// Get a random puzzle seed
	randomPuzzleSeed := getRandomPuzzleSeed(diff)

	// Rotate puzzle random n times x 90 degrees
	randomRotationTimes := rand.Intn(3 + 1)
	randomPuzzleSeedRotated := rotateSeed(randomPuzzleSeed, randomRotationTimes, N_OF_LINES)

	// Assign numbers to letters
	finishedSeed := assignRandomNumbersToSeed(randomPuzzleSeedRotated, N_OF_LINES)

	return finishedSeed
}

// Rotate seed matrix 90 degrees at a time
func rotateSeed(seed string, times int, nLines int)  string{
	seedLength := nLines * nLines
	seedChars := []rune(seed)

	for i := 0; i < times; i++ {
		var seedCharsRotated = make([]rune, len(seedChars))
		copy(seedCharsRotated, seedChars)

		for j := 0; j < seedLength; j++ {
			currentLine := (j / nLines) % seedLength + 1
			currentPositionInLine := j % nLines + 1

			seedCharsRotated[j] = seedChars[seedLength - (nLines * currentPositionInLine) - 1 + currentLine]
		}

		seedChars = seedCharsRotated
	}

	return string(seedChars)
}

func getRandomPuzzleSeed(diff int) string{
	// Get File
	file, _ := getFileByDifficulty(diff)

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// What puzzle seed to get from list
	randomPuzzleNumber := rand.Intn(100) + 1

	// Create a Scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Jump n times to get random seed
	for i := 0; i < randomPuzzleNumber; i++ {
		scanner.Scan()
	}

	return scanner.Text()
}

func getFileByDifficulty(diff int) (*os.File, error) {
	var fileName string

	// Get seed file name relative to difficulty
	if diff == 1 {
		fileName = EASY_SEEDS_FILE
	} else if diff == 2 {
		fileName = MED_SEEDS_FILE
	} else if diff == 3 {
		fileName = HARD_SEEDS_FILE
	}

	// Open file in path
	path := "seeds/" + fileName

	return os.Open(path)
}

func assignRandomNumbersToSeed(seed string, nLines int) string {
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	characters = characters[:nLines]

	// Append numbers
	numbers := []int{}
	for i := 1; i <= nLines; i++ {
		numbers = append(numbers, i)
	}

	// Shuffle numbers
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nLines; i++ {
		r := rand.Intn(nLines)
		a := numbers[i]
		b := numbers[r]

		numbers[i] = b
		numbers[r] = a
	}

	// Assign values to a map (Letter, Value)
	values := map[rune]int{}

	for i := 0; i < nLines; i++ {
		values[characters[i]] = numbers[i]
	}

	// Create new seed with values
	newSeed := []rune(seed)

	for i := 0; i < nLines * nLines; i++ {
		if newSeed[i] != '0' {
			newSeed[i] = rune(strconv.Itoa(values[newSeed[i]])[0])
		}
	}


	return string(newSeed)
}