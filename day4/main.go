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
}

func main() {
	fmt.Println("Advent of Code 2023 Day 4")
}

func PartOne() {
	file, err := os.Open("test.txt")

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

	var totalPoints int
	for _, curCard := range cards {
		fmt.Println(curCard)
	}
	fmt.Println("Part one answer: ", totalPoints)
}

func GetCard(inStr string) Card {
	var curCard Card

	curCard.Id = GetCardId(inStr)

	//Get winning numbers

	//Get picked numbers

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

	return ret
}

func GetPickedNumbers(inStr string) []int {
	var ret []int

	return ret
}
