package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type programLine struct {
	instruction string
	value       int
}

func main() {
	var program = Read()
	Execute(&program)

	for _, character := range output {
		fmt.Printf("%s", string(character))
	}
}

var output [241]string

func Read() []programLine {
	var program = make([]programLine, 0)

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

		instruction := programLine{instruction: splitLine[0], value: int(amount)}

		program = append(program, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return program
}

func Execute(program *[]programLine) {
	var cycle = 1
	var registerX = 1

	var refProgram = *program

	for _, line := range refProgram {
		if line.instruction == "addx" {
			Update(cycle, registerX)

			cycle++
			Update(cycle, registerX)

			cycle++
			registerX += line.value
			Update(cycle, registerX)
		}

		if line.instruction == "noop" {
			Update(cycle, registerX)

			cycle++
			Update(cycle, registerX)
		}
	}
}

func Update(cycle int, registerX int) {
	if cycle > 0 && cycle <= 40 {
		Draw(cycle, registerX, 0)
	} else if cycle > 40 && cycle <= 80 {
		Draw(cycle, registerX, 40)
	} else if cycle > 80 && cycle <= 120 {
		Draw(cycle, registerX, 80)
	} else if cycle > 120 && cycle <= 160 {
		Draw(cycle, registerX, 120)
	} else if cycle > 160 && cycle <= 200 {
		Draw(cycle, registerX, 160)
	} else if cycle > 200 && cycle <= 240 {
		Draw(cycle, registerX, 200)
	}
}

func Draw(cycle int, registerX int, offset int) {
	if cycle-offset == registerX || cycle-offset == registerX+1 || cycle-offset == registerX+2 {
		output[cycle] = "#"
	} else {
		output[cycle] = "."
	}
}
