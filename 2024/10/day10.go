package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

type Trail struct {
	start Point
	end Point
}

var m map[Trail]bool
var m1 map[Trail]int

func rec(maze [][]int, x int, y int, ind int, xZero int, yZero int) {
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[x]) {
		return
	}
	if maze[x][y] != ind {
		return
	}
	if maze[x][y] == 9 {
		start := Point{xZero, yZero}
		end := Point{x, y}
		trail := Trail{start, end}
		m[trail] = true
		m1[trail] = m1[trail] + 1
		return
	}
	rec(maze, x - 1, y, ind + 1, xZero, yZero)
	rec(maze, x, y - 1, ind + 1, xZero, yZero)
	rec(maze, x, y + 1, ind + 1, xZero, yZero)
	rec(maze, x + 1, y, ind + 1, xZero, yZero)
}

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	m = make(map[Trail]bool)
	m1 = make(map[Trail]int)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maze [][]int
	var index int = 0
	for scanner.Scan() {
		var input string = scanner.Text()
		maze = append(maze, []int{})
		for _, el := range input {
			maze[index] = append(maze[index], int(el) - '0')
		}
		index++
	}
	for i, row := range maze {
		for j, el := range row {
			if el == 0 {
				rec(maze, i, j, 0, i, j)
			}
		}
	}
	var sum1 int = 0
	var sum2 int = 0
	for range m {
		sum1++
	}
	for _, value := range m1 {
		sum2 += value
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
