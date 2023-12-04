package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	PickedNumbers  []int
	WinCount       int
	Copies         int
}

func main() {
	fmt.Println("Advent of Code 2023 Day 4")
	PartOne()
	PartTwo()
}

func PartOne() {
	cards := GetCards()
	var totalPoints int
	for _, curCard := range cards {
		totalPoints += int(CalculateScore(curCard.WinningNumbers, curCard.PickedNumbers))
	}
	fmt.Println("Part one answer: ", totalPoints)
}

func PartTwo() {
	var totalCards int
	cards := GetCards()

	for index, curCard := range cards {
		cards[index].WinCount = CountWins(curCard.WinningNumbers, curCard.PickedNumbers)
		cards[index].Copies = 1
	}

	for i := 0; i < len(cards); i++ {
		if cards[i].WinCount > 0 {
			for j := i + 1; j <= i+cards[i].WinCount; j++ {
				cards[j].Copies += 1 * cards[i].Copies
				if j+1 >= len(cards) {
					break
				}
			}
		}
	}

	for _, curCard := range cards {
		totalCards += curCard.Copies
	}

	fmt.Println("Part two answer: ", totalCards)
}

func GetCards() []Card {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cards []Card

	for scanner.Scan() {
		cards = append(cards, GetCard(scanner.Text()))
	}
	return cards
}

func GetCard(inStr string) Card {
	var curCard Card
	curCard.Id = GetCardId(inStr)
	curCard.WinningNumbers = GetWinningNumbers(inStr)
	curCard.PickedNumbers = GetPickedNumbers(inStr)

	return curCard
}

func GetCardId(inStr string) int {
	firstSlice := strings.Split(inStr, ":")
	secondSlice := strings.Split(firstSlice[0], " ")

	val, err := strconv.ParseInt(secondSlice[1], 10, 32)

	if err == nil {
		return int(val)
	}

	return 0
}

func GetWinningNumbers(inStr string) []int {
	var ret []int

	firstSlice := strings.Split(inStr, ":")
	secondSlice := strings.Split(firstSlice[1], "|")
	rawNums := secondSlice[0]
	rawNums = strings.Trim(rawNums, " ")

	strSlcNums := strings.Split(rawNums, " ")

	for _, curNumStr := range strSlcNums {
		val, err := strconv.ParseInt(curNumStr, 10, 32)
		if err == nil {
			ret = append(ret, int(val))
		}
	}

	return ret
}

func GetPickedNumbers(inStr string) []int {
	var ret []int

	firstSlice := strings.Split(inStr, ":")
	secondSlice := strings.Split(firstSlice[1], "|")
	rawNums := secondSlice[1]
	rawNums = strings.Trim(rawNums, " ")

	strSlcNums := strings.Split(rawNums, " ")

	for _, curNumStr := range strSlcNums {
		val, err := strconv.ParseInt(curNumStr, 10, 32)
		if err == nil {
			ret = append(ret, int(val))
		}
	}

	return ret
}

func CalculateScore(winningNums []int, pickedNums []int) int {
	cardScore := 0
	winCount := CountWins(winningNums, pickedNums)

	if winCount == 0 {
		return 0
	}

	for i := 1; i <= winCount; i++ {
		if cardScore == 0 {
			cardScore = 1
		} else {
			cardScore = cardScore * 2
		}
	}

	return cardScore
}

func CountWins(winningNums []int, pickedNums []int) int {
	winCount := 0

	for _, curPicked := range pickedNums {
		for _, curWinning := range winningNums {
			if curPicked == curWinning {
				winCount++
			}
		}
	}

	return winCount
}
