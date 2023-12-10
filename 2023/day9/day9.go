package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func diff(array []int) int {
	var nextArray []int
	var zeros bool = true
	var lastNum int
	for i := 1; i < len(array); i++ {
		var num int = array[i] - array[i-1]
		nextArray = append(nextArray, num)
		if num != 0 {
			zeros = false
		}
		lastNum = array[i]
	}
	if zeros {
		return lastNum
	}
	return lastNum + diff(nextArray)
}

func diffBeginning(array []int) int {
	var nextArray []int
	var zeros bool = true
	for i := 1; i < len(array); i++ {
		var num int = array[i] - array[i-1]
		nextArray = append(nextArray, num)
		if num != 0 {
			zeros = false
		}
	}
	if zeros {
		return array[0]
	}
	return array[0] - diffBeginning(nextArray)
}

func main() {
	file, err := os.Open("longInput.txt")
	// file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum int = 0
	var sumBackwards int = 0
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var array []int
		var input string = scanner.Text()
		for _, el := range strings.Split(input, " ") {
			num, _ := strconv.Atoi(el)
			array = append(array, num)
		}
		sum += diff(array)
		sumBackwards += diffBeginning(array)
	}
	fmt.Println("==== Part 1 ====")
	fmt.Println(sum)
	fmt.Println("==== Part 2 ====")
	fmt.Println(sumBackwards)
}
