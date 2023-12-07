package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var (
	cardValues1 = map[rune]int{
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	cardValues2 = map[rune]int{
		'J': 1,
		'T': 10,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %s", err.Error())
	}
	fileScanner := bufio.NewScanner(inputFile)
	var hands []handData
	for fileScanner.Scan() {
		play := strings.Fields(fileScanner.Text())
		bid, err := strconv.Atoi(play[1])
		if err != nil {
			log.Fatalf(err.Error())
		}
		hands = append(hands, handData{
			cards:     play[0],
			bid:       bid,
			handType1: getHandType1(play[0]),
			handType2: getHandType2(play[0]),
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].isLowerThan1(hands[j])
	})
	sum1 := 0
	for i, hand := range hands {
		sum1 += (i + 1) * hand.bid
	}
	fmt.Println(sum1)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].isLowerThan2(hands[j])
	})
	sum2 := 0
	for i, hand := range hands {
		sum2 += (i + 1) * hand.bid
	}
	fmt.Println(sum2)
}

func getHandType1(hand string) int {
	cardCounts := map[rune]int{}
	for _, card := range hand {
		n, exists := cardCounts[card]
		if exists {
			cardCounts[card] = n + 1
		} else {
			cardCounts[card] = 1
		}
	}
	switch len(cardCounts) {
	case 1:
		return 0 // 5 of a kind
	case 2:
		for _, count := range cardCounts {
			if count == 4 || count == 1 {
				return 1 // 4 of a kind
			}
		}
		return 2 // full house
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return 3 // 3 of a kind
			}
		}
		return 4 // two pair
	case 4:
		return 5 // one pair
	case 5:
		return 6 // high card
	}
	return 7
}

func getHandType2(hand string) int {
	cardCounts := map[rune]int{}
	for _, card := range hand {
		n, exists := cardCounts[card]
		if exists {
			cardCounts[card] = n + 1
		} else {
			cardCounts[card] = 1
		}
	}
	j := cardCounts['J']
	if j == 0 {
		return getHandType1(hand)
	}
	switch len(cardCounts) {
	case 1:
		return 0 // 5 of a kind J
	case 2:
		return 0 // 5 of a kind
	case 3:
		for card, count := range cardCounts {
			if card != 'J' {
				if count+j == 4 {
					return 1 // 4 of a kind
				}
			}
		}
		return 2 // full house
	case 4:
		for _, count := range cardCounts {
			if count+j == 3 {
				return 3 // 3 of a kind
			}
		}
		return 4 // two pair
	case 5:
		return 5 // one pair
	}
	return 6
}

func getCardValue1(card rune) int {
	if unicode.IsDigit(card) {
		return int(card - '0')
	}
	return cardValues1[card]
}

func getCardValue2(card rune) int {
	if unicode.IsDigit(card) {
		return int(card - '0')
	}
	return cardValues2[card]
}

type handData struct {
	cards     string
	bid       int
	handType1 int
	handType2 int
}

func (h1 handData) isLowerThan1(h2 handData) bool {
	if h1.handType1 != h2.handType1 {
		return h1.handType1 > h2.handType1
	}
	for i := 0; i < len(h1.cards); i++ {
		c1, c2 := rune(h1.cards[i]), rune(h2.cards[i])
		if getCardValue1(c1) == getCardValue1(c2) {
			continue
		}
		return getCardValue1(c1) < getCardValue1(c2)
	}
	return true
}

func (h1 handData) isLowerThan2(h2 handData) bool {
	if h1.handType2 != h2.handType2 {
		return h1.handType2 > h2.handType2
	}
	for i := 0; i < len(h1.cards); i++ {
		c1, c2 := rune(h1.cards[i]), rune(h2.cards[i])
		if getCardValue2(c1) == getCardValue2(c2) {
			continue
		}
		return getCardValue2(c1) < getCardValue2(c2)
	}
	return true
}
