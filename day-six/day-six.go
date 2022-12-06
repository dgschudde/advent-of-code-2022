package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var input = Read()
	var position = Scan(input)

	fmt.Printf("Position: %d", position)
}

func Read() string {
	var input string = ""

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
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func Scan(input string) int {
	fmt.Println(input)

	var buffer [14]string
	var count int = 0

	for i := 0; i-14 < len(input); i++ {
		for bufferCount := 0; bufferCount < 14; bufferCount++ {
			buffer[bufferCount] = string(input[i+bufferCount])
		}

		if !duplicateInArray(buffer) {
			break
		} else {
			count = i
		}
	}

	return count + 15
}

func duplicateInArray(arr [14]string) bool {
	visited := make(map[string]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
