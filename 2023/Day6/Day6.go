package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	if !fileScanner.Scan() {
		log.Fatalf("Unable to read file")
	}
	times := strings.Fields(fileScanner.Text())
	if !fileScanner.Scan() {
		log.Fatalf("Unable to read file")
	}
	distances := strings.Fields(fileScanner.Text())
	prod := 1
	timeString := ""
	distString := ""
	for i := 1; i < len(times); i++ {
		t, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatalf(err.Error())
		}
		timeString += times[i]
		d, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatalf(err.Error())
		}
		distString += distances[i]
		countWin := 0
		for t1 := 0; t1 <= t; t1++ {
			dist := (t - t1) * t1
			if dist > d {
				countWin++
			}
		}
		prod *= countWin
	}
	fmt.Println(prod)
	T, err := strconv.Atoi(timeString)
	if err != nil {
		log.Fatalf(err.Error())
	}
	D, err := strconv.Atoi(distString)
	if err != nil {
		log.Fatalf(err.Error())
	}
	time, dist := float64(T), float64(D)
	t1 := int(math.Ceil((time - math.Sqrt(math.Pow(time, 2)-4*dist)) / 2))
	t2 := int(math.Floor((time + math.Sqrt(math.Pow(time, 2)-4*dist)) / 2))
	fmt.Println(t2 - t1 + 1)
}
