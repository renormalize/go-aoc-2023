package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveDay4() {
	fmt.Println()
	fmt.Println("Solving Day 4!")
	points, points2, err := getPoints1("day4/input.txt")
	if err != nil {
		fmt.Println("Failed while trying to solve with error: ", err)
		fmt.Println()
	}
	fmt.Println("The total points 1 are: ", points)
	fmt.Println("The total points 2 are: ", points2)
	fmt.Println()
}

func getPoints1(filename string) (int, int, error) {
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("error opening the input file for day 4 with error: %v", err)
	}
	scanner := bufio.NewScanner(inputFile)
	totalCardMap := make(map[int]int)
	for i := 1; i <= 213; i++ {
		totalCardMap[i] = 1
	}
	card := 1
	var points int
	for {
		b := scanner.Scan()
		if !b {
			break
		}
		line := scanner.Text()
		lineIndex := strings.Index(line, ":")
		line = line[lineIndex+1:]
		lines := strings.Split(line, "|")
		flags := strings.Split(lines[0], " ")
		var flagMap map[int]bool = make(map[int]bool)
		for _, stringFlag := range flags {
			convVal, _ := strconv.Atoi(stringFlag)
			if convVal == 0 {
				continue
			}
			flagMap[convVal] = true
		}
		vals := strings.Split(lines[1], " ")
		var hits int
		for _, st := range vals {
			st := strings.TrimSpace(st)
			value, _ := strconv.Atoi(st)
			if value == 0 {
				continue
			}
			if flagMap[value] == true {
				hits++
			}
		}
		if hits != 0 {
			points += IntPow(2, hits-1)
			for h := hits; h >= 1; h-- {
				totalCardMap[card+h] = totalCardMap[card+h] + totalCardMap[card]
			}
		}
		card++
	}
	var finSum int = 0
	for _, vv := range totalCardMap {
		finSum += vv
	}
	return points, finSum, nil
}

func IntPow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}
	return result
}
