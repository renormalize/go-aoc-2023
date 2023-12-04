package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveDay2() {
	fmt.Println()
	fmt.Println("Solving Day 2!")
	sumIDs, sumPowers, err := solveCubeConundrum("day2/input.txt")
	if err != nil {
		fmt.Println("Failed while trying to solve with error ", err)
		fmt.Println()
		return
	}
	fmt.Println("The sum of the Game IDs is:\t\t", sumIDs)
	fmt.Println("The sum of the powers of the sets is:\t", sumPowers)
	fmt.Println()
}

func solveCubeConundrum(filename string) (int, int, error) {
	var sumIDs, sumPowers int
	var err error
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		fmt.Println("Error opening the input file for day 2 with error: ", err)
		return sumIDs, sumPowers, err
	}
	var colorToMax map[string]int = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	scanner := bufio.NewScanner(inputFile)
	for {
		var minForColor map[string]int = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		b := scanner.Scan()
		if !b {
			return sumIDs, sumPowers, err
		}
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		gameID, err := strconv.Atoi(strings.Split(line[:colonIndex], " ")[1])
		if err != nil {
			return sumIDs, sumPowers, err
		}
		line = line[colonIndex+1:]
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ";", "")
		words := strings.Split(line, " ")[1:]
		var flag bool = true
		for i := 0; i < len(words)-1; i += 2 {
			// word[i] contains number of balls, word[i+1] contains color
			balls, _ := strconv.Atoi(words[i])
			if colorToMax[words[i+1]] < balls {
				flag = false
			}
			if minForColor[words[i+1]] < balls {
				minForColor[words[i+1]] = balls
			}
		}
		if flag {
			sumIDs += gameID
		}
		sumPowers += minForColor["red"] * minForColor["green"] * minForColor["blue"]
	}
}
