package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Printf("Comparing rucksacks\n")
	Prioritize()
}

func Prioritize() {
	var currentLine = 0
	var totalScore = 0
	var badges string = ""
	var group [3]string

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

	for scanner.Scan() {
		actualLine := scanner.Text()

		group[currentLine] = actualLine

		if (currentLine+1)%3 == 0 {
			badges += DetermineBadge(group)
			currentLine = 0
		} else {
			currentLine++
		}
	}

	// Calculate the priority
	totalScore = CalculatePriority(badges)
	fmt.Println(totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DetermineBadge(group [3]string) string {
	var badge string = ""

	firstPart := group[0]
	secondPart := group[1]
	thirdPart := group[2]

	for index, item := range firstPart {
		if strings.Contains(secondPart, string(item)) && strings.Contains(thirdPart, string(item)) {
			badge = string(item)
			break
		}
		index++
	}

	return badge
}

func CalculatePriority(badges string) int {
	var totalScore = 0
	for index, item := range badges {
		totalScore += strings.Index(priority, string(item)) + 1
		index++
	}

	return totalScore
}
