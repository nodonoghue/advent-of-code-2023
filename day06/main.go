package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RaceSheet struct {
	Time     []int
	Distance []int
}

func main() {
	fmt.Println("Advent of Code 2023: Day 6")
	raceSheet := ParseFile("inputs.txt")
	PartOne(raceSheet)
	PartTwo(raceSheet)
}

func PartOne(raceSheet RaceSheet) {
	fmt.Println("Starting part 1")

	//iterate through the number of races to find the number of ways to win
	var totalWins []int
	for i := 0; i < len(raceSheet.Time); i++ {
		totalWins = append(totalWins, GetWins(raceSheet.Time[i], raceSheet.Distance[i]))
	}
	var answer int
	for _, val := range totalWins {
		if answer == 0 {
			answer = val
		} else {
			answer = answer * val
		}
	}

	fmt.Println("Part one answer: ", answer)
}

func PartTwo(raceSheet RaceSheet) {
	fmt.Println("Starting part 2")

	//put the dist and time slices together
	var time string
	var distance string
	for i := 0; i < len(raceSheet.Time); i++ {
		time = time + strconv.Itoa(raceSheet.Time[i])
		distance = distance + strconv.Itoa(raceSheet.Distance[i])
	}

	timeVal, _ := strconv.ParseInt(time, 10, 64)
	distVal, _ := strconv.ParseInt(distance, 10, 64)

	totalWins := GetWins(int(timeVal), int(distVal))

	fmt.Println("Part two answer: ", totalWins)
}

func ParseFile(fileName string) RaceSheet {
	var returnObj RaceSheet

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()

		if strings.Contains(strings.ToLower(currentLine), "time") {
			slcRemoveName := strings.Split(currentLine, ":")
			slcNumbers := strings.Split(slcRemoveName[1], " ")

			var times []int

			for _, currentNumber := range slcNumbers {
				val, err := strconv.ParseInt(currentNumber, 10, 64)
				if err == nil {
					times = append(times, int(val))
					returnObj.Time = times
				}
			}
		}

		if strings.Contains(strings.ToLower(currentLine), "distance") {
			slcRemoveName := strings.Split(currentLine, ":")
			slcNumbers := strings.Split(slcRemoveName[1], " ")

			var distances []int

			for _, currentNumber := range slcNumbers {
				val, err := strconv.ParseInt(currentNumber, 10, 64)
				if err == nil {
					distances = append(distances, int(val))
					returnObj.Distance = distances
				}
			}
		}
	}

	return returnObj
}

func GetWins(time int, distance int) int {
	wins := 0
	for i := 0; i <= time; i++ {
		dist := i * (time - i)
		if dist > distance {
			wins++
		}
	}

	return wins
}
