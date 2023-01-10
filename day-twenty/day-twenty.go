package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
)

type Coordinate struct {
	id         string
	value      int64
	multiValue int64
}

var original = make([]Coordinate, 0)
var destination = make([]Coordinate, 0)
var zeroId string

//var multiplyFactor int64 = 811589153

var multiplyFactor int64 = 1

func main() {
	Read()

	for i := 0; i < 1; i++ {
		fmt.Printf("Count: %d\n", i)
		Decrypt()
	}

	var total int64 = 0
	total += Calculate(1000)
	total += Calculate(2000)
	total += Calculate(3000)

	fmt.Println(total)
}

func CalculateMultiplyFactor(value int64, divider int64) int64 {
	mod := value % divider

	return mod
}

func Read() {
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
		value, err := strconv.ParseInt(line, 10, 64)
		item := Coordinate{id: uuid.NewString(), value: value, multiValue: multiplyFactor * value}
		if item.value == 0 {
			zeroId = item.id
		}
		original = append(original, item)

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, item := range original {
		destination = append(destination, item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Decrypt() {
	for i := 0; i < len(original); i++ {
		left := make([]Coordinate, 0)
		right := make([]Coordinate, 0)

		if i%1000 == 0 {
			fmt.Printf("Decrypting value: %d with id: %s\n", original[i].value, original[i].id)
		}

		var indexInDestination = Find(original[i].id)

		var stepsToMove = CalculateMultiplyFactor(original[i].value*multiplyFactor, 5000)

		current := original[i]

		if stepsToMove == 0 {
			continue
		}

		var delta int64 = stepsToMove + indexInDestination - 1

		if delta > 0 {

			for delta > int64(len(original)) {
				delta -= int64(len(original)) - 1
			}
		} else {
			for delta <= 0 {
				delta += int64(len(original)) - 1
			}
		}

		destination = remove(destination, indexInDestination)
		for l := 0; int64(l) < delta; l++ {
			left = append(left, destination[l])
		}

		for r := delta; r < int64(len(destination)); r++ {
			right = append(right, destination[r])
		}

		destination = make([]Coordinate, 0)

		for _, item := range left {
			destination = append(destination, item)
		}

		destination = append(destination, current)

		for _, item := range right {
			destination = append(destination, item)
		}
	}
}

func Find(id string) int64 {
	var index int64 = -1

	for i := 0; i < len(destination); i++ {
		if destination[i].id == id {
			index = int64(i)
			break
		}
	}

	return index
}

func remove(slice []Coordinate, s int64) []Coordinate {
	return append(slice[:s], slice[s+1:]...)
}

func Calculate(count int64) int64 {
	var zeroIndex = Find(zeroId)

	fmt.Println("Calculating...")

	delta := zeroIndex + count

	for delta > int64(len(destination)) {
		delta -= int64(len(destination))
	}

	fmt.Println(destination[delta-1].multiValue)
	return destination[delta-1].multiValue
}
