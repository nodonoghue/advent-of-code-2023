package main

import (
	"bufio"
	"fmt"
	"os"
)

type DesertMap struct {
	Directions []string
	Nodes      []Node
}

type Node struct {
	Label string
	Left  string
	Right string
}

func main() {
	fmt.Println("Advent of code 2023 day 8")
}

func GetInputs(fileName string) []DesertMap {
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

func ParseLines(fileLines []string) []DesertMap {
	var returnObj []DesertMap
	var hasDirections bool = false

	for _, currentLine := range fileLines {
		if len(currentLine) == 0 {
			continue
		}

		//Check for directions or nodes
		if hasDirections {
			// processing node lines
		} else {
			// process directions line
		}
	}

	return returnObj
}
