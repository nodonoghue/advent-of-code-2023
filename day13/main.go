package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func GetInputs(fileName string) map[int][][]string {
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

func ParseLines(fileLines []string) map[int][][]string {
	returnObj := make(map[int][][]string)
	index := 0
	var grid [][]string
	for _, line := range fileLines {
		if line == "" {
			returnObj[index] = grid
			grid = nil
			index++
			continue
		}
		grid = append(grid, strings.Split(line, ""))
	}
	if grid != nil {
		returnObj[index] = grid
	}

	return returnObj
}

func FindHorizontalMirror(grid [][]string) {
	//use reflect.DeepEqual() to find the mirring, return index and count
	//Can use deepEqual for multiple rows in the grid, go outside and work in to find the point of reflection
	//What happens if the point of reflection is not in the middle?

	//Maybe split the grid on the index point, meaning start at 1 instead of 0
	//check if above and below match, make sure that both above and below are of equal length
	//may need to step through every row of the grid.
	endIndex := len(grid) - 1

	for i := 1; i <= endIndex/2; i++ {

	}
}

func FindVerticalMirror(grid [][]string) {
	//Rotate the grid 90 clockwise then find the horizontal mirroring, will equal the column of mirroring
	//How to rotate though?
}

func main() {
	fmt.Println("Advent of Code 2023: Day 13")

	s := "#.##..##."
	s2 := "##....##."

	s3 := "#.##..##."
	s4 := "##....##."

	var grid1 [][]string
	var grid2 [][]string

	grid1 = append(grid1, strings.Split(s, ""))
	grid1 = append(grid1, strings.Split(s2, ""))
	grid2 = append(grid2, strings.Split(s3, ""))
	grid2 = append(grid2, strings.Split(s4, ""))

	if reflect.DeepEqual(grid1, grid2) {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not equal")
	}

	inputs := GetInputs("test.txt")

	fmt.Println(inputs)
}
