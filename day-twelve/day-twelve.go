package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var positionGrid = Read()
	Move(&positionGrid)
	PrintGrid(&positionGrid)
}

const Rows = 5
const Columns = 8

type Position struct {
	value     string
	visited   bool
	direction string
}

func Read() [Rows][Columns]Position {
	var positionGrid [Rows][Columns]Position

	// Read the input from file
	inputFile, err := os.Open("input/test-input.txt")

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

	var currentLine = 0

	for scanner.Scan() {
		line := scanner.Text()

		var currentCharacter = 0
		for _, character := range line {
			positionGrid[currentLine][currentCharacter] = Position{value: string(character), visited: false, direction: "."}
			currentCharacter++
		}

		currentLine++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return positionGrid
}

func PrintGrid(positionGrid *[Rows][Columns]Position) {
	refPositions := positionGrid

	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			fmt.Print(refPositions[i][j])
		}
		fmt.Println()
	}
}

func Move(positionGrid *[Rows][Columns]Position) {
	refPositions := positionGrid

	var currentColumn = 0
	var currentRow = 0

	// Start by moving one down
	var start = refPositions[currentRow][currentColumn]
	Search(refPositions, 1, 0)
	fmt.Println(start)

}

func Search(positionGrid *[Rows][Columns]Position, row int, column int) {
	var up = ""
	var down = ""
	var left = ""
	var right = ""

	var currentCharacter = positionGrid[row][column].value[0]

	for string(currentCharacter) != "E" {

		// Up
		if row-1 > 0 {
			up = positionGrid[column][row-1].value
		} else {
			up = "X"
		}

		// Down
		if row+1 < Rows {
			down = positionGrid[column][row+1].value
		} else {
			down = "X"
		}

		// Left
		if column-1 > 0 {
			left = positionGrid[column-1][row].value
		} else {
			left = "X"
		}

		// Right
		if column+1 < Columns {
			right = positionGrid[column+1][row].value
		} else {
			right = "X"
		}

		if up[0]-currentCharacter == 1 {
			positionGrid[column][row].direction = "^"
			positionGrid[column][row].visited = true
			row--
			Search(positionGrid, row, column)
			continue
		}

		if down[0]-currentCharacter == 1 {
			positionGrid[column][row].direction = "v"
			positionGrid[column][row].visited = true
			row++
			Search(positionGrid, row, column)
			continue
		}

		if left[0]-currentCharacter == 1 {
			positionGrid[column][row].direction = "<"
			positionGrid[column][row].visited = true
			column--
			Search(positionGrid, row, column)
			continue
		}

		if right[0]-currentCharacter == 1 {
			positionGrid[column][row].direction = ">"
			positionGrid[column][row].visited = true
			column++
			Search(positionGrid, row, column)
			continue
		}
	}
}
