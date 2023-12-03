package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	sum1 := 0
	sum2 := 0
	for fileScanner.Scan() {
		possible := true
		game := strings.Split(fileScanner.Text(), ":")
		id, err := strconv.Atoi(game[0][5:])
		if err != nil {
			log.Fatalf("Unable to open file: %s", err.Error())
		}
		minRed := 0
		minGreen := 0
		minBlue := 0
		picks := strings.Split(game[1], ";")
		for _, pick := range picks {
			for _, cubes := range strings.Split(pick, ",") {
				data := strings.Fields(cubes)
				n, err := strconv.Atoi(data[0])
				if err != nil {
					log.Fatalf(err.Error())
				}
				colour := data[1]
				maxCubes := 0
				switch colour {
				case "red":
					maxCubes = maxRed
					minRed = max(minRed, n)
				case "green":
					maxCubes = maxGreen
					minGreen = max(minGreen, n)
				case "blue":
					maxCubes = maxBlue
					minBlue = max(minBlue, n)
				}
				if n > maxCubes {
					possible = false
				}
			}
		}
		if possible {
			sum1 += id
		}
		sum2 += minRed * minGreen * minBlue
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
