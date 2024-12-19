package day04

import (
	"advent-of-go-2024/utils"
	"fmt"
	"strings"
)

func check(grid [][]string, rowIndex, colIndex int, word, wordInReverse string) (match bool, indexes []Indexes) {
	if rowIndex+2 >= len(grid) || colIndex+2 >= len(grid[0]) {
		return
	}

	diagonal1 := grid[rowIndex][colIndex] + grid[rowIndex+1][colIndex+1] + grid[rowIndex+2][colIndex+2]
	diagonal2 := grid[rowIndex+2][colIndex] + grid[rowIndex+1][colIndex+1] + grid[rowIndex][colIndex+2]
	match = (diagonal1 == word || diagonal1 == wordInReverse) && (diagonal2 == word || diagonal2 == wordInReverse)
	if match {
		utils.Debugf("Match at indexes: %d,%d\n", rowIndex, colIndex)
		utils.Debugf("%s.%s\n", string(diagonal1[0]), string(diagonal2[2]))
		utils.Debugf(".%s.\n", string(diagonal1[1]))
		utils.Debugf("%s.%s\n", string(diagonal2[0]), string(diagonal1[2]))
	}

	return match, []Indexes{
		{rowIndex: rowIndex, colIndex: colIndex},
		{rowIndex: rowIndex + 1, colIndex: colIndex + 1},
		{rowIndex: rowIndex + 2, colIndex: colIndex + 2},
		{rowIndex: rowIndex + 2, colIndex: colIndex},
		{rowIndex: rowIndex, colIndex: colIndex + 2},
	}
}

func Part2(inputPath string) {
	grid := [][]string{}
	for _, line := range utils.GetInputLines(inputPath) {
		grid = append(grid, strings.Split(line, ""))
	}

	word := "MAS"
	wordInReverse := "SAM"

	matchingIndexes := map[Indexes]struct{}{}
	totalMatchCount := 0
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			match, indexes := check(grid, rowIndex, colIndex, word, wordInReverse)
			if match {
				totalMatchCount++
				for _, i := range indexes {
					matchingIndexes[i] = struct{}{}
				}
			}
		}
	}

	if utils.DebugMode {
		for rowIndex, line := range grid {
			for colIndex, letter := range line {
				if _, ok := matchingIndexes[Indexes{rowIndex: rowIndex, colIndex: colIndex}]; ok {
					fmt.Print(letter)
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}

	fmt.Println("Total match count: ", totalMatchCount)
}
