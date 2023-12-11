package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func sliceContains(slice []string) bool {
	for _, el := range slice {
		if strings.Contains(el, "#") {
			return true
		}
	}
	return false
}

func abs(num int) int {
	if num < 0 {
		return num * (-1)
	}
	return num
}

func main() {
	file, err := os.Open("longInput.txt")
	// file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var galaxies [][]string
	for scanner.Scan() {
		var line string = scanner.Text()
		galaxies = append(galaxies, strings.Split(line, ""))
	}
	var indexes []int
	var indexesToIncrease []int
	for i, el := range galaxies {
		for j, elx := range el {
			if elx == "#" {
				indexes = append(indexes, i)
				indexesToIncrease = append(indexesToIncrease, 0)
				indexes = append(indexes, j)
				indexesToIncrease = append(indexesToIncrease, 0)
			}
		}
	}
	for i := 0; i < len(galaxies); i++ {
		if !sliceContains(galaxies[i]) {
			for j := 0; j < len(indexes); j += 2 {
				if indexes[j] > i {
					indexesToIncrease[j]++
				}
			}
		}
	}
	for i := 0; i < len(galaxies[0]); i++ {
		var hasGalaxy bool = false
		for _, el := range galaxies {
			if el[i] == "#" {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for j := 1; j < len(indexes); j += 2 {
				if indexes[j] > i {
					indexesToIncrease[j]++
				}
			}
		}
	}
	for i := range indexesToIncrease {
		indexes[i] += indexesToIncrease[i] * 999999
	}
	var sum int = 0
	for i := 0; i < len(indexes); i += 2 {
		for j := i + 2; j < len(indexes); j += 2 {
			sum += abs(indexes[j]-indexes[i]) + abs(indexes[j+1]-indexes[i+1])
		}
	}
	fmt.Println(sum)
}
