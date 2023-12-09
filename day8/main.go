package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DesertMap struct {
	Directions []string
	Nodes      map[string]Node
}

type Node struct {
	Left  string
	Right string
}

func GetInputs(fileName string) DesertMap {
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

func ParseLines(fileLines []string) DesertMap {
	var returnObj DesertMap
	nodes := make(map[string]Node)
	var hasDirections bool = false

	for _, currentLine := range fileLines {
		// skip any blank lines
		if len(currentLine) == 0 {
			continue
		}

		//Check for directions or nodes
		if hasDirections {
			// processing node lines
			var currentNodeKey string
			var currentNode Node
			slcRawNode := strings.Split(currentLine, "=")
			currentNodeKey = strings.Trim(slcRawNode[0], " ")
			slcNodeLinks := strings.Split(slcRawNode[1], ",")

			for index := range slcNodeLinks {
				slcNodeLinks[index] = strings.Trim(slcNodeLinks[index], " ")
				if index == 0 {
					slcNodeLinks[index] = strings.Trim(slcNodeLinks[index], "(")
				} else {
					slcNodeLinks[index] = strings.Trim(slcNodeLinks[index], ")")
				}
			}
			currentNode.Left = slcNodeLinks[0]
			currentNode.Right = slcNodeLinks[1]
			nodes[currentNodeKey] = currentNode
		} else {
			// process directions line, first line is always the directions (assumption based on input structures)
			slcDirections := strings.Split(currentLine, "")
			returnObj.Directions = slcDirections
			hasDirections = true
		}
	}

	returnObj.Nodes = nodes
	return returnObj
}

func MakeStep(direction string, currentNode string, nodes map[string]Node) string {
	if direction == "L" {
		return nodes[currentNode].Left
	} else {
		return nodes[currentNode].Right
	}
}

// Get the total steps needed for all inputs to reach the end
func LeastCommonMultiple(a, b int, integers []int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i], integers[i+1:])
	}

	return result
}

// Get these for each set of factors to calculate the
// least common multiple
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	fmt.Println("Advent of code 2023 day 8")
	desertMap := GetInputs("inputs.txt")

	//PartOne(desertMap)
	PartTwo(desertMap)
}

func PartOne(desertMap DesertMap) {
	var steps int
	var currentLocation string = "AAA"

	for currentLocation != "ZZZ" {
		for _, currentDirection := range desertMap.Directions {
			steps++

			currentLocation = MakeStep(currentDirection, currentLocation, desertMap.Nodes)

			if currentLocation == "ZZZ" {
				break
			}
		}
	}

	fmt.Println("Part one answer: ", steps)
}

func PartTwo(desertMap DesertMap) {
	var ghostSteps []int
	var startLocations []string

	//Get all starting locations
	for key := range desertMap.Nodes {
		slcKey := strings.Split(key, "")
		if slcKey[2] == "A" {
			startLocations = append(startLocations, key)
		}
	}

	//Iterate through each starting location the same as a single starting location to get the number of steps
	//each will need individually, then calculate the LCM for the results.
	done := false
	for !done {
		for _, startLocation := range startLocations {
			currentLocation := startLocation
			steps := 0
			for strings.Split(currentLocation, "")[2] != "Z" {
				for _, currentDirection := range desertMap.Directions {
					steps++

					currentLocation = MakeStep(currentDirection, currentLocation, desertMap.Nodes)

					if strings.Split(currentLocation, "")[2] == "Z" {
						break
					}
				}
			}
			ghostSteps = append(ghostSteps, steps)
		}
		done = true
	}

	fmt.Println(ghostSteps)
	answer := LeastCommonMultiple(ghostSteps[0], ghostSteps[1], ghostSteps[2:])
	fmt.Println("Part two answer: ", answer)
}
