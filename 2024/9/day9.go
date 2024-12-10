package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func numR(nums []int, i, num int) int {
	var n int = 0
	for i >= 0 && nums[i] == num {
		n++
		i--
	}
	return n
}

func numL(nums []int, i, num int) int {
	var n int = 0
	for i < len(nums) && nums[i] == num {
		n++
		i++
	}
	return n
}

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums []int
	scanner.Scan()
	var input string = scanner.Text()
	var index int = 0
	for i, char := range input {
		if i%2 == 0 {
			for num := 0; num < int(char)-'0'; num++ {
				nums = append(nums, index)
			}
			index++
		} else {
			for num := 0; num < int(char)-'0'; num++ {
				nums = append(nums, -1)
			}
		}
	}
	var i int = 0
	var j int = len(nums) - 1
	// for j := len(nums) - 1; j >= i; j-- {
	for {
		if j < 0 {
			break
		}
		for nums[j] == -1 {
			j--
		}
		for nums[i] != -1 {
			i++
		}
		var numR = numR(nums, j, nums[j])
		var numL = numL(nums, i, -1)
		if i >= j {
			j -= numR
			i = 0
			continue
		}
		if numR > numL {
			i = i + numL
			continue
		}
		for numR > 0 {
			nums[i] = nums[j]
			nums[j] = -1
			i++
			j--
			numR--
		}
		i = 0
	}
	
	// fmt.Println(nums)
	var sum uint64 = 0
	for i, el := range nums {
		if el == -1 {
			continue
		}
		sum += uint64(i * el)
	}
	fmt.Println(sum)
}
