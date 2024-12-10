package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
// 0 2 4 6
// ^ > v < 
var nums = [8]int{-1, 0, 0, +1, 1, 0, 0, -1}
var sum int = 0

func rec(maze []string, x int, y int, dirX int, dirY int, ind int) {
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[x]) {
		return
	}
	if maze[x][y] == byte(ind + 30) {
		sum++
		return
	}
	if maze[x][y] == '.' || maze[x][y] == '^' {
		out := []rune(maze[x])
		out[y] = rune(byte(ind + 30))
		maze[x] = string(out)
	}
	if x + dirX >= 0 && x + dirX < len(maze) && y + dirY >= 0 && y + dirY < len(maze[x + dirX]) && maze[x + dirX][y + dirY] == '#' {
		ind = (ind + 2) % 8
		dirX = nums[ind]
		dirY = nums[ind+1]
	}
	rec(maze, x + dirX, y + dirY, dirX, dirY, ind)
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
	var found bool = false
	var ii int
	var jj int
	for i, el := range maze {
		for j, c := range el {
			if c == '^' {
				ii = i
				jj = j
				found = true
				break;
			}
		}
		if found {
			break
		}
	}
	for i, el := range maze {
		for j, c := range el {
			if c == '.' {
				out := []rune(maze[i])
				out[j] = '#'
				maze[i] = string(out)
				newMaze := make([]string, len(maze))
				copy(newMaze, maze)
				rec(newMaze, ii, jj, nums[0], nums[1], 0)
				out[j] = '.'
				maze[i] = string(out)
			}
		}
	}
	// var sum int = 0
	// for _, el := range maze {
	// 	for _, c := range el {
	// 		if c == 'X' || c == '^' {
	// 			sum++
	// 		}
	// 	}
	// }
	fmt.Println(sum)
}
