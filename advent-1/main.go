package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readData(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}

//part 1  answer: 1387
// func main() {
// 	lines, err := readData("dat")
// 	if err != nil {
// 		log.Fatalf("readData: %s", err)
// 	}
// 	var total int
// 	for i, _ := range lines {
// 		if i != 0 {
// 			if lines[i] > lines[i-1] {
// 				total++
// 			}

// 		}
// 	}
// 	fmt.Println(total)
// }

//part 2 answer:1362
func main() {
	lines, err := readData("dat")
	if err != nil {
		log.Fatalf("readData: %s", err)
	}

	var total int
	for i, _ := range lines {
		if i < len(lines)-3 {
			j := i + 1
			sum1 := lines[i] + lines[i+1] + lines[i+2]
			sum2 := lines[j] + lines[j+1] + lines[j+2]
			if sum2 > sum1 {
				total++
			}
		}

	}

	fmt.Println(total)
}
