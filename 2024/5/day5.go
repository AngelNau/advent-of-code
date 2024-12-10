package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var orderings [][]int
	var index int = 0
	for scanner.Scan() {
		var input string = scanner.Text()
		if input == "" {
			break
		}
		numStr := strings.Split(input, "|")
		orderings = append(orderings, []int{})
		for _, el := range numStr {
			num, _ := strconv.Atoi(el)
			orderings[index] = append(orderings[index], num)
		}
		index++
	}
	var nums [][]int
	index = 0
	for scanner.Scan() {
		var input string = scanner.Text()
		numStr := strings.Split(input, ",")
		nums = append(nums, []int{})
		for _, el := range numStr {
			num, _ := strconv.Atoi(el)
			nums[index] = append(nums[index], num)
		}
		index++
	}
	orderedNums := make([][]int, len(nums))
	for i := range orderedNums {
		orderedNums[i] = make([]int, len(nums[i]))
		copy(orderedNums[i], nums[i])
	}
	for _, el := range orderedNums {
		sort.Slice(el, func(i, j int) bool {
			for _, sortEl := range orderings {
				if sortEl[0] == el[i] && sortEl[1] == el[j] {
					return true
				}
			}
			return false
		})
	}
	var sum int = 0
	for i := range orderedNums {
		is := false
		for j, el := range orderedNums[i] {
			if el != nums[i][j] {
				is = true
				break
			}
		}
		if is {
			var ind int = len(orderedNums[i]) / 2
			sum += orderedNums[i][ind]
		}
	}
	fmt.Println(sum)
}
