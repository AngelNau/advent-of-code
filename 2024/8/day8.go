package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
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
	var maze []string
	var dupes1 []string
	var dupes2 []string
	var m map[rune][]int = make(map[rune][]int)
	for scanner.Scan() {
		var input string = scanner.Text()
		maze = append(maze, input)
		dupes1 = append(dupes1, input)
		dupes2 = append(dupes2, input)
	}
	for i, el := range maze {
		for j, char := range el {
			if (char >= 48 && char <= 57) || (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
				m[char] = append(m[char], i)
				m[char] = append(m[char], j)
			}
		}
	}
	for i := range maze {
		for j := range maze[i] {
			for _, val := range m {
				for ii := 0; ii < len(val); ii+=2 {
					for jj := 0; jj < len(val); jj += 2 {
						if ii == jj {
							continue
						}
						d1 := abs(i - val[ii]) + abs(j - val[ii + 1])
						d2 := abs(i - val[jj]) + abs(j - val[jj + 1])
						di1 := i - val[ii]
						di2 := i - val[jj]
						dj1 := j - val[ii + 1]
						dj2 := j - val[jj + 1]
						if (di1*dj2 == dj1*di2) {
							if d1 == 2*d2 || d2 == d1*2 {
								out := []rune(dupes1[i])
								out[j] = '#'
								dupes1[i] = string(out)
							}
							out := []rune(dupes2[i])
							out[j] = '#'
							dupes2[i] = string(out)
						}
					}
				}
			}
		}
	}
	var sum1 int = 0
	var sum2 int = 0
	for _, el := range dupes1 {
		for _, char := range el {
			if char == '#' {
				sum1 ++
			}
		}
	}
	for _, el := range dupes2 {
		for _, char := range el {
			if char == '#' {
				sum2 ++
			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
