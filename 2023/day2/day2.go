package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
)

func checkIfMore(game string) bool {
	var sets []string = s.Split(game, ";")
	for _, el := range sets {
		var cubes []string = s.Split(el, ",")
		for _, elx := range cubes {
			var cubeAndColor []string = s.Split(elx, " ")
			num, _ := strconv.Atoi(cubeAndColor[1])
			switch cubeAndColor[2] {
			case "red":
				if num > 12 {
					return false
				}
			case "green":
				if num > 13 {
					return false
				}
			case "blue":
				if num > 14 {
					return false
				}
			}
		}
	}
	return true
}

func findMaximal(game string) (int, int, int) {
	var maxRed int = 0
	var maxGreen int = 0
	var maxBlue int = 0
	var sets []string = s.Split(game, ";")
	for _, el := range sets {
		var cubes []string = s.Split(el, ",")
		for _, elx := range cubes {
			var cubeAndColor []string = s.Split(elx, " ")
			num, _ := strconv.Atoi(cubeAndColor[1])
			switch cubeAndColor[2] {
			case "red":
				if num > maxRed {
					maxRed = num
				}
			case "green":
				if num > maxGreen {
					maxGreen = num
				}
			case "blue":
				if num > maxBlue {
					maxBlue = num
				}
			}
		}
	}
	return maxRed, maxGreen, maxBlue
}

func main() {
	var sum int = 0
	var sumPower int = 0
	file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var input string = scanner.Text()
		var group []string = s.Split(input, ":")
		groupId, err := strconv.Atoi(s.Split(group[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		if checkIfMore(group[1]) {
			sum += groupId
		}
		maxRed, maxGreen, maxBlue := findMaximal(group[1])
		sumPower += maxRed * maxGreen * maxBlue
	}
	fmt.Println(sum)
	fmt.Println(sumPower)
}
