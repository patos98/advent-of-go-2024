package day03

import (
	"advent-of-go-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part2(inputPath string) {
	doIsOn := true
	sumOfAllProducts := 0
	for _, line := range utils.GetInputLines(inputPath) {
		for i, splitByDonts := range strings.Split(line, "don't()") {
			for j, splitByDos := range strings.Split(splitByDonts, "do()") {
				if doIsOn {
					sumOfAllProducts += executeMul(splitByDos)
				}

				if j < len(strings.Split(splitByDonts, "do()"))-1 {
					utils.Debugf("Do is on after line: %s\n", splitByDos)
					doIsOn = true
				}
			}

			if i < len(strings.Split(line, "don't()"))-1 {
				utils.Debugf("Do is off after line: %s\n", splitByDonts)
				doIsOn = false
			}
		}
	}

	fmt.Println("\nResult: ", sumOfAllProducts)
}

func executeMul(line string) int {
	sumOfAllProducts := 0

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(line, -1)
	for _, match := range matches {
		utils.Debugf("\nMul command: %s\n", match)

		operands := strings.Split(match, ",")
		operand1 := strings.Split(operands[0], "mul(")[1]
		operand2 := strings.Split(operands[1], ")")[0]

		utils.Debugf("Operand1: %s\n", operand1)
		utils.Debugf("Operand2: %s\n", operand2)

		op1, _ := strconv.Atoi(operand1)
		op2, _ := strconv.Atoi(operand2)

		sumOfAllProducts += op1 * op2
	}

	return sumOfAllProducts
}
