package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seeds struct {
	SeedNum []int
}

type Range struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func main() {
	fmt.Println("Advent of Code 2023 day 5")
	PartOne()
}

func PartOne() {
	file, err := os.Open("test.txt")

	//Declare needed structs
	var seeds Seeds
	var seedtoSoil []Range
	var soiltoFertilizer []Range
	var fertilizertoWater []Range
	var watertoLight []Range
	var lighttoTemp []Range
	var temptoHumidity []Range
	var humiditytoLocation []Range
	var fileLocation string

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		curLine := scanner.Text()
		if len(curLine) > 0 {
			if strings.Contains(curLine, "seeds") {
				seeds.SeedNum = GetSeedNums(scanner.Text())
			}
			if strings.Contains(curLine, "seed-to-soil") {
				fileLocation = "seed-to-soil"
				continue
			}
			if strings.Contains(curLine, "soil-to-fertilizer") {
				fileLocation = "soil-to-fertilizer"
				continue
			}
			if strings.Contains(curLine, "fertilizer-to-water") {
				fileLocation = "fertilizer-to-water"
				continue
			}
			if strings.Contains(curLine, "water-to-light") {
				fileLocation = "water-to-light"
				continue
			}
			if strings.Contains(curLine, "light-to-temperature") {
				fileLocation = "light-to-temperature"
				continue
			}
			if strings.Contains(curLine, "temperature-to-humidity") {
				fileLocation = "temperature-to-humidity"
				continue
			}
			if strings.Contains(curLine, "humidity-to-location") {
				fileLocation = "humidity-to-location"
				continue
			}

			switch fileLocation {
			case "seed-to-soil":
				seedtoSoil = append(seedtoSoil, ParseRange(curLine))
			case "soil-to-fertilizer":
				soiltoFertilizer = append(soiltoFertilizer, ParseRange(curLine))
			case "fertilizer-to-water":
				fertilizertoWater = append(fertilizertoWater, ParseRange(curLine))
			case "water-to-light":
				watertoLight = append(watertoLight, ParseRange(curLine))
			case "light-to-temperature":
				lighttoTemp = append(lighttoTemp, ParseRange(curLine))
			case "temperature-to-humidity":
				temptoHumidity = append(temptoHumidity, ParseRange(curLine))
			case "humidity-to-location":
				humiditytoLocation = append(humiditytoLocation, ParseRange(curLine))
			}
		}
	}

	fmt.Println("seeds", seeds)
	fmt.Println("seed to soil: ", seedtoSoil)
	fmt.Println("soil to fertilizer: ", soiltoFertilizer)
	fmt.Println("fertilizer to water: ", fertilizertoWater)
	fmt.Println("water to light: ", watertoLight)
	fmt.Println("light to temp: ", lighttoTemp)
	fmt.Println("temp to humidity: ", temptoHumidity)
	fmt.Println("humidity to location", humiditytoLocation)

	//Walk through the steps to get to the location for each seed.
	for _, curSeed := range seeds.SeedNum {
		fmt.Println("current seed num: ", curSeed)
	}
}

func GetSeedNums(inStr string) []int {
	var retSeeds []int
	slcStr := strings.Split(inStr, ":")
	numStr := strings.Trim(slcStr[1], " ")
	slcNums := strings.Split(numStr, " ")
	for _, curSeedNum := range slcNums {
		value, err := strconv.ParseInt(curSeedNum, 10, 64)

		if err == nil {
			retSeeds = append(retSeeds, int(value))
		}
	}
	return retSeeds
}

func ParseRange(inStr string) Range {
	var curRange Range
	slcStr := strings.Split(inStr, " ")

	dest, destErr := strconv.ParseInt(slcStr[0], 10, 64)
	if destErr == nil {
		curRange.DestinationStart = int(dest)
	}

	src, srcErr := strconv.ParseInt(slcStr[1], 10, 64)
	if srcErr == nil {
		curRange.SourceStart = int(src)
	}

	len, lenErr := strconv.ParseInt(slcStr[2], 10, 64)
	if lenErr == nil {
		curRange.Length = int(len)
	}

	return curRange
}
