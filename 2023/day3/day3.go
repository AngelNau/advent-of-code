package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func checkForSymbol(start int, end int, indexes []int) bool {
	if len(indexes) == 0 {
		return false
	}
	for i := start - 1; i <= end+1; i++ {
		if slices.Contains(indexes, i) {
			return true
		}
	}
	return false
}

func checkForNums(index int, indexes []int) (int, int, int) {
	if len(indexes) == 0 {
		return 0, -1, -1
	}
	var found int = 0
	var first int = -1
	var second int = -1
	for i := 0; i < len(indexes); i += 2 {
		if index >= indexes[i]-1 && index <= indexes[i+1]+1 {
			if found == 0 {
				first = i
				found++
			} else {
				second = i
				found++
				break
			}
		}
	}
	return found, first, second
}

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var entireInput []string
	var indexNums [][]int
	var indexSymbols [][]int
	var index int = 0
	for scanner.Scan() {
		var numFound bool = false
		var input string = scanner.Text()
		indexNums = append(indexNums, []int{})
		indexSymbols = append(indexSymbols, []int{})
		entireInput = append(entireInput, input)
		for i, el := range input {
			if !numFound && el >= 48 && el <= 57 {
				numFound = true
				indexNums[index] = append(indexNums[index], i)
			}
			if numFound && (el < 48 || el > 57) {
				numFound = false
				indexNums[index] = append(indexNums[index], i-1)
			}
			if numFound && i == len(input)-1 {
				indexNums[index] = append(indexNums[index], i)
			}
			// if (el < 48 || el > 57) && el != '.' {
			// 	indexSymbols[index] = append(indexSymbols[index], i)
			// }
			if el == '*' {
				indexSymbols[index] = append(indexSymbols[index], i)
			}
		}
		index++
	}
	var sum int = 0
	// for i, el := range indexNums {
	// 	for iter := 0; iter < len(el); iter += 2 {
	// 		var first int = el[iter]
	// 		var second int = el[iter+1]
	// 		if i > 0 {
	// 			if checkForSymbol(first, second, indexSymbols[i-1]) {
	// 				add, _ := strconv.Atoi(entireInput[i][first : second+1])
	// 				sum += add
	// 				continue
	// 			}
	// 		}
	// 		if i+1 < len(indexNums) {
	// 			if checkForSymbol(first, second, indexSymbols[i+1]) {
	// 				add, _ := strconv.Atoi(entireInput[i][first : second+1])
	// 				sum += add
	// 				continue
	// 			}
	// 		}
	// 		if checkForSymbol(first, second, indexSymbols[i]) {
	// 			add, _ := strconv.Atoi(entireInput[i][first : second+1])
	// 			sum += add
	// 		}
	// 	}
	// }
	for i, el := range indexSymbols {
		for iter := 0; iter < len(el); iter++ {
			var product int = 1
			var numOfNums int = 0
			if i > 0 {
				numOfAdj, first, second := checkForNums(el[iter], indexNums[i-1])
				numOfNums += numOfAdj
				fmt.Printf("first1: %d, second1: %d\n", first, second)
				switch numOfAdj {
				case 1:
					multiply, _ := strconv.Atoi(entireInput[i-1][indexNums[i-1][first] : indexNums[i-1][first+1]+1])
					fmt.Println(entireInput[i-1][indexNums[i-1][first]:indexNums[i-1][first+1]])
					product = product * multiply
				case 2:
					multiplyF, _ := strconv.Atoi(entireInput[i-1][indexNums[i-1][first] : indexNums[i-1][first+1]+1])
					multiplyS, _ := strconv.Atoi(entireInput[i-1][indexNums[i-1][second] : indexNums[i-1][second+1]+1])
					product = product * multiplyF * multiplyS
				}
			}
			if i+1 < len(indexNums) {
				numOfAdj, first, second := checkForNums(el[iter], indexNums[i+1])
				numOfNums += numOfAdj
				fmt.Printf("first2: %d, second2: %d\n", first, second)
				switch numOfAdj {
				case 1:
					multiply, _ := strconv.Atoi(entireInput[i+1][indexNums[i+1][first] : indexNums[i+1][first+1]+1])
					fmt.Println(entireInput[i+1][indexNums[i+1][first]:indexNums[i+1][first+1]])
					product = product * multiply
				case 2:
					multiplyF, _ := strconv.Atoi(entireInput[i+1][indexNums[i+1][first] : indexNums[i+1][first+1]+1])
					multiplyS, _ := strconv.Atoi(entireInput[i+1][indexNums[i+1][second] : indexNums[i+1][second+1]+1])
					product = product * multiplyF * multiplyS
				}
			}
			numOfAdj, first, second := checkForNums(el[iter], indexNums[i])
			numOfNums += numOfAdj
			fmt.Printf("first3: %d, second3: %d\n", first, second)
			switch numOfAdj {
			case 1:
				multiply, _ := strconv.Atoi(entireInput[i][indexNums[i][first] : indexNums[i][first+1]+1])
				product = product * multiply
			case 2:
				multiplyF, _ := strconv.Atoi(entireInput[i][indexNums[i][first] : indexNums[i][first+1]+1])
				multiplyS, _ := strconv.Atoi(entireInput[i][indexNums[i][second] : indexNums[i][second+1]+1])
				product = product * multiplyF * multiplyS
			}
			if numOfNums == 2 {
				fmt.Printf("i: %d, el[iter]: %d\n", i, el[iter])
				sum += product
			}
		}
	}
	fmt.Println(sum)
}
