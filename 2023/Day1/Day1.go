package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	inputFile, err := os.Open("input.txt")
	writtenDigits := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	sum1 := 0
	sum2 := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		nums1 := make([]rune, 0)
		for _, r := range value {
			if unicode.IsDigit(r) {
				nums1 = append(nums1, r)
			}
		}
		if len(nums1) != 0 {
			n1, err := strconv.Atoi(string(nums1[0]) + string(nums1[len(nums1)-1]))
			if err != nil {
				log.Fatalf(err.Error())
			}
			sum1 += n1
		}

		num2 := ""
		for i := 0; i < len(value); i++ {
			foundWritten := false
			if unicode.IsDigit(rune(value[i])) {
				num2 += string(value[i])
				break
			}
			for number := range writtenDigits {
				if i+len(number) <= len(value) && value[i:i+len(number)] == number {
					num2 += string(writtenDigits[number])
					foundWritten = true
					break
				}
			}
			if foundWritten {
				break
			}
		}
		for i := len(value); i > 0; i-- {
			foundWritten := false
			if unicode.IsDigit(rune(value[i-1])) {
				num2 += string(value[i-1])
				break
			}
			for number := range writtenDigits {
				if i-len(number) >= 0 && value[i-len(number):i] == number {
					num2 += string(writtenDigits[number])
					foundWritten = true
					break
				}
			}
			if foundWritten {
				break
			}
		}
		n2, err := strconv.Atoi(num2)
		if err != nil {
			log.Fatalf(err.Error())
		}
		sum2 += n2
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
