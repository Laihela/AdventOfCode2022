package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

const (
	GroupSize = 3
)

func main() {
	fmt.Println(solvePart1("example.txt"))
	fmt.Println(solvePart1("input.txt"))
	fmt.Println(solvePart2("example.txt"))
	fmt.Println(solvePart2("input.txt"))
}

func solvePart1(fileName string) int {
	rucksacks := bytes.Split(readFile(fileName), []byte("\n"))
	totalPriority := 0

	for _, rucksack := range rucksacks {
		fullSize := len(rucksack)
		halfSize := fullSize / 2
		compartment1 := rucksack[0:halfSize]
		compartment2 := rucksack[halfSize:fullSize]
		item := findFirstDuplicate(compartment1, compartment2)
		totalPriority += getPriority(item)
	}

	return totalPriority
}

func solvePart2(fileName string) int {
	rucksacks := bytes.Split(readFile(fileName), []byte("\n"))
	groups := [][][]byte{}
	totalPriority := 0

	for i := 0; i < len(rucksacks); i += GroupSize {
		end := i + GroupSize
		groups = append(groups, rucksacks[i:end])
	}

	for _, sacks := range groups {
		badgeItem := findFirstCommon(sacks[0], sacks[1], sacks[2])
		totalPriority += getPriority(badgeItem)
	}

	return totalPriority
}

func getPriority(item byte) int {
	if unicode.IsUpper(rune(item)) {
		return int(item) - 38
	} else {
		return int(item) - 96
	}
}

func findFirstCommon(array1, array2, array3 []byte) byte {
	for i1 := 0; i1 < len(array1); i1++ {
		for i2 := 0; i2 < len(array2); i2++ {
			for i3 := 0; i3 < len(array3); i3++ {
				if array1[i1] == array2[i2] && array1[i1] == array3[i3] {
					return array1[i1]
				}
			}
		}
	}
	return 0
}

func findFirstDuplicate(array1, array2 []byte) byte {
	for i1 := 0; i1 < len(array1); i1++ {
		for i2 := 0; i2 < len(array2); i2++ {
			if array1[i1] == array2[i2] {
				return array1[i1]
			}
		}
	}
	return 0
}

func readFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
