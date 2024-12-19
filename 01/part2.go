package day01

import (
	"advent-of-go-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part2(inputPath string) {
	debugMode := strings.Contains(inputPath, "test")

	var leftNumbers []int
	rightNumbersAppearances := map[int]int{}
	for _, line := range utils.GetInputLines(inputPath) {
		numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])
		leftNumbers = append(leftNumbers, leftNumber)
		appearances := rightNumbersAppearances[rightNumber]
		rightNumbersAppearances[rightNumber] = appearances + 1
	}

	similarityScore := 0
	length := len(leftNumbers)
	for i := 0; i < length; i++ {
		leftNumber := leftNumbers[i]
		appearances := rightNumbersAppearances[leftNumber]
		score := leftNumber * appearances
		similarityScore += score

		if debugMode {
			fmt.Println(leftNumber, " * ", appearances, " = ", score)
		}
	}

	fmt.Println("Similarity score: ", similarityScore)
}
