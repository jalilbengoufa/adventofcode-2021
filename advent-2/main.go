package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//part 1  = 2150351
// func main() {
// 	lines, err := readData("dat")
// 	if err != nil {
// 		log.Fatalf("readData: %s", err)
// 	}
// 	var horizontal int
// 	var depth int
// 	for i, _ := range lines {
// 		s := strings.Split(lines[i], " ")
// 		command := s[0]

// 		value, err := strconv.Atoi(s[1])
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(2)
// 		}
// 		if command == "forward" {
// 			horizontal = horizontal + value

// 		} else if command == "up" {
// 			depth = depth - value

// 		} else if command == "down" {
// 			depth = depth + value

// 		}
// 	}
// 	fmt.Println(horizontal * depth)
// }

//part 2 = 1842742223
func main() {
	lines, err := readData("dat2")
	if err != nil {
		log.Fatalf("readData: %s", err)
	}
	var horizontal int
	var depth int
	aim := 0
	for i, _ := range lines {
		s := strings.Split(lines[i], " ")
		command := s[0]

		value, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if command == "forward" {
			horizontal = horizontal + value
			depth += aim * value

		} else if command == "up" {
			aim = aim - value

		} else if command == "down" {
			aim = aim + value

		}
	}
	fmt.Println(horizontal * depth)
}
