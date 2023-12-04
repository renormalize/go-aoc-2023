package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveDay3() {
	fmt.Println()
	fmt.Println("Solving Day 3!")
	sumNumbers, gearRatioSum, err := sumNumbersSchematic("day3/input.txt")
	if err != nil {
		fmt.Println("Failed while trying to solve with error: ", err)
		fmt.Println()
	}
	fmt.Println("The sum of the engine schematic numbers is: ", sumNumbers)
	fmt.Println("The sum of the gear ratios is: ", gearRatioSum)
	fmt.Println()
}

func sumNumbersSchematic(filename string) (int, int, error) {
	var err error
	inputFile, err := os.Open(filename)
	defer inputFile.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("error opening the input file for day 3 with error: %v", err)
	}
	scanner := bufio.NewScanner(inputFile)
	b := scanner.Scan()
	if !b {
		return 0, 0, nil
	}
	str := scanner.Text()
	lines := []string{strings.Repeat(".", len(str)+3), "." + str + ".."}
	for {
		b = scanner.Scan()
		if !b {
			break
		}
		lines = append(lines, "."+scanner.Text()+"..")
	}
	lines = append(lines, strings.Repeat(".", len(str)+3))
	var sum int
	var gearRatioSum int
	for i := 1; i < len(lines)-1; i++ {
		var rolling string
		var foundValidNumber bool
		for j := 1; j < len(lines[i])-1; j++ {
			// proceed if a digit and keep adding
			if lines[i][j] == '*' {
				// check surroundings 2
				// if not 2, ignore
				// if just 2 find the numbers
				if checkSurroundingsForRatio(i, j, lines) == 2 {
					gearRatioSum += getGearRatio(i, j, lines)
				}
			}
			if isDigit(lines[i][j]) {
				rolling += string(lines[i][j])
				foundValidNumber = foundValidNumber || checkSurroundings(i, j, lines)
			} else {
				if foundValidNumber {
					number, err := strconv.Atoi(rolling)
					if err != nil {
						return sum, gearRatioSum, err
					}
					sum += number
				}
				rolling = ""
				foundValidNumber = false
			}
		}
	}
	return sum, gearRatioSum, nil
}

func getGearRatio(x, y int, lines []string) int {
	var gearRatio int = 1
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if isDigit(lines[x+i][y+j]) {
				var temp string = string(lines[x+i][y+j])
				var left, right int = y + j - 1, y + j + 1
				for isDigit(lines[x+i][left]) {
					temp = string(lines[x+i][left]) + temp
					left--
				}
				for isDigit(lines[x+i][right]) {
					temp = temp + string(lines[x+i][right])
					right++
				}
				ratioString, _ := strconv.Atoi(temp)
				gearRatio *= ratioString
			}
		}
	}
	return gearRatio
}

func checkSurroundingsForRatio(x, y int, lines []string) int {
	var hits int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if isDigit(lines[x+i][y+j]) {
				hits++
			}
		}
	}
	return hits
}

func isDigit(c byte) bool {
	if c < '0' || c > '9' {
		return false
	}
	return true
}

func checkSurroundings(x, y int, lines []string) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !isDigit(lines[x+i][y+j]) && lines[x+i][y+j] != '.' {
				return true
			}
		}
	}
	return false
}
