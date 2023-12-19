package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func ProcessLine(currentRow SpringRow) int {
	replaceMap, bools := SplitUnknowns(currentRow.row)
	regexpstr := BuildRegExpString(currentRow.pattern)
	permutations := GetPermutations(bools)
	count := 0

	for _, permutation := range permutations {
		//put the string back together from the bool array
		reconstructed := ReconstructString(currentRow.row, permutation, replaceMap)

		//get a count of validations
		if IsValid(reconstructed, regexpstr) {
			count++
		}
	}

	return count
}

func GetPermutations(bools []bool) [][]bool {
	var result [][]bool
	n := len(bools)

	var permutate func(int)
	permutate = func(index int) {
		if index == n {
			permutationCopy := make([]bool, n)
			copy(permutationCopy, bools)
			result = append(result, permutationCopy)
			return
		}

		bools[index] = true
		permutate(index + 1)

		bools[index] = false
		permutate(index + 1)
	}

	permutate(0)
	return result
}

func ReconstructString(linestr string, bools []bool, replaceMap map[int]bool) string {
	slice := strings.Split(linestr, "")
	boolIndex := 0

	for index := range slice {
		if _, isOk := replaceMap[index]; isOk {
			if bools[boolIndex] {
				slice[index] = "#"
			} else {
				slice[index] = "."
			}
			boolIndex++
		}
	}

	return strings.Join(slice, "")
}

func SplitUnknowns(row string) (map[int]bool, []bool) {
	unknownMap := make(map[int]bool)
	var boolSlc []bool

	slcRow := strings.Split(row, "")

	for index, value := range slcRow {
		if value == "?" {
			unknownMap[index] = false
			boolSlc = append(boolSlc, false)
		}
	}

	return unknownMap, boolSlc
}

func IsValid(line string, regexpstr string) bool {
	isValid := false
	regEx, _ := regexp.Compile(regexpstr)
	isValid = regEx.MatchString(line)
	return isValid
}

func BuildRegExpString(pattern []int) string {
	returnVal := `^\.*`
	patternLen := len(pattern)
	//`^#{1}\.+#{1}\.+#{3}$`gm
	for i := 0; i < patternLen; i++ {
		returnVal = returnVal + `#{` + strconv.FormatInt(int64(pattern[i]), 10) + `}`
		if i < (patternLen - 1) {
			returnVal = returnVal + `\.+`
		}
	}
	returnVal += `\.*$`

	return returnVal
}

func UnfoldLine(line string) string {
	newLine := line

	for i := 1; i < 5; i++ {
		newLine += "?" + line
	}

	return newLine
}

func UnfoldPattern(pattern []int) []int {
	var newPattern []int

	for i := 1; i <= 5; i++ {
		for j := 0; j < len(pattern); j++ {
			newPattern = append(newPattern, pattern[j])
		}
	}

	return newPattern
}

func PartOne(inputs []SpringRow) {
	fmt.Println("Starting part one")
	answer := 0

	for _, val := range inputs {
		fmt.Println("Starting Line: ", val)
		answer += ProcessLine(val)
	}

	fmt.Println("Part one answer: ", answer)
}

func PartTwo(inputs []SpringRow) {
	fmt.Println("Starting Part Two")

	answer := 0

	//Need to unfold the records, both line and pattern
	//send each line through the same logic as above
	//Will run very, very slowly....
	for _, val := range inputs {
		var newRow SpringRow
		newRow.pattern = UnfoldPattern(val.pattern)
		newRow.row = UnfoldLine(val.row)

		fmt.Println("Starting Line: ", newRow)
		answer += ProcessLine(newRow)
	}

}

func main() {
	fmt.Println("Advent of Code 2023: Day 12")
	inputs := GetInputs("test.txt")
	PartOne(inputs)
	//PartTwo(inputs)
}
