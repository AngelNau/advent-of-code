package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"
)

func main() {
	file, err := os.Open("longInput.txt")
	// file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var seeds []uint
	scanner.Scan()
	var firstLine []string = s.Split(s.Split(scanner.Text(), ": ")[1], " ")
	for _, el := range firstLine {
		num, _ := strconv.ParseUint(el, 10, 32)
		seeds = append(seeds, uint(num))
	}
	var maps [][]uint
	var mapsIndex int = -1
	var isMapLine bool = false
	for scanner.Scan() {
		var input string = scanner.Text()
		if len(input) == 0 {
			isMapLine = true
			continue
		}
		if isMapLine {
			mapsIndex++
			maps = append(maps, []uint{})
			isMapLine = false
			continue
		}
		var line []string = s.Split(input, " ")
		for _, el := range line {
			num, _ := strconv.ParseUint(el, 10, 32)
			maps[mapsIndex] = append(maps[mapsIndex], uint(num))
		}
	}
	var minLocation = ^uint(0)
	// for _, seed := range seeds {
	for iter := 0; iter < len(seeds); iter += 2 {
		for seed := seeds[iter]; seed < seeds[iter]+seeds[iter+1]; seed++ {
			// fmt.Printf("Seed: %d", seed)
			var currentNumber uint = seed
			for _, mapp := range maps {
				for i := 0; i < len(mapp); i += 3 {
					if mapp[i+1] <= currentNumber && mapp[i+1]+mapp[i+2] > currentNumber {
						currentNumber = (currentNumber - mapp[i+1]) + mapp[i]
						break
					}
				}
			}
			// fmt.Printf(" Location: %d\n", currentNumber)
			if currentNumber < minLocation {
				minLocation = currentNumber
			}
		}
	}
	// fmt.Println(seeds)
	fmt.Println(minLocation)
}
