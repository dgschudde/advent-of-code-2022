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
	var rucksacks [300]string
	var commonTypes string = ""

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
	fmt.Printf("Reading rucksacks contents\n")

	for scanner.Scan() {
		actualLine := scanner.Text()
		rucksacks[currentLine] = actualLine

		// Split the actual line
		var actualLength = len(actualLine)

		var half = actualLength / 2

		var firstPart = string(actualLine[0:half])
		var secondPart = string(actualLine[half:actualLength])

		fmt.Println(actualLine)
		fmt.Printf("%s\n", firstPart)
		fmt.Printf("%s\n", secondPart)

		// Intersect the both results
		commonTypes += CompareInventories(firstPart, secondPart)
		fmt.Printf("Types in common for rucksack %d: %s\n\n", currentLine+1, commonTypes)

		currentLine++
	}

	fmt.Printf("%d lines read\n\n", currentLine)

	// Calculate the priority
	totalScore = CalculatePriority(commonTypes)
	fmt.Printf("Total priority points: %d\n", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func CompareInventories(firstPart string, secondPart string) string {
	var commonTypes string = ""

	for index, item := range firstPart {
		if strings.Contains(secondPart, string(item)) {
			if !strings.Contains(commonTypes, string(item)) {
				commonTypes += string(item)
			}
		}
		index++
	}

	return commonTypes
}

func CalculatePriority(commonTypes string) int {
	var totalScore = 0
	for index, item := range commonTypes {
		totalScore += strings.Index(priority, string(item)) + 1
		index++
	}

	return totalScore
}
