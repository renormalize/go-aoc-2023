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
	points, scratchCards, err := solveScratchCards("day4/input.txt")
	if err != nil {
		fmt.Println("Failed while trying to solve with error: ", err)
		fmt.Println()
		return
	}
	fmt.Println("The total points are:\t\t\t", points)
	fmt.Println("The total total scratch cards are:\t", scratchCards)
	fmt.Println()
}

func solveScratchCards(filename string) (int, int, error) {
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("error opening the input file for day 4 with error: %v", err)
	}
	scanner := bufio.NewScanner(inputFile)
	totalCardMap := make(map[int]int)
	card := 1
	var points int
	for {
		b := scanner.Scan()
		if !b {
			break
		}
		if _, ok := totalCardMap[card]; !ok {
			totalCardMap[card] = 1
		}
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+1:]
		cardNumbers := strings.Split(line, "|")
		winningCards := strings.Split(cardNumbers[0], " ")
		var winningCardMap map[int]bool = make(map[int]bool)
		for _, stringFlag := range winningCards {
			winningCardNumber, _ := strconv.Atoi(stringFlag)
			if winningCardNumber == 0 {
				continue
			}
			winningCardMap[winningCardNumber] = true
		}
		numbers := strings.Split(cardNumbers[1], " ")
		var wins int
		for _, number := range numbers {
			number := strings.TrimSpace(number)
			value, _ := strconv.Atoi(number)
			if value == 0 {
				continue
			}
			if winningCardMap[value] == true {
				wins++
			}
		}
		if wins != 0 {
			points += intPow(2, wins-1)
			for i := wins; i >= 1; i-- {
				if _, ok := totalCardMap[card+i]; !ok {
					totalCardMap[card+i] = 1
				}
				// card + i guaranteed to be within the total number of cards
				totalCardMap[card+i] += totalCardMap[card]
			}
		}
		card++
	}
	var totalScratchCards int = 0
	for _, numberCards := range totalCardMap {
		totalScratchCards += numberCards
	}
	return points, totalScratchCards, nil
}

func intPow(base, exp int) int {
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
