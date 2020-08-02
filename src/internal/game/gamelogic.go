package game

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const HARD_SEEDS_FILE string = "hard.txt"
const MED_SEEDS_FILE  string = "med.txt"
const EASY_SEEDS_FILE string = "easy.txt"

const SEEDS_FOLDER string ="assets/seeds/"

const N_OF_LINES int = 9

const ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


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
	path := SEEDS_FOLDER + fileName

	return os.Open(path)
}

func assignRandomNumbersToSeed(seed string, nLines int) string {
	characters := []rune(ALPHABET)
	characters = characters[:nLines]

	// Append numbers
	var numbers []int
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

// Returns true if node is wrong
func (g GameGrid) VerifyNode(node *Node) bool {
	// Verify node square
	nodeCol := node.X / 3
	nodeRow := node.Y / 3

	for x := nodeCol * 3; x < nodeCol * 3 + 3; x++ {
		for y := nodeRow * 3; y < nodeRow * 3 + 3; y++ {
			if node.X != x && node.Y != y {
				if node.Value == g.Nodes[x][y].Value {
					return true
				}
			}
		}
	}

	// Verify Horizontal
	for y := 0; y < N_OF_LINES; y++ {
		if node.Y != y {
			if node.Value == g.Nodes[node.X][y].Value {
				return true
			}
		}
	}

	// Verify Vertical
	for x := 0; x < N_OF_LINES; x++ {
		if node.X != x {
			if node.Value == g.Nodes[x][node.Y].Value {
				return true
			}
		}
	}

	return false
}

// Returns true if any node is wrong
func (g GameGrid) VerifyAllNodes() bool {
	for x := 0; x < N_OF_LINES; x++ {
		for y := 0; y < N_OF_LINES; y++ {
			if g.Nodes[x][y].isWrong || (g.Nodes[x][y].Value == 0 && g.Nodes[x][y].isActive) {
				return true
			}
		}
	}

	return false
}

// Set all nodes as inactive
func (g *GameGrid) SetAllNodesAsInactive() {
	for x := 0; x < N_OF_LINES; x++ {
		for y := 0; y < N_OF_LINES; y++ {
			g.Nodes[x][y].SetInactive()
		}
	}
}