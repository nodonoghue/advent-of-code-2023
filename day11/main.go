package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
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

func ExpandGrid(grid [][]string, expansionSize int) [][]string {
	var returnObj [][]string

	rowMap := FindEmptyRows(grid)
	colMap := FindEmptyColumns(grid)

	for rowIndex, row := range grid {
		var newRow []string
		for colIndex, col := range row {
			if colMap[colIndex] {
				if expansionSize > 0 {
					for i := 1; i <= expansionSize; i++ {
						newRow = append(newRow, ".")
					}
				}
			}
			newRow = append(newRow, col)
		}

		if rowMap[rowIndex] {
			if expansionSize > 0 {
				for i := 1; i <= expansionSize; i++ {
					returnObj = append(returnObj, newRow)
				}
			}
		}
		returnObj = append(returnObj, newRow)
	}

	return returnObj
}

func FindEmptyRows(grid [][]string) map[int]bool {
	rowMap := make(map[int]bool)

	for index, row := range grid {
		isEmpty := true
		for _, char := range row {
			if char == "#" {
				isEmpty = false
			}
		}
		if isEmpty {
			rowMap[index] = true
		} else {
			rowMap[index] = false
		}
	}

	return rowMap
}

func FindEmptyColumns(grid [][]string) map[int]bool {
	var colMap = make(map[int]bool)

	for i := 0; i < len(grid[0]); i++ {
		colMap[i] = true
	}

	for _, row := range grid {
		for colIndex, char := range row {
			if char == "#" {
				colMap[colIndex] = false
			}
		}
	}

	return colMap
}

func FindGalaxies(grid [][]string) []Point {
	var returnObj []Point

	for rowIndex, row := range grid {
		for colIndex, col := range row {
			if col == "#" {
				returnObj = append(returnObj, Point{x: colIndex, y: rowIndex})
			}
		}
	}

	return returnObj
}

func FindDistances(grid [][]string, galaxies []Point) {
	//Don't actually need to find a pair here, just need to
	//loop staring from the first galaxy, checking the diff
	//of both the x and y of all others, then to the next
	//iteration where the first item is ignore, and so on...
}

func main() {
	fmt.Println("Advent of Code 2023: Day 11")
	grid := GetInputs("test.txt")
	PartOne(grid)
}

func PartOne(grid [][]string) {
	fmt.Println("Starting part one")
	expandedGrid := ExpandGrid(grid, 1)
	galaxyLocations := FindGalaxies(expandedGrid)

	fmt.Println(galaxyLocations)

}

func PartTwo() {
	fmt.Println("Starting part two")
}
