package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveDay6() {
	fmt.Println()
	fmt.Println("Solving Day 6!")
	badKerning, properKerning, err := solveDay6("day6/input.txt")
	if err != nil {
		fmt.Println("Error while solving for the answer of day 6!", err)
		return
	}
	fmt.Println("Ways to win, bad kerning input:\t\t", badKerning)
	fmt.Println("Ways to win, proper kerning input:\t", properKerning)
	fmt.Println()
}

func solveDay6(filename string) (uint64, uint64, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return 0, 0, fmt.Errorf("Error while opening the file %v", err)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	times := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	distances := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	var timeToRecordDistanceBadKerning map[uint64]uint64 = make(map[uint64]uint64)
	for i := range times {
		time, err := strconv.ParseUint(times[i], 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("Error while parsing the input time into an integer %v", err)
		}
		distance, err := strconv.ParseUint(distances[i], 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("Error while parsing the input distance into an integer %v", err)
		}
		timeToRecordDistanceBadKerning[time] = distance
	}

	correctKerningTimeString := strings.Join(times, "")
	correctKerningDistanceString := strings.Join(distances, "")
	correctKerningTime, err := strconv.ParseUint(correctKerningTimeString, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Error while parsing the correct kerning time into an integer %v", err)
	}
	correctKerningDistance, err := strconv.ParseUint(correctKerningDistanceString, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Error while parsing the input distance into an integer %v", err)
	}
	var timeToRecordDistanceProperKerning map[uint64]uint64 = map[uint64]uint64{
		correctKerningTime: correctKerningDistance,
	}

	badKerning := findPossibleWinProduct(timeToRecordDistanceBadKerning)
	properKerning := findPossibleWinProduct(timeToRecordDistanceProperKerning)
	return badKerning, properKerning, nil
}

func findPossibleWinProduct(timeToRecordDistance map[uint64]uint64) uint64 {
	var waysWinProduct uint64 = 1
	for time, distanceToBeat := range timeToRecordDistance {
		var waysWinPossible uint64 = 0
		var i uint64 = 1
		for i = 1; i <= time; i++ {
			if i*(time-i) > distanceToBeat {
				waysWinPossible++
			}
		}
		waysWinProduct *= waysWinPossible
	}
	return waysWinProduct
}
