package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input = Read()

	var result = CalculateVisibleTrees(&input)

	fmt.Printf("Total visible trees: %d\n", result)
}

func Read() [][]int64 {
	var input [][]int64
	var temp []int64
	var treeHeight int64

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

	// Create a 2d array
	input = make([][]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()

		temp = make([]int64, 0)

		for _, character := range line {
			treeHeight, _ = strconv.ParseInt(string(character), 10, 64)
			temp = append(temp, treeHeight)
		}

		input = append(input, [][]int64{temp}...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func CalculateVisibleTrees(input *[][]int64) int {
	refInput := *input

	height := len(refInput)
	width := len(refInput[0])

	visibleFromCorners := (height+width)*2 - 4

	innerTreesVisible := DiscoverInnerTrees(input)

	fmt.Printf("Inner trees visible: %d\n", innerTreesVisible)

	visibleTrees := visibleFromCorners + innerTreesVisible

	fmt.Printf("Visible from the corners: %d\n", visibleFromCorners)

	return visibleTrees
}

func DiscoverInnerTrees(input *[][]int64) int {
	refInput := *input
	var innerCount = 0

	var vertical = make([]int64, 0)

	height := len(refInput)
	width := len(refInput[0])

	for row := 1; row <= height-2; row++ {
		for column := 1; column < width-1; column++ {
			vertical = make([]int64, 0)
			for i := 0; i < height; i++ {
				vertical = append(vertical, refInput[row][i])
			}

			if CheckVisibility(refInput[row][column], refInput[row], vertical, column, row) {
				innerCount++
			}
		}
		fmt.Println()
	}

	return innerCount
}

func CheckVisibility(tree int64, horizontal []int64, vertical []int64, xPos int, yPos int) bool {
	fmt.Printf("Checking tree %d\n", tree)

	horizontalVisible := CheckHorizontal(tree, xPos, horizontal)
	verticalVisible := CheckVertical(tree, yPos, vertical)

	if horizontalVisible {
		fmt.Println("Horizontal visible")
	}

	if verticalVisible {
		fmt.Println("Vertical visible")
	}

	fmt.Println()

	return horizontalVisible || verticalVisible
}

func CheckHorizontal(tree int64, xPos int, horizontal []int64) bool {
	var forward = false
	var backward = false

	// Check forward
	for i := xPos + 1; i < len(horizontal); i++ {
		if tree-1 > horizontal[i] {
			forward = true
		} else {
			forward = false
			break
		}
	}

	// Check backward
	for i := xPos - 1; i >= 0; i-- {
		if tree-1 > horizontal[i] {
			backward = true
		} else {
			backward = false
			break
		}
	}

	return forward || backward
}

func CheckVertical(tree int64, yPos int, vertical []int64) bool {
	var forward = false
	var backward = false

	// Check forward
	for i := yPos + 1; i < len(vertical); i++ {
		if tree-1 > vertical[i] {
			forward = true
		} else {
			forward = false
			break
		}
	}

	// Check backward
	for i := yPos - 1; i >= 0; i-- {
		if tree-1 > vertical[i] {
			backward = true
		} else {
			backward = false
			break
		}
	}

	return forward || backward
}
