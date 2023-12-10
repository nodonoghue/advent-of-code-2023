package main

import (
	"bufio"
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

func PartOne(inputs [][]string) {
	fmt.Println("Starting part one")
	for _, line := range inputs {
		fmt.Println(line)
	}
	startLoc := FindStart(inputs)
	fmt.Println(startLoc)

	fmt.Println(inputs[startLoc.Y][startLoc.X])
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
