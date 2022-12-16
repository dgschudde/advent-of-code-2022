package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	xPos int64
	yPos int64
}

type Structure struct {
	position Position
	value    string
}

func main() {
	var cave [1000][1000]Structure
	Read(&cave)

	Search(&cave, Position{500, 0})

	PrintCave(&cave, Position{490, 0}, Position{510, 9})
	//PrintCave(&cave, Position{xPos: 460, yPos: 0}, Position{xPos: 540, yPos: 150})
}

func Read(cave *[1000][1000]Structure) {
	refCave := cave

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			refCave[y][x] = Structure{position: Position{int64(x), int64(y)}, value: "."}
		}
	}

	refCave[500][0] = Structure{position: Position{int64(500), int64(0)}, value: "+"}

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

	var currentPosition = Position{xPos: 0, yPos: 0}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "->")
		startPosition := splitLine[0]
		rockStart := strings.Split(startPosition, ",")
		startXPos, _ := strconv.ParseInt(strings.Trim(rockStart[0], " "), 10, 32)
		startYPos, _ := strconv.ParseInt(strings.Trim(rockStart[1], " "), 10, 32)
		currentPosition = Position{xPos: startXPos, yPos: startYPos}
		part := Structure{position: Position{startXPos, startYPos}, value: "#"}
		refCave[startXPos][startYPos] = part
		for i := 1; i < len(splitLine); i++ {
			position := splitLine[i]
			positionSplit := strings.Split(position, ",")
			xPos, _ := strconv.ParseInt(strings.Trim(positionSplit[0], " "), 10, 32)
			yPos, _ := strconv.ParseInt(strings.Trim(positionSplit[1], " "), 10, 32)
			part = Structure{position: Position{xPos, yPos}, value: "#"}

			refCave[xPos][yPos] = part

			if xPos != currentPosition.xPos {
				if xPos > currentPosition.xPos {
					for i := xPos; i > currentPosition.xPos; i-- {
						part = Structure{position: Position{i, yPos}, value: "#"}
						refCave[i][yPos] = part
					}
				} else {
					for i := xPos; i <= currentPosition.xPos; i++ {
						part = Structure{position: Position{i, yPos}, value: "#"}
						refCave[i][yPos] = part
					}
				}
			}

			if yPos != currentPosition.yPos {
				if yPos > currentPosition.yPos {
					for i := yPos; i > currentPosition.yPos; i-- {
						part = Structure{position: Position{xPos, i}, value: "#"}
						refCave[xPos][i] = part
					}
				} else {
					for i := yPos; i <= currentPosition.yPos; i++ {
						part = Structure{position: Position{xPos, i}, value: "#"}
						refCave[xPos][i] = part
					}
				}
			}

			currentPosition = Position{xPos, yPos}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var left = true

func Search(cave *[1000][1000]Structure, position Position) {
	refCave := cave

	xPos := position.xPos
	yPos := position.yPos

	if xPos < 0 || xPos >= 1000 || yPos < 0 || yPos >= 1000 {
		return
	}

	currentChar := cave[xPos][yPos].value

	switch currentChar {
	case ".":
		yPos++
		Search(refCave, Position{xPos, yPos})
	case "+":
		yPos++
		Search(refCave, Position{xPos, yPos})
	case "o":
		// Check Left
		if left {
			currentChar = refCave[xPos-1][yPos].value
			if currentChar == "." {
				refCave[xPos-1][yPos] = Structure{Position{xPos, yPos - 1}, "o"}
			}
		} else {
			currentChar = refCave[xPos+1][yPos].value
			if currentChar == "." {
				refCave[xPos-1][yPos] = Structure{Position{xPos, yPos - 1}, "o"}
			}
		}

		left = false
		Search(refCave, Position{500, 0})
	case "#":
		refCave[xPos][yPos-1] = Structure{Position{xPos, yPos - 1}, "o"}
		Search(refCave, Position{500, 0})
	}
}

func PrintCave(cave *[1000][1000]Structure, startPos Position, endPos Position) {
	refCave := cave

	startXPos := startPos.xPos
	startYPos := startPos.yPos

	endXPos := endPos.xPos
	endYPos := endPos.yPos

	for y := startYPos; y <= endYPos; y++ {
		for x := startXPos; x <= endXPos; x++ {
			fmt.Printf("%s", refCave[x][y].value)
		}
		fmt.Println()
	}
	fmt.Println()
}
