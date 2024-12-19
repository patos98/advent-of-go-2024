package day03

import (
	"advent-of-go-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(inputPath string) {
	sumOfAllProducts := 0

	for _, line := range utils.GetInputLines(inputPath) {
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
	}

	fmt.Println("\nResult: ", sumOfAllProducts)
}
