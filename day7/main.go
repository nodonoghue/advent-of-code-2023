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
	Cards     []string
	Wager     int
	Rank      int
	CardScore int
}

type HandRank struct {
	Rank      int
	RankScore int
}

type GroupScore struct {
	Card  string
	Count int
	Score int
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

	answer = CalculateTotalScore(calculatedHands)
	fmt.Println("Answer: ", answer)
}

func PartTwo(hands []Hand) {
	fmt.Println("Starting Part Two")

}

func CalculateHandScore(hand Hand) Hand {
	var returnObj Hand

	rankObj := CalculateHandType(hand.Cards)
	returnObj.Rank = rankObj.Rank
	returnObj.CardScore = rankObj.RankScore
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

	for card, count := range cardDict {
		switch count {
		case 2:
			returnObj.RankScore += 2 * CardScore()[card]
			countPair++
		case 3:
			returnObj.RankScore += 3 * CardScore()[card]
			countTriple++
		case 4:
			returnObj.RankScore += 4 * CardScore()[card]
			countQuad++
		case 5:
			returnObj.RankScore += 5 * CardScore()[card]
			countQuint++
		}
	}

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

func CalculateHandTypeWild(cards []string) HandRank {
	//J cards are wild, need to remove the J cards, then count the numver of each card
	// find the highest scoring group and add all Js to it to make the best possible hand
	// How to handle Full house?  if a Hand has 2 pairs and 1 J, add a J to the highest scoring
	// group
}

func CalculateTotalScore(hands []Hand) int {
	var returnObj int

	for i := 0; i < len(hands); i++ {
		returnObj += (i + 1) * hands[i].Wager
	}

	return returnObj
}
