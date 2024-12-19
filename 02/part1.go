package day02

import (
	"advent-of-go-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(inputPath string, debugMode bool) {
	minDiff := 1
	maxDiff := 3

	safeReportCount := 0
	for _, report := range utils.GetInputLines(inputPath) {
		levels := strings.Split(report, " ")

		safe := true
		previousDiff := 0
		for i := 0; i < len(levels)-1; i++ {
			level1, _ := strconv.Atoi(levels[i])
			level2, _ := strconv.Atoi(levels[i+1])
			currentDiff := level1 - level2

			if debugMode {
				// fmt.Printf("Checking: %d %d\n", level1, level2)
			}

			if currentDiff == 0 {
				if debugMode {
					fmt.Printf("%s: Unsafe because %d %d is neither an increase or a decrease.\n", report, level1, level2)
				}

				safe = false
				break
			}

			if currentDiff > 0 && previousDiff < 0 || currentDiff < 0 && previousDiff > 0 {
				if debugMode {
					fmt.Printf("%s: Unsafe because %d and %d diffs mismatch.\n", report, previousDiff, currentDiff)
				}

				safe = false
				break
			}

			absDiff := math.Abs(float64(currentDiff))
			if int(absDiff) < minDiff || int(absDiff) > maxDiff {
				if debugMode {
					fmt.Printf("%s: Unsafe because %d is not within the bounds of %d and %d\n", report, currentDiff, minDiff, maxDiff)
				}

				safe = false
				break
			}

			previousDiff = currentDiff
		}

		if safe {
			if debugMode {
				fmt.Printf("%s: Safe\n", report)
			}
			safeReportCount++
		}
	}

	fmt.Println("Safe reports: ", safeReportCount)
}
