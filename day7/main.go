package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards    []string
	Wager    int
	Strength int
	Rank     int
}

var CardScore = func() map[string]int {
	return map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
}

var HandScore = func() map[string]int {
	return map[string]int{
		"HighCard":     0,
		"OnePair":      1,
		"TwoPair":      2,
		"ThreeOfAKind": 3,
		"FullHouse":    4,
		"FourOfAKind":  5,
		"FiveOfAKind":  6,
	}
}

func main() {
	fmt.Println("Advent of code 2023: Day7")
	hands := GetInputs("inputs.txt")
	PartOne(hands)
}

func GetInputs(fileName string) []Hand {
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

func ParseLines(fileLines []string) []Hand {
	var returnObj []Hand

	for _, currentLine := range fileLines {
		slcLine := strings.Split(currentLine, " ")
		slcCards := strings.Split(slcLine[0], "")
		val, _ := strconv.ParseInt(slcLine[1], 10, 64)
		currentHand := Hand{Cards: slcCards, Wager: int(val)}
		returnObj = append(returnObj, currentHand)
	}

	return returnObj
}

func PartOne(hands []Hand) {
	fmt.Println("Starting Part One")
	var answer int
	var calculatedHands []Hand

	for _, currentHand := range hands {
		calculatedHands = append(calculatedHands, CalculateHandScore(currentHand))
	}
	sort.Slice(calculatedHands, func(i, j int) bool {
		var returnObj bool
		if calculatedHands[i].Strength == calculatedHands[j].Strength {
			//walk the cards in each hand to see when one has a higher card
			for index := range calculatedHands[i].Cards {
				if CardScore()[calculatedHands[i].Cards[index]] == CardScore()[calculatedHands[j].Cards[index]] {
					continue
				} else {
					returnObj = CardScore()[calculatedHands[i].Cards[index]] < CardScore()[calculatedHands[j].Cards[index]]
					break
				}
			}
		} else {
			returnObj = calculatedHands[i].Strength < calculatedHands[j].Strength
		}

		return returnObj
	})

	for _, curHand := range calculatedHands {
		fmt.Println(curHand)
	}

	answer = CalculateTotalScore(calculatedHands)
	fmt.Println("Answer: ", answer)
}

func CalculateHandScore(hand Hand) Hand {
	var returnObj Hand
	returnObj.Strength = CalculateHandType(hand.Cards)
	returnObj.Wager = hand.Wager
	returnObj.Cards = hand.Cards
	return returnObj
}

func CalculateHandType(cards []string) int {
	countPair := 0
	countTriple := 0
	countQuad := 0
	countQuint := 0

	cardDict := make(map[string]int)
	for _, card := range cards {
		cardDict[card]++
	}

	for _, count := range cardDict {
		switch count {
		case 2:
			countPair++
		case 3:
			countTriple++
		case 4:
			countQuad++
		case 5:
			countQuint++
		}
	}

	if countPair == 1 {
		return HandScore()["OnePair"]
	}
	if countPair == 2 {
		return HandScore()["TwoPair"]
	}
	if countTriple == 1 && countPair == 0 {
		return HandScore()["ThreeOfAKind"]
	}
	if countTriple == 1 && countPair == 1 {
		return HandScore()["FullHouse"]
	}
	if countQuad == 1 {
		return HandScore()["FourOfAKind"]
	}
	if countQuint == 1 {
		return HandScore()["FiveOfAKind"]
	} else {
		return HandScore()["HighCard"]
	}

}

func CalculateTotalScore(hands []Hand) int {
	var returnObj int

	for i := 0; i < len(hands); i++ {
		returnObj += (i + 1) * hands[i].Wager
	}

	return returnObj
}
