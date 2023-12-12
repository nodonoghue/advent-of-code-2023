package main

import (
	"bufio"
	"fmt"
	"math"
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

func FindDistances(grid [][]string, galaxies []Point) int {
	//Don't actually need to find a pair here, just need to
	//loop staring from the first galaxy, checking the diff
	//of both the x and y of all others, then to the next
	//iteration where the first item is ignore, and so on...
	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			//calc distance and get abs and sum
			firstPoint := galaxies[i]
			secondPoint := galaxies[j]

			distance := int(math.Abs(float64(firstPoint.x)-float64(secondPoint.x))) + int(math.Abs(float64(firstPoint.y)-float64(secondPoint.y)))
			totalDistance += distance
		}
	}
	return totalDistance
}

func FindDistancesPartTwo(colMap map[int]bool, rowMap map[int]bool, galaxies []Point, expansion int) int {
	totalDistance := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			//calc distance and get abs and sum
			firstPoint := galaxies[i]
			secondPoint := galaxies[j]

			if firstPoint.y > secondPoint.y {
				//iterate backwards from secondPoint.y to firstPoint.y indexes, if these cross an empty row, add 1mil for each
				for rowi := firstPoint.y; rowi >= secondPoint.y; rowi-- {
					if rowMap[rowi] {
						totalDistance += expansion
					}
				}
				totalDistance += firstPoint.y - secondPoint.y
			} else {
				for rowi := secondPoint.y; rowi >= firstPoint.y; rowi-- {
					if rowMap[rowi] {
						totalDistance += expansion
					}
				}
				totalDistance += secondPoint.y - firstPoint.y
			}

			if firstPoint.x > secondPoint.x {
				//iterate backwards from secondPoint.x to firstPoint.x indexes, for each emprty col this crosses add 1mil
				for coli := firstPoint.x; coli >= secondPoint.x; coli-- {
					if colMap[coli] {
						totalDistance += expansion
					}
				}
				totalDistance += firstPoint.x - secondPoint.x
			} else {
				//iterate backwards from firstPoint.x to secondPoint.x indexes, for each emprty col this crosses add 1mil
				for coli := secondPoint.x; coli >= firstPoint.x; coli-- {
					if colMap[coli] {
						totalDistance += expansion
					}
				}
				totalDistance += secondPoint.x - firstPoint.x
			}

			//totalDistance += int(math.Abs(float64(firstPoint.x)-float64(secondPoint.x))) + int(math.Abs(float64(firstPoint.y)-float64(secondPoint.y)))
		}
	}

	return totalDistance
}

func GetExpansionMaps(grid [][]string) (map[int]bool, map[int]bool) {
	colMap := FindEmptyColumns(grid)
	rowMap := FindEmptyRows(grid)
	return rowMap, colMap
}

func main() {
	fmt.Println("Advent of Code 2023: Day 11")
	grid := GetInputs("test.txt")
	PartOne(grid)
	PartTwo(grid)
}

func PartOne(grid [][]string) {
	fmt.Println("Starting part one")
	expandedGrid := ExpandGrid(grid, 1)
	galaxyLocations := FindGalaxies(expandedGrid)
	answer := FindDistances(expandedGrid, galaxyLocations)
	fmt.Println(answer)
}

func PartTwo(grid [][]string) {
	fmt.Println("Starting part two")

	rowMap, colMap := GetExpansionMaps(grid)
	galaxies := FindGalaxies(grid)
	answer := FindDistancesPartTwo(colMap, rowMap, galaxies, 100)

	fmt.Println("Part two answer: ", answer)
}
