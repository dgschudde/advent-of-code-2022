package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Reading pairs.\n")
	Read()
}

func Read() {
	var totalOverlap = 0

	// Read the input from file
	inputFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	// Read the values per line from the input file
	for scanner.Scan() {
		actualLine := scanner.Text()
		pairs := strings.Split(actualLine, ",")

		if DetermineOverlap(pairs) {
			fmt.Printf("Found overlap for pair: %s\n", actualLine)
			totalOverlap++
		}
	}

	fmt.Printf("Total overlapping pairs: %d\n\n", totalOverlap)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DetermineOverlap(pairs []string) bool {

	minMaxElfOne := strings.Split(pairs[0], "-")
	minMaxEfTwo := strings.Split(pairs[1], "-")

	minElfOne, err := strconv.ParseInt(minMaxElfOne[0], 10, 64)
	maxElfOne, err := strconv.ParseInt(minMaxElfOne[1], 10, 64)

	minElfTwo, err := strconv.ParseInt(minMaxEfTwo[0], 10, 64)
	maxElfTwo, err := strconv.ParseInt(minMaxEfTwo[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return CheckOverlap(minElfOne, maxElfOne, minElfTwo, maxElfTwo) && CheckOverlap(minElfTwo, maxElfTwo, minElfOne, maxElfOne)
}

func CheckOverlap(min1 int64, max1 int64, min2 int64, max2 int64) bool {
	if min1 < min2 && max1 < min2 || min1 > max2 && min2 > max2 {
		return false
	}
	return true
}
