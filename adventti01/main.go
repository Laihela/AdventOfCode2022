package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(solve("example.txt"))
	fmt.Println(solve("input.txt"))
}

func solve(fileName string) int {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	elfInventories := bytes.Split(content, []byte("\n\n"))
	firstHighestTotalCalories := 0
	secondHighestTotalCalories := 0
	thirdHighestTotalCalories := 0

	for _, inventory := range elfInventories {
		items := bytes.Split(inventory, []byte("\n"))
		totalCalories := 0

		for _, item := range items {
			calories, _ := strconv.Atoi(string(item))
			totalCalories += calories
		}

		if totalCalories > firstHighestTotalCalories {
			thirdHighestTotalCalories = secondHighestTotalCalories
			secondHighestTotalCalories = firstHighestTotalCalories
			firstHighestTotalCalories = totalCalories
		} else if totalCalories > secondHighestTotalCalories {
			thirdHighestTotalCalories = secondHighestTotalCalories
			secondHighestTotalCalories = totalCalories
		} else if totalCalories > thirdHighestTotalCalories {
			thirdHighestTotalCalories = totalCalories
		}
	}

	return firstHighestTotalCalories + secondHighestTotalCalories + thirdHighestTotalCalories
}
