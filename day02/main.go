package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeGame struct {
	Id    string
	Games []string
}

type CubeCount struct {
	Red   int
	Green int
	Blue  int
}

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

	var games []CubeGame

	for scanner.Scan() {
		slicedRecord := SliceString(scanner.Text())
		slicedGames := SplitGames(slicedRecord[1])
		currentGame := CubeGame{slicedRecord[0], slicedGames}

		games = append(games, currentGame)
	}

	totalSum := 0
	for _, curGame := range games {
		if CheckPossibility(curGame) {
			slcId := strings.Split(curGame.Id, " ")
			id, _ := strconv.ParseInt(slcId[1], 10, 32)
			totalSum += int(id)
		}
	}
	fmt.Println("Sum of possible game Id: ", totalSum)
}

func SliceString(inStr string) []string {
	return strings.Split(inStr, ":")
}

func SplitGames(inStr string) []string {
	return strings.Split(inStr, ";")
}

func CheckPossibility(inGame CubeGame) bool {
	var gameCounts []CubeCount

	for _, gameCubes := range inGame.Games {
		counts := GetCounts(gameCubes)
		gameCounts = append(gameCounts, counts)
	}

	for _, currentCount := range gameCounts {
		if currentCount.Red > 12 || currentCount.Blue > 14 || currentCount.Green > 13 {
			return false
		}
	}

	return true
}

func GetCounts(inCounts string) CubeCount {
	slcCubes := strings.Split(inCounts, ",")
	var cubeCount CubeCount

	for _, count := range slcCubes {
		count = strings.Trim(count, " ")
		slcCube := strings.Split(count, " ")

		switch slcCube[1] {
		case "red":
			numRed, _ := strconv.ParseInt(slcCube[0], 10, 32)
			cubeCount.Red = int(numRed)
		case "blue":
			numBlue, _ := strconv.ParseInt(slcCube[0], 10, 32)
			cubeCount.Blue = int(numBlue)
		case "green":
			numGreen, _ := strconv.ParseInt(slcCube[0], 10, 32)
			cubeCount.Green = int(numGreen)
		}
	}

	return cubeCount
}

// Part 2
func PartTwo() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []CubeGame

	//i := 0

	for scanner.Scan() {
		slicedRecord := SliceString(scanner.Text())
		slicedGames := SplitGames(slicedRecord[1])
		currentGame := CubeGame{slicedRecord[0], slicedGames}

		games = append(games, currentGame)

		//i++
		//if i >= 5 {
		//	break
		//}
	}

	totalSum := 0
	for _, curGame := range games {
		totalSum += GetGamePower(curGame)
	}
	fmt.Println("Total Power: ", totalSum)
}

func GetGamePower(inGame CubeGame) int {
	var gameCounts []CubeCount
	minCount := CubeCount{1, 1, 1}

	for _, gameCubes := range inGame.Games {
		counts := GetCounts(gameCubes)
		gameCounts = append(gameCounts, counts)
	}

	for _, curGame := range gameCounts {
		if curGame.Red > minCount.Red && curGame.Red > 0 {
			minCount.Red = curGame.Red
		}
		if curGame.Blue > minCount.Blue && curGame.Blue > 0 {
			minCount.Blue = curGame.Blue
		}
		if curGame.Green > minCount.Green && curGame.Green > 0 {
			minCount.Green = curGame.Green
		}
	}

	return minCount.Red * minCount.Blue * minCount.Green
}
