package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var xmas string = "XMAS"

func rec(maze []string, x int, y int, dirX int, dirY int, letter int) bool {
	if letter >= 4 {
		return true
	}
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[x]) || maze[x][y] != xmas[letter] {
		return false
	}
	return rec(maze, x + dirX, y + dirY, dirX, dirY, letter + 1)
}

func circle(maze []string, x int, y int) int {
	sum := 0
	nums := []int{-1, -1, -1, 0, -1, 1, 0, -1, 0, 1, 1, -1, 1, 0, 1, 1}
	for i := 0; i < len(nums); i += 2 {
		if rec(maze, x, y, nums[i], nums[i + 1], 0) {
			sum++
		}
	}
	return sum
}

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maze []string
	for scanner.Scan() {
		var input string = scanner.Text()
		maze = append(maze, input)
	}
	var sum int = 0
	for i, el := range maze {
		if i == 0 || i == len(maze) - 1 {
			continue
		}
		for j, c := range el {
			if j == 0 || j == len(el) - 1 {
				continue
			}
			// if c == 'X' {
			// 	sum = sum + circle(maze, i, j)
			// }
			if c == 'A' {
				if maze[i-1][j-1] == 'M' && maze[i+1][j-1] == 'M' && maze[i-1][j+1] == 'S' && maze[i+1][j+1] == 'S' {
					sum++
				} else if maze[i-1][j-1] == 'S' && maze[i+1][j-1] == 'S' && maze[i-1][j+1] == 'M' && maze[i+1][j+1] == 'M' {
					sum++
				} else if maze[i-1][j-1] == 'M' && maze[i+1][j-1] == 'S' && maze[i-1][j+1] == 'M' && maze[i+1][j+1] == 'S' {
					sum++
				} else if maze[i-1][j-1] == 'S' && maze[i+1][j-1] == 'M' && maze[i-1][j+1] == 'S' && maze[i+1][j+1] == 'M' {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}
