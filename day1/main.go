package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	PartOne()
	PartTwo()
}

func PartOne() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalSum = 0

	for scanner.Scan() {
		currentSlice := SplitStrings(scanner.Text())
		currentInts := GetCalibrationInts(currentSlice)
		totalSum += SumValue(currentInts)
	}

	fmt.Println("Part1 - The total sum is: ", totalSum)
}

func PartTwo() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	//spelled out numbers between one and nine are still considered numbers
	scanner := bufio.NewScanner(file)

	var totalSum = 0

	for scanner.Scan() {
		currentInts := ReplaceNumbersPositional(scanner.Text())
		totalSum += SumValue(currentInts)
	}

	fmt.Println("Part2 - The total sum is: ", totalSum)
}

// Splits the incoming string into a string slice
func SplitStrings(inStr string) []string {
	return strings.Split(inStr, "")
}

// Takes in a string slice and finds the outer most integer values, returning the two digit pair
func GetCalibrationInts(inStr []string) string {
	var strFirstNum string
	var strSecondNum string

	// loop forwards for the first number
	for i := 0; i < len(inStr); i++ {
		_, err := strconv.ParseInt(inStr[i], 10, 32)

		if err == nil {
			strFirstNum = inStr[i]
			break
		}
	}

	// loop backwards for the second number
	for i := len(inStr) - 1; i >= 0; i-- {
		_, err := strconv.ParseInt(inStr[i], 10, 32)

		if err == nil {
			strSecondNum = inStr[i]
			break
		}
	}

	return strFirstNum + strSecondNum
}

// Get the sum of each slice of number pairs
func SumValue(inStr string) int {
	result, err := strconv.ParseInt(inStr, 10, 64)

	if err != nil {
		return 0
	} else {
		return int(result)
	}
}

// This doesn't work, as it doesn't take into account the positioning in the input line
// leaving in place for posterity
func ReplaceNumbers(inStr string, isReverse bool) string {
	inStr = strings.Replace(inStr, "one", "1", -1)
	inStr = strings.Replace(inStr, "two", "2", -1)
	inStr = strings.Replace(inStr, "three", "3", -1)
	inStr = strings.Replace(inStr, "four", "4", -1)
	inStr = strings.Replace(inStr, "five", "5", -1)
	inStr = strings.Replace(inStr, "six", "6", -1)
	inStr = strings.Replace(inStr, "seven", "7", -1)
	inStr = strings.Replace(inStr, "eight", "8", -1)
	inStr = strings.Replace(inStr, "nine", "9", -1)

	//dig out any numbers that were placed, the correct order
	inStrSlice := SplitStrings(inStr)
	if isReverse {
		// loop forwards for the first number
		for i := 0; i < len(inStrSlice); i++ {
			_, err := strconv.ParseInt(inStrSlice[i], 10, 32)

			if err == nil {
				return inStrSlice[i]
			}
		}
	} else {
		// loop backwards for the second number
		for i := len(inStrSlice) - 1; i >= 0; i-- {
			_, err := strconv.ParseInt(inStrSlice[i], 10, 32)

			if err == nil {
				return inStrSlice[i]
			}
		}
	}

	return inStr
}

func ReplaceNumbersPositional(inStr string) string {
	//walk left to right, every iteration run through the replace func
	//then look for digits
	var lString string = ""
	var lStringNum string
	var rString string = ""
	var rStringNum string

	// Create a string slice
	strSlice := SplitStrings(inStr)

	// This doesn't account for word burried inside the string...

	//Left to right
	for i := 0; i < len(strSlice); i++ {
		lString = lString + strSlice[i]
		_, err := strconv.ParseInt(lString, 10, 32)

		if err == nil {
			lStringNum = lString
			break
		}

		lString = ReplaceNumbers(lString, false)
		_, err2 := strconv.ParseInt(lString, 10, 32)
		if err2 == nil {
			lStringNum = lString
			break
		}
	}

	// loop backwards for the second number
	for i := len(strSlice) - 1; i >= 0; i-- {
		_, err := strconv.ParseInt(strSlice[i], 10, 32)
		if err == nil {
			rStringNum = strSlice[i]
			break
		}

		rString = ReplaceNumbers(inStr[i:], true)

		_, err2 := strconv.ParseInt(rString, 10, 32)

		if err2 == nil {
			rStringNum = rString
			break
		}
	}

	return lStringNum + rStringNum
}
