package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SymbolLoc struct {
	Row    int
	Column int
}

type NumberLoc struct {
	Row            int
	StartingColumn int
	Length         int
	Value          int
}

func main() {
	PartOne()
}

func PartOne() {
	fmt.Println("Advent of code 2023: Day 3 Part 1")

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		curRow := SliceString(scanner.Text())
		grid = append(grid, curRow)
		fmt.Println(curRow)
	}

	// Need to traverse the gride to find all symbols and record their  coords
	foundSymbols := FindSymbols(grid)

	//Need to look around each symbol to fund the numbers
	foundNumbers := FindNumbers(grid, foundSymbols)

	var totalSum int
	for _, curNumber := range foundNumbers {
		totalSum += curNumber.Value
	}

	fmt.Println("Part One Answer: ", totalSum)
}

func SliceString(inStr string) []string {
	return strings.Split(inStr, "")
}

func FindSymbols(inGrid [][]string) []SymbolLoc {
	var ret []SymbolLoc

	for rowNum, rowVal := range inGrid {
		for colNum, colVal := range rowVal {

			if colVal == "*" || colVal == "#" || colVal == "$" || colVal == "+" {
				var foundSymbol = SymbolLoc{rowNum, colNum}
				ret = append(ret, foundSymbol)
			}
		}
	}

	fmt.Println(ret)

	return ret
}

func FindNumbers(inGrid [][]string, symbols []SymbolLoc) []NumberLoc {
	var foundNumbers []NumberLoc

	//Scan around the found symbols to find a numnber char
	// Ensure that the row and col to scan are not invalid
	//Up + Left = [row - 1][col -1]
	//up = [row - 1][col]
	//up + right = [row - 1][col + 1]
	//left = [row][col - 1]
	//right = [row][col + 1]
	//down + left = [row + 1][col - 1]
	//down = [row + 1][col]
	//down + right = [row + 1][col + 1]
	for _, curSymbol := range symbols {
		curRow := inGrid[curSymbol.Row]

		fmt.Println("Current Symbol: ", inGrid[curSymbol.Row][curSymbol.Column], curSymbol)

		//Up
		if (curSymbol.Row-1) > -1 && (curSymbol.Column-1) > -1 {
			GetNumber(curSymbol.Row-1, curSymbol.Column-1, inGrid)
		}
		if curSymbol.Row-1 > -1 {
			GetNumber(curSymbol.Row-1, curSymbol.Column, inGrid)
		}
		if curSymbol.Row-1 > -1 && curSymbol.Column+1 < len(curRow) {
			GetNumber(curSymbol.Row-1, curSymbol.Column+1, inGrid)
		}

		// Same Row
		if curSymbol.Column-1 > -1 {
			GetNumber(curSymbol.Row, curSymbol.Column-1, inGrid)
		}
		if curSymbol.Column+1 < len(curRow) {
			GetNumber(curSymbol.Row, curSymbol.Column+1, inGrid)
		}

		//Down
		if curSymbol.Row+1 < len(inGrid) && curSymbol.Column-1 > 0 {
			GetNumber(curSymbol.Row+1, curSymbol.Column-1, inGrid)
		}
		if curSymbol.Row+1 < len(inGrid) {
			GetNumber(curSymbol.Row+1, curSymbol.Column, inGrid)
		}
		if curSymbol.Row+1 < len(inGrid) && curSymbol.Column+1 < len(curRow) {
			GetNumber(curSymbol.Row+1, curSymbol.Column+1, inGrid)
		}
	}

	//If a number is found move left, then right to find the whole number
	//Create a new NumberLoc type containing the current row, startingCol and number length

	return foundNumbers
}

func IsDuplicateNumber(inNum NumberLoc, foundNumbers []NumberLoc) bool {
	for _, curNumber := range foundNumbers {
		if inNum.Row == curNumber.Row && inNum.StartingColumn == curNumber.StartingColumn {
			return true
		}
	}

	return false
}

func GetNumber(inRow int, inCol int, grid [][]string) NumberLoc {
	var num NumberLoc = NumberLoc{-1, -1, 0, -1}

	curRow := grid[inRow]

	//go left first to find the beginning index
	for i := inCol; i >= 0; i-- {
		_, err := strconv.ParseInt(curRow[i], 10, 32)

		if err == nil {
			num.StartingColumn = i
			//fmt.Println("val: ", val)
		}
	}
	if num.StartingColumn > -1 {
		//fmt.Println("Starting Index: row, col: ", inRow, num.StartingColumn)
	}

	return num
}
