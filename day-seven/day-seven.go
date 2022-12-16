package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tree struct {
	Child   *Tree
	Parent  *Tree
	Current *Tree
	name    string
	files   map[string]int64
}

var directoryStructure *Tree

func insert(t *Tree, name string, file string, size int64) *Tree {
	if t == nil {
		return &Tree{nil, nil, nil, name, nil}
	}

	if t.files == nil {
		t.files = make(map[string]int64)
	}

	if file != "" {
		t.files[file] = size
	}

	t = insert(t, name, file, size)
	return t
}

func navigateup(t *Tree) *Tree {
	if t != nil && t.Parent != nil {
		t.Current = t.Parent
	}

	return t
}

func navigateto(t *Tree, name string) *Tree {
	if t != nil {
		if t.name == name {
			t.Current = t
		} else {
			navigateto(t.Child, name)
		}
	}

	return t
}

func main() {
	//var commands = Read()
	//ParseCommands(&commands)
	Play()
	fmt.Println(directoryStructure)
}

func Read() []string {
	var input []string
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
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func Play() {
	directoryStructure = insert(nil, "root", "", 0)
	fmt.Println(directoryStructure.name)
	fmt.Printf("Current: %x\n", directoryStructure.Current)
	fmt.Printf("Parent: %x\n", directoryStructure.Parent)
	fmt.Printf("Address: %x\n", &directoryStructure)
}

func ParseCommands(commands *[]string) {
	var refCommands = *commands
	var currentNode *Tree

	for index, line := range refCommands {
		splitLine := strings.Split(line, " ")
		action := splitLine[0]
		command := splitLine[1]

		if action == "$" && command == "cd" {
			directory := splitLine[2]

			if directory == "/" {
				directoryStructure = insert(nil, "root", "", 0)
				fmt.Println(directoryStructure.name)
			} else if directory == ".." {
				//currentNode = navigateup(directoryStructure)
				//fmt.Printf("Navigate up %s", currentNode.name)
			} else {
				//currentNode = navigateto(directoryStructure, directory)
				//fmt.Printf("Navigate to %s", currentNode.name)
			}
		}

		if action == "$" && command == "ls" {
			var contents = ""

			for !strings.HasPrefix(contents, "$") && index < len(refCommands)-1 {
				index++
				contents = refCommands[index]
				splitContents := strings.Split(contents, " ")

				if splitContents[0] == "$" {
					break
				}

				if splitContents[0] == "dir" {
					currentNode = insert(directoryStructure, splitContents[1], "", 0)
					fmt.Printf("Create dir '%s' -> %s\n", splitContents[1], currentNode.name)
				} else {
					//size, _ := strconv.ParseInt(splitContents[1], 10, 64)
					//currentNode = insert(directoryStructure, "", splitContents[0], size)
					fmt.Printf("Create file '%s' with size '%s'\n", splitContents[1], splitContents[0])
				}
			}
		}
	}
}
