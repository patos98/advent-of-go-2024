package day01

import (
	"advent-of-go-2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(inputPath string) {
	debugMode := strings.Contains(inputPath, "test")

	var leftNumbers, rightNumbers []int
	for _, line := range utils.GetInputLines(inputPath) {
		numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}

	if debugMode {
		fmt.Println("Left numbers: ", leftNumbers)
		fmt.Println("Right numbers: ", rightNumbers)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	if debugMode {
		fmt.Println("Left numbers sorted: ", leftNumbers)
		fmt.Println("Right numbers sorted: ", rightNumbers)
	}

	totalDistance := 0
	length := len(leftNumbers)
	for i := 0; i < length; i++ {
		leftNumber := leftNumbers[i]
		rightNumber := rightNumbers[i]
		distance := int(math.Abs(float64(leftNumber) - float64(rightNumber)))
		totalDistance += distance

		if debugMode {
			fmt.Println(leftNumber, " - ", rightNumber, " = ", distance)
		}
	}

	fmt.Println("Total distance: ", totalDistance)
}
