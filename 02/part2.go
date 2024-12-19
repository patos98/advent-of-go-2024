package day02

import (
	"advent-of-go-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2(inputPath string, debugMode bool) {
	minDiff := 1
	maxDiff := 3

	safeReportCount := 0
	for _, report := range utils.GetInputLines(inputPath) {

		if debugMode {
			fmt.Printf("\nChecking report: %s\n", report)
		}

		reportIsSafe := isReportSafe(report, minDiff, maxDiff, debugMode)

		if reportIsSafe {
			if debugMode {
				fmt.Printf("%s: Safe without removing any level.\n", report)
			}
			safeReportCount++
		} else {
			reportIsSafeWithRemovals := isReportSafeWithRemovals(report, minDiff, maxDiff, debugMode)
			if reportIsSafeWithRemovals {
				safeReportCount++
			}
		}
	}

	fmt.Println("Safe reports: ", safeReportCount)
}

func isReportSafeWithRemovals(report string, minDiff int, maxDiff int, debugMode bool) bool {
	levels := strings.Split(report, " ")
	for i := 0; i < len(levels); i++ {
		var newLevels []string
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		newReport := strings.Join(newLevels, " ")

		if debugMode {
			fmt.Printf("Checking new report: %s\n", newReport)
		}

		newReportIsSafe := isReportSafe(newReport, minDiff, maxDiff, debugMode)
		if newReportIsSafe {
			if debugMode {
				fmt.Printf("%s: Safe by removing the %d. level, %s.\n", report, i+1, levels[i])
			}
			return true
		}
	}

	if debugMode {
		fmt.Printf("%s: Unsafe regardless of which level is removed.\n", report)
	}

	return false
}

func isReportSafe(report string, minDiff int, maxDiff int, debugMode bool) bool {
	levels := strings.Split(report, " ")

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

			return false
		}

		if currentDiff > 0 && previousDiff < 0 || currentDiff < 0 && previousDiff > 0 {
			if debugMode {
				fmt.Printf("%s: Unsafe because %d and %d diffs mismatch.\n", report, previousDiff, currentDiff)
			}

			return false
		}

		absDiff := math.Abs(float64(currentDiff))
		if int(absDiff) < minDiff || int(absDiff) > maxDiff {
			if debugMode {
				fmt.Printf("%s: Unsafe because %d is not within the bounds of %d and %d\n", report, currentDiff, minDiff, maxDiff)
			}

			return false
		}

		previousDiff = currentDiff
	}

	return true
}
