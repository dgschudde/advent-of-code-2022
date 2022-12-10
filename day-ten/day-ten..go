package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type programline struct {
	instruction string
	value       int
}

func main() {
	var program = Read()
	Execute(&program)
}

func Read() []programline {
	var program = make([]programline, 0)

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
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		var amount int64 = 0
		if len(splitLine) > 1 {
			amount, _ = strconv.ParseInt(splitLine[1], 10, 32)
		}

		instruction := programline{instruction: splitLine[0], value: int(amount)}

		program = append(program, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return program
}

func Execute(program *[]programline) {
	var cycle = 1
	var registerX = 1
	var totalSignalStrength = 0

	var refProgram = *program

	for _, line := range refProgram {
		if line.instruction == "addx" {
			cycle++
			totalSignalStrength += CheckSignalStrength(cycle, registerX)
			registerX += line.value
			cycle++
			totalSignalStrength += CheckSignalStrength(cycle, registerX)

		}

		if line.instruction == "noop" {
			cycle++
			totalSignalStrength += CheckSignalStrength(cycle, registerX)
		}
	}

	fmt.Printf("Total signal strength is %d\n", totalSignalStrength)
}

func CheckSignalStrength(cycle int, registerX int) int {
	var signalStrength = 0

	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		fmt.Printf("Signal strength at cycle: %d is %d\n", cycle, registerX)
		signalStrength += registerX * cycle
	}

	return signalStrength
}