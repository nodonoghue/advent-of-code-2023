package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

var PipeChunk = func() map[string]Directions {
	return map[string]Directions{
		"|": {X: "1 -1"},
		"-": {Y: "1 -1"},
		"L": {X: "-1", Y: "1"},
		"J": {X: "-1", Y: "-1"},
		"7": {X: "-1", Y: "1"},
		"F": {X: "1", Y: "1"},
		"S": {X: "1 -1", Y: "1 -1"},
		".": {X: "0", Y: "0"},
	}
}

type Directions struct {
	X string
	Y string
}

type Location struct {
	X int
	Y int
}

type VisitedLocaction struct {
	X    int
	Y    int
	Step int
}

func GetInputs(fileName string) [][]string {
	file := OpenFile(fileName)
	fileLines := ReadFile(file)
	return ParseLines(fileLines)
}

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}
	return file
}

func ReadFile(file *os.File) []string {
	var returnObj []string

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		returnObj = append(returnObj, scanner.Text())
	}

	return returnObj
}

func ParseLines(fileLines []string) [][]string {
	var returnObj [][]string

	for _, line := range fileLines {
		slcLine := strings.Split(line, "")
		returnObj = append(returnObj, slcLine)
	}

	return returnObj
}

func FindStart(inputs [][]string) Location {
	var returnObj Location
	for rowIndex, row := range inputs {
		for colIndex, col := range row {
			if col == "S" {
				returnObj.X = colIndex
				returnObj.Y = rowIndex
			}
		}
	}
	return returnObj
}

func TracePath(startLoc Location, grid [][]string) int {
	returnVal := 0
	nextChar := ""
	currentLoc := startLoc
	currentChar := "S"
	var previousLoc Location

	for nextChar != "S" {
		if currentChar == "S" {
			//check check around for pipes that would connect
			if currentLoc.Y > 0 {
				currentChar = grid[currentLoc.Y-1][currentLoc.X]
				if currentChar == "|" || currentChar == "7" || currentChar == "F" {
					currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y - 1}
					previousLoc = startLoc
					nextChar = currentChar
					returnVal++
					continue
				}
			}
			if currentLoc.X < len(grid[0])-1 {
				currentChar = grid[currentLoc.Y][currentLoc.X+1]
				if currentChar == "-" || currentChar == "7" || currentChar == "J" {
					currentLoc = Location{X: currentLoc.X + 1, Y: currentLoc.Y}
					nextChar = currentChar
					previousLoc = startLoc
					returnVal++
					continue
				}
			}
			if currentLoc.Y < len(grid)-1 {
				currentChar = grid[currentLoc.Y+1][currentLoc.X]
				if currentChar == "|" || currentChar == "J" || currentChar == "L" {
					currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y + 1}
					previousLoc = startLoc
					nextChar = currentChar
					returnVal++
					continue
				}
			}
			if currentLoc.X > 0 {
				currentChar = grid[currentLoc.Y][currentLoc.X-1]
				if currentChar == "-" || currentChar == "L" || currentChar == "F" {
					currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y - 1}
					previousLoc = startLoc
					nextChar = currentChar
					returnVal++
					continue
				}
			}
		}
		switch currentChar {
		case "-": //can move left or right
			if currentLoc.X > previousLoc.X {
				//move right
				currentChar = grid[currentLoc.Y][currentLoc.X+1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X + 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.X < previousLoc.X {
				//move left
				currentChar = grid[currentLoc.Y][currentLoc.X-1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X - 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
		case "7": //can move left or down
			if currentLoc.X > previousLoc.X {
				//move down
				currentChar = grid[currentLoc.Y+1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y + 1}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.Y < previousLoc.Y {
				//move left
				currentChar = grid[currentLoc.Y][currentLoc.X-1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X - 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
		case "|": //can move up or down
			if currentLoc.Y < previousLoc.Y {
				//move up
				currentChar = grid[currentLoc.Y-1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y - 1}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.Y > previousLoc.Y {
				//move down
				currentChar = grid[currentLoc.Y+1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y + 1}
				nextChar = currentChar
				returnVal++
				continue
			}
		case "J": // canse move right or up
			if currentLoc.Y > previousLoc.Y {
				//move left
				currentChar = grid[currentLoc.Y][currentLoc.X-1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X - 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.X > previousLoc.X {
				//move up
				currentChar = grid[currentLoc.Y-1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y - 1}
				nextChar = currentChar
				returnVal++
				continue
			}
		case "L": //can move right or up
			if currentLoc.Y > previousLoc.Y {
				//move right
				currentChar = grid[currentLoc.Y][currentLoc.X+1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X + 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.X < previousLoc.X {
				//move up
				currentChar = grid[currentLoc.Y-1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y - 1}
				nextChar = currentChar
				returnVal++
				continue
			}
		case "F":
			if currentLoc.X < previousLoc.X {
				//move down
				currentChar = grid[currentLoc.Y+1][currentLoc.X]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X, Y: currentLoc.Y + 1}
				nextChar = currentChar
				returnVal++
				continue
			}
			if currentLoc.Y < previousLoc.Y {
				//move right
				currentChar = grid[currentLoc.Y][currentLoc.X+1]
				previousLoc = currentLoc
				currentLoc = Location{X: currentLoc.X + 1, Y: currentLoc.Y}
				nextChar = currentChar
				returnVal++
				continue
			}
		default:
			nextChar = "S"
		}

	}

	return returnVal / 2
}

func RoughBFS(startLoc Location, grid [][]string) int {
	var visited []VisitedLocaction
	steps := 0
	//add the starting location to the visited array
	visited = append(visited, VisitedLocaction{X: startLoc.X, Y: startLoc.Y, Step: steps})
	queue := list.New()

	//Add surrounding cells to "S" to the queue
	if startLoc.X > 0 {
		queue.PushBack(Location{X: startLoc.X - 1, Y: startLoc.Y})
	}
	if startLoc.X < len(grid[0]) {
		queue.PushBack(Location{X: startLoc.X + 1, Y: startLoc.Y})
	}
	if startLoc.Y > 0 {
		queue.PushBack(Location{X: startLoc.X, Y: startLoc.Y - 1})
	}
	if startLoc.Y < len(grid) {
		queue.PushBack(Location{X: startLoc.X, Y: startLoc.Y + 1})
	}

	for queue.Len() < 0 {
		//pop off queue and search for a valid connector to the current location
		thing := queue.Front()

		fmt.Println(thing.Value.(*Location))
	}

	return steps
}

func PartOne(inputs [][]string) {
	fmt.Println("Starting part one")
	startLoc := FindStart(inputs)

	answer := TracePath(startLoc, inputs)
	fmt.Println("Answer ", answer)

	//Need to navigate the pipe loop until start is found again, then divide by 2 to fined the answer?
	//being mindful to take only steps that are legal with the current pipe symbol, particularly important
	//from start where every direction is legal, but not every direction will be allows by the pieces next
	//to it.  --Solved it initially like this.  Will refactor to learn BFS algorithms

	//Or should I learn how to and implement a breadth first search algorithm
	//search outward from S like layers of an onion keep/visit matches and discard non-matches
	//store visits using a struct that will also note the step in which it was found
	//Need to know what to queue
}

func PartTwo(inputs [][]string) {
	fmt.Println("Starting part two")

	//I really have no idea how to tackle this part.  Need to do some research into what algorithms would work.
	//or figure out a way to brute force the answer out.

	//Maybe storing a map of the nodes that make up the outline of the circuit and figuring out a way to scan
	//down the rows and have an understanding of when inside the lines....  but how?
}

func main() {
	fmt.Println("Advent of Code: Day 10")
	inputs := GetInputs("inputs.txt")
	PartOne(inputs)
}

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
