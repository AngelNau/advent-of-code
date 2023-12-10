package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	s "strings"
)

var fiveOfAKind int = 7
var fourOfAKind int = 6
var fullHouse int = 5
var threeOfAKind int = 4
var twoPair int = 3
var onePair int = 2
var highCard int = 1

func checkFullHouse(hand string, card byte) bool {
	var num int = 0
	var currentCard byte
	for _, el := range hand {
		if byte(el) != card {
			currentCard = byte(el)
		}
	}
	for _, elx := range hand {
		if byte(elx) == currentCard {
			num++
		}
	}
	return num == 2
}

func handStrength(hand string) int {
	var chars = "23456789TJQKA"
	var numOfChar = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, el := range hand {
		var what int = s.Index(chars, string(el))
		numOfChar[what]++
	}
	if slices.Contains(numOfChar, 5) {
		return fiveOfAKind
	}
	if slices.Contains(numOfChar, 4) {
		return fourOfAKind
	}
	if slices.Contains(numOfChar, 3) {
		if slices.Contains(numOfChar, 2) {
			return fullHouse
		}
		return threeOfAKind
	}
	if slices.Contains(numOfChar, 2) {
		var index int = slices.Index(numOfChar, 2)
		if index < len(numOfChar) {
			if slices.Contains(numOfChar[index+1:], 2) {
				return twoPair
			}
		}
		return onePair
	}
	return highCard
}

func strongerHand(myHand string, opponentHand string) bool {
	var cards = "23456789TJQKA"
	for i := 0; i < len(myHand); i++ {
		if myHand[i] == opponentHand[i] {
			continue
		}
		return s.Index(cards, string(myHand[i])) > s.Index(cards, string(opponentHand[i]))
	}
	return false
}

func strongerHandJoker(myHand string, opponentHand string) bool {
	var cards = "J23456789TQKA"
	for i := 0; i < len(myHand); i++ {
		if myHand[i] == opponentHand[i] {
			continue
		}
		return s.Index(cards, string(myHand[i])) > s.Index(cards, string(opponentHand[i]))
	}
	return false
}

func findJokerStrength(hand string) int {
	var maxStrength int = 0
	for _, el := range hand {
		var tempStr string = s.Replace(hand, "J", string(el), -1)
		var num = handStrength(tempStr)
		if maxStrength < num {
			maxStrength = num
		}
	}
	return maxStrength
}

func main() {
	file, err := os.Open("longInput.txt")
	// file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cards []string
	var bids []int
	var cardsStrength []int
	var jokerStrength []int
	for scanner.Scan() {
		cardStr := s.Split(scanner.Text(), " ")[0]
		bidStr := s.Split(scanner.Text(), " ")[1]
		bid, _ := strconv.Atoi(bidStr)
		cards = append(cards, cardStr)
		cardsStrength = append(cardsStrength, handStrength(cardStr))
		jokerStrength = append(jokerStrength, findJokerStrength(cardStr))
		bids = append(bids, bid)
	}
	var sum int = 0
	var sumJoker int = 0
	for i, myHand := range cards {
		var numWorseHands int = 0
		var numWorseHandsJoker int = 0
		for j, opponentHand := range cards {
			if cardsStrength[i] > cardsStrength[j] || (cardsStrength[i] == cardsStrength[j] && strongerHand(myHand, opponentHand)) {
				numWorseHands++
			}
			if jokerStrength[i] > jokerStrength[j] || (jokerStrength[i] == jokerStrength[j] && strongerHandJoker(myHand, opponentHand)) {
				numWorseHandsJoker++
			}
		}
		sum += (numWorseHands + 1) * bids[i]
		sumJoker += (numWorseHandsJoker + 1) * bids[i]
	}
	fmt.Println("==== Part 1 ====")
	fmt.Println(sum)
	fmt.Println("==== Part 2 ====")
	fmt.Println(sumJoker)
}
