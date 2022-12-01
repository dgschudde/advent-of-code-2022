package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("Calculating calories.\n")
	CalculateCalories()
}

func CalculateCalories() {
	var currentElf int64 = 1
	elves := make([]int64, 1)

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
		fmt.Printf("Counting calories for elf: %d\n", currentElf)
		actualLine := scanner.Text()

		if actualLine == "" {
			// When a new line is detected, start counting for a new elf
			currentElf++

			// Append a new item to the elves array
			elves = append(elves, 0)
		} else {
			// Convert the actual value to a int64, it's read as string
			actualValue, err := strconv.ParseInt(actualLine, 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			elves[currentElf-1] += actualValue
		}
	}

	// Sort the array with elves to determine the maximum value
	sort.Slice(elves, func(i, j int) bool { return elves[i] < elves[j] })
	fmt.Printf("Elf with the largest amount of calories %d\n", elves[len(elves)-1])

	// Calculate the calories for the top 3 elves
	topThreeTotal := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Printf("Amount of calories carried by the top 3 elves: %d\n", topThreeTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
