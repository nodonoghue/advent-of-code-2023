package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputs(fileName string) [][]int {
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

func ParseLines(fileLines []string) [][]int {
	var returnObj [][]int

	for _, line := range fileLines {
		slcLine := strings.Split(line, " ")
		var convertedNums []int

		for _, num := range slcLine {
			cNum, err := strconv.ParseInt(num, 10, 64)
			if err == nil {
				convertedNums = append(convertedNums, int(cNum))
			}
		}
		returnObj = append(returnObj, convertedNums)
	}

	return returnObj
}

// func to recursively get the differences into an array of array of ints
// will recurse until all diffs are 0 and return this [][]int
func GetDiffs(sequence []int) [][]int {
	var returnObj [][]int
	returnObj = append(returnObj, sequence)
	done := false
	firstRun := true
	var latestSequence []int

	for !done {
		var diffs []int
		if firstRun {
			diffs = CalculateDiffs(sequence)
			latestSequence = diffs
			firstRun = false
		} else {
			diffs = CalculateDiffs(latestSequence)
			latestSequence = diffs
		}
		returnObj = append(returnObj, diffs)

		allZeros := true
		for _, diff := range diffs {
			if diff != 0 {
				allZeros = false
			}
		}
		done = allZeros
	}

	return returnObj
}

func CalculateDiffs(sequence []int) []int {
	var returnObj []int

	for i := 0; i < len(sequence)-1; i++ {
		returnObj = append(returnObj, sequence[i+1]-sequence[i])
	}

	return returnObj
}

// func to add to the sequence by adding an additional item to the final element in
// the passed in [][]int then work backwards adding an additional item to each one until
// a new item has been added, returns the new item as int
func GetNextNum(diffs [][]int) int {
	returnObj := 0

	addNum := 0
	for i := len(diffs) - 1; i > 0; i-- {
		addNum = diffs[i][len(diffs[i])-1] + addNum
	}

	returnObj = addNum + diffs[0][len(diffs[0])-1]

	return returnObj
}

// func to add to the sequene by subtracting the first diff from the first number
func GetPreviousNumber(diffs [][]int) int {
	returnObj := 0

	addNum := 0
	for i := len(diffs) - 1; i > 0; i-- {
		addNum = diffs[i][0] - addNum
	}
	returnObj = diffs[0][0] - addNum

	return returnObj
}

func PartOne(inputs [][]int) {
	answer := 0

	for _, sequence := range inputs {
		slcDiffs := GetDiffs(sequence)
		nextNum := GetNextNum(slcDiffs)
		answer += nextNum
	}

	fmt.Println("Part one answer: ", answer)
}

func PartTwo(inputs [][]int) {
	answer := 0

	for _, sequence := range inputs {
		slcDiffs := GetDiffs(sequence)
		prevNum := GetPreviousNumber(slcDiffs)
		answer += prevNum
	}

	fmt.Println("Pard two answer: ", answer)
}

func main() {
	fmt.Println("Advent of Code 2023: Day 9")
	inputs := GetInputs("inputs.txt")

	PartOne(inputs)
	PartTwo(inputs)
}
