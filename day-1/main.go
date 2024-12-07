package main

import (
	"aoc-2024-day-1/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("The distance is: %v", findOutDistanceToSanta("puzzle_input.txt"))
}

func findOutDistanceToSanta(fileLocation string) int {

	leftList, rightList := readInputFile(fileLocation)

	if len(leftList) != len(rightList) {
		fmt.Printf("Something went horribly wrong...The left & right sides aren't equal!")
	}

	sortedLeft := utils.QuickSortStart(leftList)
	sortedRight := utils.QuickSortStart(rightList)

	totalDistance := 0

	for i := 0; i < len(sortedLeft); i++ {
		totalDistance += fintOutDifference(sortedLeft[i], sortedRight[i])
	}

	return totalDistance
}

func fintOutDifference(left, right int) int {
	fmt.Printf("Left: %v, Right: %v\n", left, right)
	if left > right {
		fmt.Printf("Left is greater. Difference: %v\n", left-right)
		return left - right
	} else {
		fmt.Printf("Right is greater. Difference: %v\n", right-left)
		return right - left
	}
}

func readInputFile(fileLocation string) ([]int, []int) {
	file, err := os.Open(fileLocation)

	if err != nil {

		fmt.Errorf("Error: %v", err)
		return nil, nil
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSuffix(line, "\n")

		splitted := strings.Split(line, "   ")

		i, err := strconv.ParseInt(splitted[0], 10, 64)
		if err != nil {
			fmt.Errorf("Error: %v", err)
		}
		leftList = append(leftList, int(i))

		i, err = strconv.ParseInt(splitted[1], 10, 64)
		if err != nil {
			fmt.Errorf("Error: %v", err)
		}
		rightList = append(rightList, int(i))

	}

	return leftList, rightList
}
