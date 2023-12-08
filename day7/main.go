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
	Cards      []string
	Wager      int
	Rank       int
	RankSource map[string]int
}

type HandRank struct {
	Rank       int
	RankSource map[string]int
}

var CardScore = func() map[string]int {
	return map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
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
	hands := GetInputs("test.txt")
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
		if calculatedHands[i].Rank == calculatedHands[j].Rank {
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
			returnObj = calculatedHands[i].Rank < calculatedHands[j].Rank
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

	rankObj := CalculateHandType(hand.Cards)
	fmt.Println(rankObj)
	returnObj.Rank = rankObj.Rank
	returnObj.RankSource = rankObj.RankSource
	returnObj.Wager = hand.Wager
	returnObj.Cards = hand.Cards

	return returnObj
}

func CalculateHandType(cards []string) HandRank {
	var returnObj HandRank
	//Need to add variations to the ranks based on which cards are in each group.
	// 4 Ts should never be ranked about 4 Ks
	countPair := 0
	countTriple := 0
	countQuad := 0
	countQuint := 0

	cardDict := make(map[string]int)
	for _, card := range cards {
		cardDict[card]++
	}

	rankSource := make(map[string]int)
	for card, count := range cardDict {
		switch count {
		case 2:
			rankSource[card] = 2
			countPair++
		case 3:
			rankSource[card] = 3
			countTriple++
		case 4:
			rankSource[card] = 4
			countQuad++
		case 5:
			rankSource[card] = 5
			countQuint++
		}
	}
	fmt.Println(cards)
	fmt.Println(rankSource)

	if countPair == 1 {
		returnObj.Rank = HandScore()["OnePair"]
	}
	if countPair == 2 {
		returnObj.Rank = HandScore()["TwoPair"]
	}
	if countTriple == 1 && countPair == 0 {
		returnObj.Rank = HandScore()["ThreeOfAKind"]
	}
	if countTriple == 1 && countPair == 1 {
		returnObj.Rank = HandScore()["FullHouse"]
	}
	if countQuad == 1 {
		returnObj.Rank = HandScore()["FourOfAKind"]
	}
	if countQuint == 1 {
		returnObj.Rank = HandScore()["FiveOfAKind"]
	}
	if countPair == 0 && countTriple == 0 && countQuad == 0 && countQuint == 0 {
		returnObj.Rank = HandScore()["HighCard"]
	}

	return returnObj
}

func CalculateTotalScore(hands []Hand) int {
	var returnObj int

	for i := 0; i < len(hands); i++ {
		returnObj += (i + 1) * hands[i].Wager
	}

	return returnObj
}
