package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//The calibration values are the outer mose integers, will be pairs
	//if only one integer, repeat to make the pair
	firstValue := "1abc2"
	secondValue := "pqr3stu8vwx"
	thirdValue := "a1b2c3d4e5f"
	fourthValue := "treb7uchet"

	// split the strings into slices
	firstValSlice := SplitStrings(firstValue)
	secondValSlice := SplitStrings(secondValue)
	thirdValSlice := SplitStrings(thirdValue)
	fourthValSlice := SplitStrings(fourthValue)

	// find the integers
	firstInts := GetCalibrationInts(firstValSlice)
	secondInts := GetCalibrationInts(secondValSlice)
	thirdInts := GetCalibrationInts(thirdValSlice)
	fourthInts := GetCalibrationInts(fourthValSlice)

	totalSum := SumValue(firstInts) + SumValue(secondInts) + SumValue(thirdInts) + SumValue(fourthInts)

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
