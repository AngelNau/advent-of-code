package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println("Heho")
	file, err := os.Open("longInput.txt")
	// file, err := os.Open("shortInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var times []string = re.FindAllString(scanner.Text(), -1)
	var timesInt []int
	for _, el := range times {
		num, _ := strconv.Atoi(el)
		timesInt = append(timesInt, num)
	}
	var time string = s.Replace(s.Split(scanner.Text(), ":")[1], " ", "", -1)
	timeInt, _ := strconv.ParseUint(time, 10, 64)
	scanner.Scan()
	var distances []string = re.FindAllString(scanner.Text(), -1)
	var distancesInt []int
	for _, el := range distances {
		num, _ := strconv.Atoi(el)
		distancesInt = append(distancesInt, num)
	}
	var distance string = s.Replace(s.Split(scanner.Text(), ":")[1], " ", "", -1)
	distanceInt, _ := strconv.ParseUint(distance, 10, 64)
	fmt.Println(timeInt)
	fmt.Println(distanceInt)
	var product int = 1
	for i := 0; i < len(timesInt); i++ {
		var sum int = 0
		for j := 1; j < timesInt[i]; j++ {
			if j*(timesInt[i]-j) > distancesInt[i] {
				sum += 1
			}
		}
		product = product * sum
	}
	fmt.Println(product)
	var sum int = 0
	for i := uint64(1); i < timeInt; i++ {
		if i*(timeInt-i) > distanceInt {
			sum++
		}
	}
	fmt.Println(sum)
}
