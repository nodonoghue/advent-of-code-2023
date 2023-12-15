package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SpringRow struct {
	row     string
	pattern []int
}

func GetInputs(fileName string) []SpringRow {
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

func ParseLines(fileLines []string) []SpringRow {
	var returnObj []SpringRow

	for _, line := range fileLines {
		var currentRow SpringRow
		var rowPattern []int
		slcLine := strings.Split(line, " ")
		currentRow.row = slcLine[0]
		slcPatternNums := strings.Split(slcLine[1], ",")
		for _, numStr := range slcPatternNums {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err == nil {
				rowPattern = append(rowPattern, int(num))
			}
		}
		currentRow.pattern = rowPattern
		returnObj = append(returnObj, currentRow)
	}

	return returnObj
}

func IsValid(line string) bool {
	isValid := false

	//How do I check against the pattern here?
	//Regex use?  How do I make a regEx to test 1, 1, 3 with any number of spaces between each group?

	return isValid
}

func PartOne(inputs []SpringRow) {
	fmt.Println("Starting part one")
	//General idea is to splut out the "?" from the rest of the string, map back to the index in the original string
	//make an array of bool the same length, permutate over the bool array.  For each permutation check if valid by
	//using the following regex:
}

func main() {
	fmt.Println("Advent of Code 2023: Day 12")
	inputs := GetInputs("test.txt")

	fmt.Println(inputs)
}
