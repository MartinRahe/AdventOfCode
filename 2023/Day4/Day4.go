package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	var arr []string
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}
	copyCounts := make([]int, len(arr))
	sum1 := 0
	sum2 := len(arr)
	for i := 0; i < len(arr); i++ {
		card := strings.FieldsFunc(arr[i], func(r rune) bool {
			return r == ':' || r == '|'
		})
		winners := strings.Fields(card[1])
		have := strings.Fields(card[2])
		cardSum := 0
		winCount := 0
		for _, number := range have {
			if stringContains(winners, number) {
				if cardSum == 0 {
					cardSum = 1
				} else {
					cardSum *= 2
				}
				winCount += 1
			}
		}
		for j := 0; j < winCount; j++ {
			if i+j+1 < len(arr) {
				copyCounts[i+j+1] += 1 + copyCounts[i]
			}
		}
		sum1 += cardSum
		sum2 += copyCounts[i]
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func stringContains(lsd []string, element string) bool {
	for _, el := range lsd {
		if el == element {
			return true
		}
	}
	return false
}
