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
}

func FindVerticalMirror(grid [][]string) {
	//Not entirely sure how to do this...
	//Could rotate the grid and call FindHorizontalMirror()?
	//Is there an algorithm to rotate the grid or do it the hard way and brute force it.
}

func main() {
	fmt.Println("Advent of Code 2023: Day 13")

	s := "#.##..##."
	s2 := "##....##."

	slc1 := strings.Split(s, "")
	slc2 := strings.Split(s2, "")

	if reflect.DeepEqual(slc1, slc2) {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not equal")
	}

	inputs := GetInputs("test.txt")

	fmt.Println(inputs)
}
