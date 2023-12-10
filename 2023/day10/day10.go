package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findS(maze []string) (int, int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

// before
// up 1
// down 2
// left 3
// right 4
var count int = 0

func rec(maze []string, visited [][]bool, x, y, before int, insideLoop [][]int, starts int) bool {
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[x]) {
		return false
	}
	if maze[x][y] == 'S' && visited[x][y] {
		if (starts == 1 && before == 2) || (starts == 2 && before == 1) {
			insideLoop[x][y] = 4
		} else if (starts == 3 && before == 4) || (starts == 4 && before == 3) {
			insideLoop[x][y] = 3
		} else if (starts == 1 && before == 3) || (starts == 3 && before == 1) {
			insideLoop[x][y] = 2
		} else if (starts == 2 && before == 3) || (starts == 3 && before == 2) {
			insideLoop[x][y] = 5
		} else if (starts == 1 && before == 4) || (starts == 4 && before == 1) {
			insideLoop[x][y] = 1
		} else if (starts == 2 && before == 4) || (starts == 4 && before == 2) {
			insideLoop[x][y] = 6
		}
		return true
	}
	if visited[x][y] {
		return false
	}
	visited[x][y] = true
	switch before {
	case 1:
		if maze[x][y] == 'L' && rec(maze, visited, x, y+1, 3, insideLoop, starts) {
			insideLoop[x][y] = 1
			count++
			return true
		}
		if maze[x][y] == '|' && rec(maze, visited, x+1, y, 1, insideLoop, starts) {
			insideLoop[x][y] = 4
			count++
			return true
		}
		if maze[x][y] == 'J' && rec(maze, visited, x, y-1, 4, insideLoop, starts) {
			insideLoop[x][y] = 2
			count++
			return true
		}
	case 2:
		if maze[x][y] == '7' && rec(maze, visited, x, y-1, 4, insideLoop, starts) {
			insideLoop[x][y] = 5
			count++
			return true
		}
		if maze[x][y] == '|' && rec(maze, visited, x-1, y, 2, insideLoop, starts) {
			insideLoop[x][y] = 4
			count++
			return true
		}
		if maze[x][y] == 'F' && rec(maze, visited, x, y+1, 3, insideLoop, starts) {
			insideLoop[x][y] = 6
			count++
			return true
		}
	case 3:
		if maze[x][y] == '-' && rec(maze, visited, x, y+1, 3, insideLoop, starts) {
			insideLoop[x][y] = 3
			count++
			return true
		}
		if maze[x][y] == '7' && rec(maze, visited, x+1, y, 1, insideLoop, starts) {
			insideLoop[x][y] = 5
			count++
			return true
		}
		if maze[x][y] == 'J' && rec(maze, visited, x-1, y, 2, insideLoop, starts) {
			insideLoop[x][y] = 2
			count++
			return true
		}
	case 4:
		if maze[x][y] == '-' && rec(maze, visited, x, y-1, 4, insideLoop, starts) {
			insideLoop[x][y] = 3
			count++
			return true
		}
		if maze[x][y] == 'F' && rec(maze, visited, x+1, y, 1, insideLoop, starts) {
			insideLoop[x][y] = 6
			count++
			return true
		}
		if maze[x][y] == 'L' && rec(maze, visited, x-1, y, 2, insideLoop, starts) {
			insideLoop[x][y] = 1
			count++
			return true
		}
	case 0:
		if rec(maze, visited, x, y+1, 3, insideLoop, 4) {
			count++
			return true
		}
		if rec(maze, visited, x+1, y, 1, insideLoop, 2) {
			count++
			return true
		}
	}
	return false
}

func inflection(first int, rest []int) bool {
	if first != 1 && first != 6 {
		return false
	}
	for _, el := range rest {
		if el == 3 {
			continue
		}
		if first == 1 && el == 5 {
			return true
		}
		if first == 6 && el == 2 {
			return true
		}
		return false
	}
	return false
}

func main() {
	// file, err := os.Open("shortInput2.txt")
	file, err := os.Open("longInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maze []string
	var visited [][]bool
	var index int = 0
	var insideLoop [][]int
	for scanner.Scan() {
		insideLoop = append(insideLoop, []int{})
		visited = append(visited, []bool{})
		var input string = scanner.Text()
		maze = append(maze, input)
		for i := 0; i < len(input); i++ {
			insideLoop[index] = append(insideLoop[index], 0)
			visited[index] = append(visited[index], false)
		}
		index++
	}
	x, y := findS(maze)
	rec(maze, visited, x, y, 0, insideLoop, 0)
	var sum int = 0
	var inside bool = false
	for _, el := range insideLoop {
		for i := 0; i < len(el); i++ {
			if el[i] > 0 && el[i] < 7 && el[i] != 3 {
				if !inflection(el[i], el[i+1:]) {
					inside = !inside
				}
			}
			if el[i] == 0 && inside {
				sum++
			}
		}
	}
	fmt.Println("==== Part 1 ====")
	fmt.Println(count / 2)
	fmt.Println("==== Part 2 ====")
	fmt.Println(sum)
}
