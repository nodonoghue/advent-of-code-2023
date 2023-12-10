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

func RoughBFS(startLoc Location, grid [][]string) int {
	var visited []VisitedLocaction
	step := 0
	//add the starting location to the visited array
	visited = append(visited, VisitedLocaction{X: startLoc.X, Y: startLoc.Y, Step: step})
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
	}
}

func PartOne(inputs [][]string) {
	fmt.Println("Starting part one")
	for _, line := range inputs {
		fmt.Println(line)
	}
	startLoc := FindStart(inputs)
	fmt.Println(startLoc)

	//Need to navigate the pipe loop until start is found again, then divide by 2 to fined the answer?
	//being mindful to take only steps that are legal with the current pipe symbol, particularly important
	//from start where every direction is legal, but not every direction will be allows by the pieces next
	//to it.

	//Or should I learn how to and implement a breadth first search algorithm
	//search outward from S like layers of an onion keep/visit matches and discard non-matches
	//store visits using a struct that will also note the step in which it was found
	//Need to know what to queue
}

func main() {
	fmt.Println("Advent of Code: Day 10")
	inputs := GetInputs("test.txt")
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
