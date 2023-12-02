package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	s "strings"
)

func findFirstWord(str string, arrayOfNumStrings [9]string) (int, int) {
	for i := 0; i < len(str); i++ {
		for index, el := range arrayOfNumStrings {
			if i+len(el) < len(str) && s.Contains(str[i:i+len(el)], el) {
				return index + 1, i
			}
		}
	}
	return 0, math.MaxInt
}

func findLastWord(str string, arrayOfNumStrings [9]string) (int, int) {
	for i := len(str) - 1; i >= 0; i-- {
		for index, el := range arrayOfNumStrings {
			if i+len(el) <= len(str) && s.Contains(str[i:i+len(el)], el) {
				return index + 1, i
			}
		}
	}
	return 0, -1
}

func main() {
	var arrayOfNumStrings = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var sum int = 0
	file, err := os.Open("longInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first int = 0
		var last int = 0
		var firstIndex int = 0
		var lastIndex int = 0
		var firstWord int = 0
		var lastWord int = 0
		var firstWordIndex int = 0
		var lastWordIndex int = 0
		var input string = scanner.Text()
		firstWord, firstWordIndex = findFirstWord(scanner.Text(), arrayOfNumStrings)
		lastWord, lastWordIndex = findLastWord(scanner.Text(), arrayOfNumStrings)
		for i := 0; i < len(input); i++ {
			if input[i] >= 48 && input[i] <= 57 {
				first = int(input[i]) - '0'
				firstIndex = i
				break
			}
		}
		for i := len(input) - 1; i >= 0; i-- {
			if input[i] >= 48 && input[i] <= 57 {
				last = int(input[i]) - '0'
				lastIndex = i
				break
			}
		}
		if firstIndex < firstWordIndex {
			sum += (10 * first)
		} else {
			sum += (10 * firstWord)
		}
		if lastIndex > lastWordIndex {
			sum += last
		} else {
			sum += lastWord
		}
	}
	fmt.Println(sum)
}
