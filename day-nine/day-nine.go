package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	steps     int
}

type position struct {
	visited bool
}

func main() {
	var input = Read()
	Move(&input)
}

func Read() []move {
	var moves = make([]move, 0)

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

	var total = 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		amountSteps, _ := strconv.ParseInt(splitLine[1], 10, 32)
		move := move{direction: splitLine[0], steps: int(amountSteps)}

		moves = append(moves, move)
	}

	fmt.Printf("total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves
}

func Move(input *[]move) {
	var refMoves = *input
	var field = [40000][40000]position{}

	var xPos = 0
	var yPos = 0

	var lines = 0
	var previousDirection = ""

	for _, move := range refMoves {
		lines++
		fmt.Println(lines)
		switch move.direction {
		case "R":
			if previousDirection == "U" || previousDirection == "D" {
				field[xPos+10000-1][yPos+10000] = position{visited: false}
			}

			for i := 0; i < move.steps; i++ {
				xPos++

				if i != move.steps-1 {
					field[xPos+10000][yPos+10000] = position{visited: true}
					previousDirection = move.direction
				}
			}
		case "L":
			if previousDirection == "U" || previousDirection == "D" {
				field[xPos+10000+1][yPos+10000] = position{visited: false}
			}

			for i := 0; i < move.steps; i++ {
				xPos--
				if i != move.steps-1 {
					field[xPos+10000][yPos+10000] = position{visited: true}
					previousDirection = move.direction
				}
			}
		case "U":
			if previousDirection == "L" || previousDirection == "R" {
				field[xPos+10000][yPos+10000+1] = position{visited: false}
			}

			for i := 0; i < move.steps; i++ {
				yPos++
				if i != move.steps-1 {
					field[xPos+10000][yPos+10000] = position{visited: true}
					previousDirection = move.direction
				}
			}
		case "D":
			if previousDirection == "L" || previousDirection == "R" {
				field[xPos+10000][yPos+10000-1] = position{visited: false}
			}

			for i := 0; i < move.steps; i++ {
				yPos--
				if i != move.steps-1 {
					field[xPos+10000][yPos+10000] = position{visited: true}
					previousDirection = move.direction
				}
			}
		}
	}

	var count = 0
	for i := 0; i < len(field[0]); i++ {
		for y := 0; y < len(field); y++ {
			if field[y][i].visited {
				count++
			}
		}
	}
	fmt.Println(count)
}
