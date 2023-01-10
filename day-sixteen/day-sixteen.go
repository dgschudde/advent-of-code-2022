package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	Value    string
	FlowRate int32
	IsOpen   bool
	PathTo   []*Valve
}

var Valves []Valve

var ValveLookUp = make(map[string]*Valve, 0)

func main() {
	inputLines := Read()

	ParseInput(&inputLines)

	Traverse(&Valves[0], 1)
}

func Read() []string {
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

	var inputLines = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputLines
}

func ParseInput(inputLines *[]string) {
	refInput := *inputLines

	// Get all valves
	for _, line := range refInput {
		splitLine := strings.Split(line, " ")

		sensorValue := splitLine[1]

		unparsedRate := splitLine[4]
		unparsedRate = strings.Replace(unparsedRate, "rate=", "", 1)
		flowRateString := strings.Replace(unparsedRate, ";", "", 1)
		flowRate, _ := strconv.ParseInt(flowRateString, 10, 32)

		valve := Valve{Value: sensorValue, FlowRate: int32(flowRate), IsOpen: true}
		Valves = append(Valves, valve)
		ValveLookUp[valve.Value] = &valve
	}

	for _, line := range refInput {
		splitLine := strings.Split(line, " ")
		lookupSensor := splitLine[1]

		valvesToAdd := make([]string, 0)

		for i := 9; i < len(splitLine); i++ {
			splitSensor := splitLine[i]
			splitSensor = strings.Replace(splitSensor, ",", "", 1)
			valvesToAdd = append(valvesToAdd, splitSensor)
		}

		foundValve := ValveLookUp[lookupSensor]

		for _, valve := range valvesToAdd {
			sensorToAdd := ValveLookUp[valve]
			foundValve.PathTo = append(foundValve.PathTo, sensorToAdd)
		}
		index := Find(foundValve)
		Valves[index] = *foundValve
	}
}

func Find(item *Valve) int {
	itemToSearch := *item
	index := 0

	for _, i := range Valves {

		if itemToSearch.Value == i.Value {
			return index
		}
		index++
	}
	return -1
}

func Traverse(valve *Valve, timeSpent int) {
	currentValve := valve
	nextValve := valve

	closedValves := make([]*Valve, 0)
	openValves := make([]*Valve, 0)
	stuckValves := make([]*Valve, 0)

	if timeSpent > 30 {
		return
	}

	// Output the valves that are visited
	PrintTraverse(timeSpent, currentValve.Value)

	for i := 0; i < len(currentValve.PathTo); i++ {
		// Select valves that are not open
		if currentValve.PathTo[i].IsOpen == false && currentValve.PathTo[i].FlowRate > 0 {
			closedValves = append(closedValves, currentValve.PathTo[i])
		} else if currentValve.PathTo[i].IsOpen == true && currentValve.PathTo[i].FlowRate > 0 {
			openValves = append(openValves, currentValve.PathTo[i])
		} else if currentValve.PathTo[i].FlowRate == 0 {
			stuckValves = append(stuckValves, currentValve.PathTo[i])
		}

		// Get open valves and count the to total released pressure of all open valves
		// Count the pressure released every minute

		// When all valves are open, select the first valve
	}

	if currentValve.FlowRate > 0 {
		// Set the valve to open
		currentValve.IsOpen = false

		// Update current valve in valves array
		index := Find(currentValve)
		Valves[index] = *currentValve
	}

	// Select the first valve
	if len(openValves) > 0 {
		nextValve = openValves[0]
	} else if len(stuckValves) > 0 {
		nextValve = stuckValves[0]
	} else if len(closedValves) > 0 {
		nextValve = closedValves[0]
	}

	// Increase the time spent searching
	timeSpent++

	// Traverse to the next valve
	Traverse(nextValve, timeSpent)
}

func Print(timeSpent int, openVales map[string]int32, value string) {
	fmt.Printf("== Minute %d ==\n", timeSpent)
	fmt.Printf("%d valves are open\n", len(openVales))
	fmt.Printf("You move to valve %s\n", value)
	fmt.Println()
}

func PrintTraverse(timeSpent int, value string) {
	fmt.Printf("== Minute %d ==\n", timeSpent)
	fmt.Printf("You are at valve %s\n", value)
	fmt.Println()
}

func GetOpenValves() map[string]int32 {
	openValves := make(map[string]int32, 0)

	for _, valve := range Valves {
		if valve.IsOpen {
			openValves[valve.Value] = valve.FlowRate
		}
	}

	return openValves
}

func GetNextValve(currentValve *Valve) *Valve {
	var nextSensor *Valve

	for _, item := range currentValve.PathTo {
		if !item.IsOpen && item.FlowRate > 0 {
			nextSensor = item
		} else {
			nextSensor = item
		}
	}

	return nextSensor
}

func UpdateValve(sensor *Valve) {
	index := Find(sensor)
	Valves[index] = *sensor
}
