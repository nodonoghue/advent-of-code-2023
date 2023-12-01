package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	fmt.Println("The total sum is: ", totalSum)
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

// Add a func to read the file line by line, then take each line through each of the
// other funcs to get to the final integer and add to a runnint total sum
