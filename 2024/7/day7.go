package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var xmas string = "XMAS"

func rec(nums []int, index int, curValue int) bool {
	if index == len(nums) && curValue == nums[0] {
		return true
	} 
	if curValue > nums[0] || index >= len(nums) {
		return false
	}
	if rec(nums, index + 1, curValue + nums[index]) {
		return true
	}
	if rec(nums, index + 1, curValue * nums[index]) {
		return true
	}
	if rec(nums, index + 1, concat(curValue, nums[index])) {
		return true
	}
	return false
}

func concat(x int, y int) int {
	var sb strings.Builder
	s1 := strconv.FormatInt(int64(x), 10)
	s2 := strconv.FormatInt(int64(y), 10)
	sb.WriteString(s1)
	sb.WriteString(s2)
	num, _ := strconv.Atoi(sb.String())
	// fmt.Println(num)
	return num
}

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums [][]int
	var index int = 0
	for scanner.Scan() {
		var input string = scanner.Text()
		numStr := strings.Split(input, " ")
		nums = append(nums, []int{})
		num, _ := strconv.Atoi(strings.TrimRight(numStr[0], ":"))
		nums[index] = append(nums[index], num)
		for i := 1; i < len(numStr); i++ {
			num, _ := strconv.Atoi(numStr[i])
			nums[index] = append(nums[index], num)
		}
		index++
	}
	var sum int64 = 0
	for _, el := range nums {
		if (rec(el, 2, el[1])) {
			sum += int64(el[0])
		}
	}
	fmt.Println(sum)
}
