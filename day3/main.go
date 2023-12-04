package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SymbolLoc struct {
	Row             int
	Column          int
	AdjacentNumbers int
}

type NumberLoc struct {
	Row            int
	StartingColumn int
	Length         int
	Value          int
}

func main() {
	fmt.Println("Advent of code 2023: Day 3 Part 1")
	PartOne()
	PartTwo()
}

func PartOne() {
	grid := MakeGrid()

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

func PartTwo() {
	grid := MakeGrid()

	foundGearSymbols := FindGearSymbols(grid)

	var totalRatio int = 0

	for _, curSymbol := range foundGearSymbols {
		totalRatio += GetGearRatio(grid, curSymbol)
	}
	fmt.Println("Part Two Answer: ", totalRatio)
}

func MakeGrid() [][]string {
	file, err := os.Open("inputs.txt")

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
	}

	return grid
}

func SliceString(inStr string) []string {
	return strings.Split(inStr, "")
}

func FindSymbols(inGrid [][]string) []SymbolLoc {
	var ret []SymbolLoc

	for rowNum, rowVal := range inGrid {
		for colNum, colVal := range rowVal {

			if colVal != "." && !IsNumeric(colVal) {
				var foundSymbol = SymbolLoc{rowNum, colNum, 0}
				ret = append(ret, foundSymbol)
			}
		}
	}

	return ret
}

func FindGearSymbols(inGrid [][]string) []SymbolLoc {
	var ret []SymbolLoc

	for rowNum, rowVal := range inGrid {
		for colNum, colVal := range rowVal {

			if colVal == "*" && !IsNumeric(colVal) {
				var foundSymbol = SymbolLoc{rowNum, colNum, 0}
				ret = append(ret, foundSymbol)
			}
		}
	}

	return ret
}

func IsNumeric(inStr string) bool {
	_, err := strconv.ParseInt(inStr, 10, 32)

	if err == nil {
		return true
	} else {
		return false
	}
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
		var count int = 0
		curRow := inGrid[curSymbol.Row]
		var foundNum NumberLoc

		//Up
		if (curSymbol.Row-1) > -1 && (curSymbol.Column-1) > -1 {
			foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column-1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
		if curSymbol.Row-1 > -1 {
			foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
		if curSymbol.Row-1 > -1 && curSymbol.Column+1 < len(curRow) {
			foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column+1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}

		// Same Row
		if curSymbol.Column-1 > -1 {
			foundNum = GetNumber(curSymbol.Row, curSymbol.Column-1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
		if curSymbol.Column+1 < len(curRow) {
			foundNum = GetNumber(curSymbol.Row, curSymbol.Column+1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}

		//Down
		if curSymbol.Row+1 < len(inGrid) && curSymbol.Column-1 > 0 {
			foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column-1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
		if curSymbol.Row+1 < len(inGrid) {
			foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
		if curSymbol.Row+1 < len(inGrid) && curSymbol.Column+1 < len(curRow) {
			foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column+1, inGrid)
			if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
				foundNumbers = append(foundNumbers, foundNum)
				count++
			}
		}
	}

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
	var num NumberLoc = NumberLoc{inRow, -1, 0, -1}

	curRow := grid[inRow]

	//go left first to find the beginning index
	for i := inCol; i >= 0; i-- {
		_, err := strconv.ParseInt(curRow[i], 10, 32)

		if err == nil {
			num.StartingColumn = i
		} else {
			break
		}
	}

	// walk right to find the end of the number
	endIndex := -1
	for i := inCol; i < len(curRow); i++ {
		_, err := strconv.ParseInt(curRow[i], 10, 32)

		if err == nil {
			endIndex = i
		} else {
			break
		}
	}

	if num.StartingColumn > -1 {
		//Get the current Number
		num.Length = endIndex - num.StartingColumn

		var curNum string
		for i := num.StartingColumn; i < len(curRow); i++ {
			_, err := strconv.ParseInt(curRow[i], 10, 32)

			if err == nil {
				curNum = curNum + curRow[i]
			} else {
				break
			}
		}

		val, err := strconv.ParseInt(curNum, 10, 32)

		if err == nil {
			num.Value = int(val)
		}
	}

	if num.StartingColumn == -1 {
		num.Value = 0
	}

	return num
}

func GetGearRatio(inGrid [][]string, curSymbol SymbolLoc) int {
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

	var count int = 0
	curRow := inGrid[curSymbol.Row]
	var foundNum NumberLoc

	//Up
	if (curSymbol.Row-1) > -1 && (curSymbol.Column-1) > -1 {
		foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column-1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}
	if curSymbol.Row-1 > -1 {
		foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}
	if curSymbol.Row-1 > -1 && curSymbol.Column+1 < len(curRow) {
		foundNum = GetNumber(curSymbol.Row-1, curSymbol.Column+1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}

	// Same Row
	if curSymbol.Column-1 > -1 {
		foundNum = GetNumber(curSymbol.Row, curSymbol.Column-1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}
	if curSymbol.Column+1 < len(curRow) {
		foundNum = GetNumber(curSymbol.Row, curSymbol.Column+1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}

	//Down
	if curSymbol.Row+1 < len(inGrid) && curSymbol.Column-1 > 0 {
		foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column-1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}
	if curSymbol.Row+1 < len(inGrid) {
		foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}
	if curSymbol.Row+1 < len(inGrid) && curSymbol.Column+1 < len(curRow) {
		foundNum = GetNumber(curSymbol.Row+1, curSymbol.Column+1, inGrid)
		if !IsDuplicateNumber(foundNum, foundNumbers) && foundNum.Value > 0 {
			foundNumbers = append(foundNumbers, foundNum)
			count++
		}
	}

	if count == 2 {
		return foundNumbers[0].Value * foundNumbers[1].Value
	}

	return 0
}
