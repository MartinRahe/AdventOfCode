package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	seeds                 []int
	seedToSoil            []almanacMapping
	soilToFertilizer      []almanacMapping
	fertilizerToWater     []almanacMapping
	waterToLight          []almanacMapping
	lightToTemperature    []almanacMapping
	temperatureToHumidity []almanacMapping
	humidityToLocation    []almanacMapping
	seedToLocation        = map[int]int{}
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	mapName := ""
	for fileScanner.Scan() {
		info := strings.Fields(fileScanner.Text())
		if len(info) == 0 {
			continue
		}
		if info[0] == "seeds:" {
			for _, s := range info[1:] {
				seed, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalf(err.Error())
				}
				seeds = append(seeds, seed)
			}
			continue
		}
		if len(info) == 2 {
			mapName = info[0]
			continue
		}
		destStart, err := strconv.Atoi(info[0])
		if err != nil {
			log.Fatalf(err.Error())
		}
		sourceStart, err := strconv.Atoi(info[1])
		if err != nil {
			log.Fatalf(err.Error())
		}
		rangeLen, err := strconv.Atoi(info[2])
		if err != nil {
			log.Fatalf(err.Error())
		}
		switch mapName {
		case "seed-to-soil":
			assignToMap(&seedToSoil, destStart, sourceStart, rangeLen)
		case "soil-to-fertilizer":
			assignToMap(&soilToFertilizer, destStart, sourceStart, rangeLen)
		case "fertilizer-to-water":
			assignToMap(&fertilizerToWater, destStart, sourceStart, rangeLen)
		case "water-to-light":
			assignToMap(&waterToLight, destStart, sourceStart, rangeLen)
		case "light-to-temperature":
			assignToMap(&lightToTemperature, destStart, sourceStart, rangeLen)
		case "temperature-to-humidity":
			assignToMap(&temperatureToHumidity, destStart, sourceStart, rangeLen)
		case "humidity-to-location":
			assignToMap(&humidityToLocation, destStart, sourceStart, rangeLen)
		}
	}
	for _, seed := range seeds {
		seedToLocation[seed] = getFromMap(&humidityToLocation,
			getFromMap(&temperatureToHumidity,
				getFromMap(&lightToTemperature,
					getFromMap(&waterToLight,
						getFromMap(&fertilizerToWater,
							getFromMap(&soilToFertilizer,
								getFromMap(&seedToSoil, seed)))))))
	}
	minLoc1 := seedToLocation[seeds[0]]
	for _, seed := range seeds {
		minLoc1 = min(minLoc1, seedToLocation[seed])
	}
	fmt.Println(minLoc1)
	minLoc2 := seedToLocation[seeds[0]]
	for i := 0; i < len(seeds); i += 2 {
		startSeed := seeds[i]
		rangeLen := seeds[i+1]
		for j := 0; j < rangeLen; j++ {
			minLoc2 = min(minLoc2, getFromMap(&humidityToLocation,
				getFromMap(&temperatureToHumidity,
					getFromMap(&lightToTemperature,
						getFromMap(&waterToLight,
							getFromMap(&fertilizerToWater,
								getFromMap(&soilToFertilizer,
									getFromMap(&seedToSoil, startSeed+j))))))))
		}
	}
	fmt.Println(minLoc2)
}

func assignToMap(m *[]almanacMapping, dStart, sStart, rLen int) {
	*m = append(*m, almanacMapping{
		destStart:   dStart,
		sourceStart: sStart,
		rangeLen:    rLen,
	})
}

func getFromMap(m *[]almanacMapping, q int) int {
	for _, mapping := range *m {
		if q >= mapping.sourceStart && q < mapping.sourceStart+mapping.rangeLen {
			return mapping.destStart + q - mapping.sourceStart
		}
	}
	return q
}

type almanacMapping struct {
	destStart   int
	sourceStart int
	rangeLen    int
}
