package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	xPos  int64
	yPos  int64
	value string
}

func main() {
	itemCollection := Read()

	var minXPos = DiscoverXPos(&itemCollection, true)
	var maxXPos = DiscoverXPos(&itemCollection, false)

	fmt.Printf("Min x=%d, Max x=%d\n", minXPos, maxXPos)
	fmt.Println()

	CalculateOccupiedPositions(&itemCollection, minXPos, maxXPos)
}

func Read() [][2]item {
	var input = make([][2]item, 0)

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

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		sensorXPosString := strings.Replace(splitLine[2], "x=", "", 1)
		sensorXPosString = strings.Replace(sensorXPosString, ",", "", 1)

		sensorYPosString := strings.Replace(splitLine[3], "y=", "", 1)
		sensorYPosString = strings.Replace(sensorYPosString, ":", "", 1)

		beaconXPosString := strings.Replace(splitLine[8], "x=", "", 1)
		beaconXPosString = strings.Replace(beaconXPosString, ",", "", 1)

		beaconYPosString := strings.Replace(splitLine[9], "y=", "", 1)

		sensorXPos, _ := strconv.ParseInt(sensorXPosString, 10, 64)
		sensorYPos, _ := strconv.ParseInt(sensorYPosString, 10, 64)

		beaconXPos, _ := strconv.ParseInt(beaconXPosString, 10, 64)
		beaconYPos, _ := strconv.ParseInt(beaconYPosString, 10, 64)

		sensor := item{sensorXPos, sensorYPos, "S"}
		beacon := item{beaconXPos, beaconYPos, "B"}

		input = append(input, [2]item{sensor, beacon})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func DiscoverXPos(itemCollection *[][2]item, direction bool) int64 {
	var refItemCollection = *itemCollection

	var xPositions = make([]int64, 0)

	for _, item := range refItemCollection {
		sensor := item[0]
		beacon := item[1]

		xPositions = append(xPositions, sensor.xPos, beacon.xPos)
	}

	if direction {
		sort.Slice(xPositions, func(i, j int) bool { return xPositions[i] < xPositions[j] })
	} else {
		sort.Slice(xPositions, func(i, j int) bool { return xPositions[i] >= xPositions[j] })
	}

	return xPositions[0]
}

func CalculateOccupiedPositions(filteredCollection *[][2]item, minXPos int64, maxXPos int64) {
	refFilteredCollection := *filteredCollection

	positions := maxXPos - minXPos

	var occupiedPositions = make([]string, positions+100)

	for i := 0; i < len(occupiedPositions); i++ {
		occupiedPositions[i] = "."
	}

	for index, item := range refFilteredCollection {
		sensor := item[0]
		beacon := item[1]

		var deltaXStoB = sensor.xPos - beacon.xPos
		fmt.Printf("deltaX S%d-B%d = %d\n", index, index, deltaXStoB)

		fmt.Println(sensor.xPos - deltaXStoB)

		for position := sensor.xPos - deltaXStoB; position < sensor.xPos+deltaXStoB; position++ {
			occupiedPositions[position] = "#"
		}
	}

	var amount = 0

	for _, position := range occupiedPositions {
		if position == "#" {
			amount++
		}
	}

	totalAmount := amount - 1

	fmt.Printf("Total occupied positions: %d\n", totalAmount)
	fmt.Println()
}
