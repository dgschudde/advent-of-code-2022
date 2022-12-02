package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Playing Rock, Paper, Scissors\n")
	PlayGamesPartOne()
	PlayGamesPartTwo()
}

func PlayGamesPartOne() {
	var currentLine = 0
	var totalScore = 0
	var strategyGuide [2500][2]string

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
	fmt.Printf("Reading strategy guide\n")

	for scanner.Scan() {
		actualLine := scanner.Text()

		var split = strings.Split(actualLine, " ")
		strategyGuide[currentLine][0] = split[0]
		strategyGuide[currentLine][1] = split[1]

		currentLine++
	}
	fmt.Printf("%d lines read\n", currentLine)

	// Loop to the 2-d array
	for line, guide := range strategyGuide {
		currentGame := line + 1
		opponent := guide[0]
		me := guide[1]

		fmt.Printf("Game %d -> opponent %s - me %s\n", currentGame, opponent, me)
		currentScore := CalculateScore(opponent, me)
		fmt.Printf("Current game results in %d points\n\n", currentScore)
		totalScore += currentScore
	}

	fmt.Printf("You colected %d points", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func PlayGamesPartTwo() {
	var currentLine = 0
	var totalScore = 0
	var strategyGuide [2500][2]string
	var me string

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
	fmt.Printf("Reading strategy guide\n")

	for scanner.Scan() {
		actualLine := scanner.Text()

		var split = strings.Split(actualLine, " ")
		strategyGuide[currentLine][0] = split[0]
		strategyGuide[currentLine][1] = split[1]

		currentLine++
	}
	fmt.Printf("%d lines read\n", currentLine)

	// Loop to the 2-d array
	for line, guide := range strategyGuide {
		currentGame := line + 1
		opponent := guide[0]
		desirableResult := guide[1]

		// Select to correct shape for the desired result
		me = ChooseShape(opponent, desirableResult)

		fmt.Printf("Game %d -> opponent %s - me %s\n", currentGame, opponent, me)
		currentScore := CalculateScore(opponent, me)
		fmt.Printf("Current game results in %d points\n\n", currentScore)
		totalScore += currentScore
	}

	fmt.Printf("You colected %d points", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func CalculateScore(opponent string, me string) int {
	switch me {

	// Rock
	case "X":
		// 1 point
		return 1 + Outcome(opponent, me)

	//Paper
	case "Y":
		// 2 points
		return 2 + Outcome(opponent, me)

	// Scissors
	case "Z":
		// 3 points
		return 3 + Outcome(opponent, me)
	}

	return 0
}

func Outcome(opponent string, me string) int {
	// Wins
	if (opponent == "A" && me == "Y") || (opponent == "B" && me == "Z") || (opponent == "C" && me == "X") {
		return 6
	}

	// Draw
	if (opponent == "A" && me == "X") || (opponent == "B" && me == "Y") || (opponent == "C" && me == "Z") {
		return 3
	}

	// Else is loss
	return 0
}

func ChooseShape(opponent string, desirableResult string) string {
	var me = ""

	// Choose if I need to win (Z), draw (Y) or lose (X)
	switch desirableResult {
	// Draw
	case "Y":
		switch opponent {
		case "A":
			me = "X"
		case "B":
			me = "Y"
		case "C":
			me = "Z"
		}

	// Win
	case "Z":
		switch opponent {
		case "A":
			me = "Y"
		case "B":
			me = "Z"
		case "C":
			me = "X"
		}

	// Lose
	case "X":
		switch opponent {
		case "A":
			me = "Z"
		case "B":
			me = "X"
		case "C":
			me = "Y"
		}
	}

	return me
}
