package day1

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SolveDay1() {
	fmt.Println("Solving Day 1!")
	fmt.Println("Solution for part 1 is: ", getDigitsAndSum())
	fmt.Println("Solution for part 2 is: ", wordDigitsPossible())
}

func getDigitsAndSum() (answer int) {
	inputFile, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening the input file for day 1 with error: ", err)
		return
	}
	for {
		var s string
		_, err := fmt.Fscanln(inputFile, &s)
		if err == io.EOF {
			return
		}
		// input as runes
		rs := []rune(s)
		lenRunes := len(rs)
		var i, j int = 0, lenRunes - 1
		// first digit
		for ; i < lenRunes; i++ {
			if rs[i] >= '0' && rs[i] <= '9' {
				break
			}
		}
		// last digit
		for ; j >= 0; j-- {
			if rs[j] >= '0' && rs[j] <= '9' {
				break
			}
		}
		answer += int((rs[i]-'0'))*10 + int((rs[j] - '0'))
	}
}

var wordToNumber map[string]int = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func wordDigitsPossible() (answer int) {
	inputFile, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening the input file for day 1 with error: ", err)
		return
	}
	for {
		var s string
		_, err := fmt.Fscanln(inputFile, &s)
		if err == io.EOF {
			return
		}
		// input as runes
		rs := []rune(s)
		lenRunes := len(rs)
		var i, j int = 0, lenRunes - 1
		// first digit
		for ; i < lenRunes; i++ {
			if rs[i] >= '0' && rs[i] <= '9' {
				break
			}
		}
		// last digit
		for ; j >= 0; j-- {
			if rs[j] >= '0' && rs[j] <= '9' {
				break
			}
		}
		// finding words which are numbers
		var minIndex, maxIndex int = len(s), -1
		var minString, maxString string
		for word := range wordToNumber {
			word := word
			// finding words which could be the leftmost "digit"
			if index := strings.Index(s, word); index != -1 {
				if index < minIndex {
					minIndex = index
					minString = word
				}
			}
			// finding words which could be the rightmost "digit"
			// last occurence of a word in a string is the first
			// occurence of the reverse word in the reversed string
			if index := strings.Index(Reverse(s), Reverse(word)); index != -1 {
				index = len(s) - index - 1
				if index > maxIndex {
					maxIndex = index
					maxString = word
				}
			}
		}
		// first digit: character which was a digit occurs before the word "digit"
		if i < minIndex {
			answer += int((rs[i] - '0')) * 10
		} else {
			answer += wordToNumber[minString] * 10
		}
		// second digit:  character which was a digit occurs after the word "digit"
		if j > maxIndex {
			answer += int((rs[j] - '0'))
		} else {
			answer += wordToNumber[maxString]
		}
	}
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
