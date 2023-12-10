package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

func powerOfTwo(powerNum int) int {
	var rez int = 1 << powerNum
	return rez
}

func main() {
	re := regexp.MustCompile("[0-9]+")
	file, err := os.Open("longInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int = 0
	var winningCards [][]int
	var index int = 0
	for scanner.Scan() {
		winningCards = append(winningCards, []int{})
		game, _ := strconv.Atoi(re.FindString(s.Split(scanner.Text(), ": ")[0]))
		// fmt.Printf("game: %d\n", game)
		var input string = s.Split(scanner.Text(), ": ")[1]
		var winners []string = re.FindAllString(s.Split(input, " | ")[0], -1)
		var myList []string = re.FindAllString(s.Split(input, " | ")[1], -1)
		var numOfWinners int = 0
		for _, el := range winners {
			for _, elx := range myList {
				if el == elx {
					numOfWinners++
					winningCards[index] = append(winningCards[index], game+numOfWinners)
					break
				}
			}
		}
		if numOfWinners > 0 {
			sum += powerOfTwo(numOfWinners - 1)
		}
		index++
	}
	// fmt.Println(sum)
	// fmt.Println(winningCards)
	// for _, el := range winningCards {
	// 	for _, elx := range el {
	// 		fmt.Println(elx)
	// 		winningCards = append(winningCards, winningCards[elx-1])
	// 		secondSum++
	// 	}
	// fmt.Println(winningCards)
	// }
	var length int = len(winningCards)
	for i := 0; i < length; i++ {
		length = len(winningCards)
		for _, elx := range winningCards[i] {
			// fmt.Println(elx)
			winningCards = append(winningCards, winningCards[elx-1])
		}
	}
	fmt.Println(len(winningCards))
}
