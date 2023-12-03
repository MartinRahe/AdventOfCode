package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var (
	directions = []direction{
		{
			x: 0,
			y: 1,
		},
		{
			x: 1,
			y: 1,
		},
		{
			x: -1,
			y: 1,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: 1,
			y: -1,
		},
		{
			x: -1,
			y: -1,
		},
		{
			x: 1,
			y: 0,
		},
		{
			x: -1,
			y: 0,
		},
	}

	numberList   []int
	numberPosMap [][]int
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
	sum1 := 0
	numberPosMap = make([][]int, len(arr))
	numberCount := -1
	for i := 0; i < len(arr); i++ {
		row := arr[i]
		numberPosMap[i] = make([]int, len(row))
		for j := 0; j < len(row); j++ {
			numberPosMap[i][j] = -1
		}
		number := ""
		counts := false
		for j := 0; j < len(row); j++ {
			symbol := rune(row[j])
			if unicode.IsDigit(symbol) {
				number += string(symbol)
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Fatalf(err.Error())
				}
				if len(number) == 1 {
					numberCount += 1
					numberList = append(numberList, n)
				} else {
					numberList[numberCount] = n
				}
				numberPosMap[i][j] = numberCount
				for _, d := range directions {
					if i+d.y >= 0 && i+d.y < len(arr) && j+d.x >= 0 && j+d.x < len(row) {
						neighbour := rune(arr[i+d.y][j+d.x])
						if neighbour != '.' && !unicode.IsDigit(neighbour) {
							counts = true
						}
					}
				}
			} else {
				if counts {
					n, err := strconv.Atoi(number)
					if err != nil {
						log.Fatalf(err.Error())
					}
					sum1 += n
				}
				number = ""
				counts = false
			}
		}
		if counts {
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf(err.Error())
			}
			sum1 += n
		}
		number = ""
		counts = false
	}
	fmt.Println(sum1)
	sum2 := 0
	for i := 0; i < len(arr); i++ {
		row := arr[i]
		for j := 0; j < len(row); j++ {
			symbol := rune(row[j])
			if symbol == '*' {
				numbers := map[int]bool{}
				for _, d := range directions {
					if i+d.y >= 0 && i+d.y < len(arr) && j+d.x >= 0 && j+d.x < len(row) {
						neighbourNumberPos := numberPosMap[i+d.y][j+d.x]
						if neighbourNumberPos != -1 {
							numbers[numberList[neighbourNumberPos]] = true
						}
					}
				}
				if len(numbers) == 2 {
					p := 1
					for n := range numbers {
						p *= n
					}
					sum2 += p
				}
			}
		}
	}
	fmt.Println(sum2)
}

type direction struct {
	x int
	y int
}
