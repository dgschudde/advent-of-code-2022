package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

func main() {
	var layout [9]Stack

	Initialize(&layout)

	Read(&layout)
}

func Initialize(stack *[9]Stack) {
	stack[0].Push("B")
	stack[0].Push("Q")
	stack[0].Push("C")

	stack[1].Push("R")
	stack[1].Push("Q")
	stack[1].Push("W")
	stack[1].Push("Z")

	stack[2].Push("B")
	stack[2].Push("M")
	stack[2].Push("R")
	stack[2].Push("L")
	stack[2].Push("V")

	stack[3].Push("C")
	stack[3].Push("Z")
	stack[3].Push("H")
	stack[3].Push("V")
	stack[3].Push("T")
	stack[3].Push("W")

	stack[4].Push("D")
	stack[4].Push("Z")
	stack[4].Push("H")
	stack[4].Push("B")
	stack[4].Push("N")
	stack[4].Push("V")
	stack[4].Push("G")

	stack[5].Push("H")
	stack[5].Push("N")
	stack[5].Push("P")
	stack[5].Push("C")
	stack[5].Push("J")
	stack[5].Push("F")
	stack[5].Push("V")
	stack[5].Push("Q")

	stack[6].Push("D")
	stack[6].Push("G")
	stack[6].Push("T")
	stack[6].Push("R")
	stack[6].Push("W")
	stack[6].Push("Z")
	stack[6].Push("S")

	stack[7].Push("C")
	stack[7].Push("G")
	stack[7].Push("M")
	stack[7].Push("N")
	stack[7].Push("B")
	stack[7].Push("W")
	stack[7].Push("Z")
	stack[7].Push("P")

	stack[8].Push("N")
	stack[8].Push("J")
	stack[8].Push("B")
	stack[8].Push("M")
	stack[8].Push("W")
	stack[8].Push("Q")
	stack[8].Push("F")
	stack[8].Push("P")
}

func Read(stack *[9]Stack) {
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
		actualLine := scanner.Text()

		values := strings.Split(actualLine, " ")
		amount, err1 := strconv.ParseInt(values[1], 10, 64)
		from, err2 := strconv.ParseInt(values[3], 10, 64)
		to, err3 := strconv.ParseInt(values[5], 10, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal(err)
		}

		Move(amount, from, to, stack)
	}

	fmt.Printf("Result:%s%s%s%s%s%s%s%s%s\n", stack[0].Pop(), stack[1].Pop(), stack[2].Pop(), stack[3].Pop(), stack[4].Pop(), stack[5].Pop(), stack[6].Pop(), stack[7].Pop(), stack[8].Pop())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Move(amount int64, from int64, to int64, stack *[9]Stack) {
	fmt.Printf("Move %d from %d to %d\n", amount, from, to)

	var tempStack Stack

	var i int64 = 1
	for ; i <= amount; i++ {
		tempValue := stack[from-1].Pop()

		// Push to temp stack
		tempStack.Push(tempValue)

	}

	// Pop temp stack until it's empty
	for !tempStack.IsEmpty() {
		var value = tempStack.Pop()
		// Push item to destination stack
		stack[to-1].Push(value)
	}
}
