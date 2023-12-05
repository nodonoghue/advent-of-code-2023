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

type SeedRange struct {
	StartNum int
	Length   int
}

type Range struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func main() {
	fmt.Println("Advent of Code 2023 day 5")
	PartOne()
	PartTwo()
}

func PartOne() {
	file, err := os.Open("inputs.txt")

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

	//Walk through the steps to get to the location for each seed.
	location := 0
	for _, curSeed := range seeds.SeedNum {
		//Seed to soil
		soilNum := GetResult(curSeed, seedtoSoil)
		fertilizerNum := GetResult(soilNum, soiltoFertilizer)
		waterNum := GetResult(fertilizerNum, fertilizertoWater)
		lightNum := GetResult(waterNum, watertoLight)
		tempNumber := GetResult(lightNum, lighttoTemp)
		humidityNum := GetResult(tempNumber, temptoHumidity)
		locationNum := GetResult(humidityNum, humiditytoLocation)

		if location == 0 {
			location = locationNum
		} else if locationNum < location {
			location = locationNum
		}
	}

	fmt.Println("Part 1 closest location: ", location)
}

func PartTwo() {
	file, err := os.Open("inputs.txt")

	//Declare needed structs
	var seedRanges []SeedRange
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
				seedRanges = GetSeedRanges(scanner.Text())
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

	//Walk through the steps to get to the location for each seed.
	//locChannel := make(chan int)
	//var wg sync.WaitGroup
	var location int = 0
	groupNum := 1
	for _, curSeedRange := range seedRanges {
		//go GetLocation(curSeedRange, seedtoSoil, soiltoFertilizer, fertilizertoWater, watertoLight, lighttoTemp, temptoHumidity, humiditytoLocation, locChannel)
		//wg.Add(1)
		fmt.Println(curSeedRange)
		fmt.Println("Starting group number ", groupNum)

		for curSeedNum := curSeedRange.StartNum; curSeedNum < (curSeedRange.StartNum + curSeedRange.Length); curSeedNum++ {
			//Seed to soil
			soilNum := GetResult(curSeedNum, seedtoSoil)
			fertilizerNum := GetResult(soilNum, soiltoFertilizer)
			waterNum := GetResult(fertilizerNum, fertilizertoWater)
			lightNum := GetResult(waterNum, watertoLight)
			tempNumber := GetResult(lightNum, lighttoTemp)
			humidityNum := GetResult(tempNumber, temptoHumidity)
			locationNum := GetResult(humidityNum, humiditytoLocation)

			if location == 0 {
				location = locationNum
			} else if locationNum < location {
				location = locationNum
			}
		}
		groupNum++
	}

	//wg.Wait()
	//location := <-locChannel

	fmt.Println("Part 2 closest location: ", location)
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

func GetSeedRanges(inStr string) []SeedRange {
	var retSeeds []SeedRange
	slcStr := strings.Split(inStr, ":")
	numStr := strings.Trim(slcStr[1], " ")
	slcNums := strings.Split(numStr, " ")
	slcLen := len(slcNums)

	for i := 0; i < slcLen; i += 2 {
		var curSeedRange SeedRange

		startVal, startErr := strconv.ParseInt(slcNums[i], 10, 64)
		lenVal, lenErr := strconv.ParseInt(slcNums[i+1], 10, 64)

		if startErr == nil && lenErr == nil {
			curSeedRange = SeedRange{int(startVal), int(lenVal)}
			retSeeds = append(retSeeds, curSeedRange)
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

func GetResult(inNum int, checkRanges []Range) int {
	retNum := 0
	//Determine rules for source to destination mapping here
	//looks like the difference between the incoming number and the start of the
	//contained range affects the destination in the same way
	for _, curRange := range checkRanges {
		if inNum < curRange.SourceStart+curRange.Length && inNum >= curRange.SourceStart {
			retNum = curRange.DestinationStart + (inNum - curRange.SourceStart)
			break
		}
	}

	if retNum == 0 {
		retNum = inNum
	}

	return retNum
}
