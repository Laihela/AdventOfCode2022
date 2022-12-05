package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(solvePart1("example.txt"))
	fmt.Println(solvePart1("input.txt"))
	fmt.Println(solvePart2("example.txt"))
	fmt.Println(solvePart2("input.txt"))
}

func solvePart1(fileName string) int {
	var (
		scoreHand     = []int{1, 2, 3}
		handByteToInt = map[byte]int{
			'A': 0,
			'B': 1,
			'C': 2,
			'X': 0,
			'Y': 1,
			'Z': 2,
		}
		resultScoreChart = [][]int{
			{3, 0, 6},
			{6, 3, 0},
			{0, 6, 3},
		}
	)

	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	gameStrategies := bytes.Split(content, []byte("\n"))
	totalScore := 0

	for _, gameStrategy := range gameStrategies {
		handOpponent := handByteToInt[gameStrategy[0]]
		handPlayer := handByteToInt[gameStrategy[2]]

		totalScore += scoreHand[handPlayer] + resultScoreChart[handPlayer][handOpponent]
	}

	return totalScore
}

func solvePart2(fileName string) int {
	var (
		scoreGame     = []int{0, 3, 6}
		handByteToInt = map[byte]int{
			'A': 0,
			'B': 1,
			'C': 2,
		}
		resultByteToInt = map[byte]int{
			'X': 0,
			'Y': 1,
			'Z': 2,
		}
		handScoreChart = [][]int{
			{3, 1, 2},
			{1, 2, 3},
			{2, 3, 1},
		}
	)

	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	gameStrategies := bytes.Split(content, []byte("\n"))
	totalScore := 0

	for _, gameStrategy := range gameStrategies {
		handOpponent := handByteToInt[gameStrategy[0]]
		gameResult := resultByteToInt[gameStrategy[2]]

		totalScore += handScoreChart[handOpponent][gameResult] + scoreGame[gameResult]
	}

	return totalScore
}
