package day04

import (
	"advent-of-go-2024/utils"
	"fmt"
	"strings"
)

type CheckerFunction = func([][]string, int, int, string, string) (bool, []Indexes)

type Indexes struct {
	rowIndex int
	colIndex int
}

func checkTopDown(grid [][]string, rowIndex, colIndex int, word, wordInReverse string) (match bool, indexes []Indexes) {
	indexes = []Indexes{
		{rowIndex: rowIndex, colIndex: colIndex},
		{rowIndex: rowIndex + 1, colIndex: colIndex},
		{rowIndex: rowIndex + 2, colIndex: colIndex},
		{rowIndex: rowIndex + 3, colIndex: colIndex},
	}
	buffer := getBuffer(grid, indexes)
	match = buffer == word || buffer == wordInReverse
	if match {
		utils.Debugf("Top down match at position: %d,%d\n", rowIndex, colIndex)
	}
	return
}

func checkLeftRight(grid [][]string, rowIndex, colIndex int, word, wordInReverse string) (match bool, indexes []Indexes) {
	indexes = []Indexes{
		{rowIndex: rowIndex, colIndex: colIndex},
		{rowIndex: rowIndex, colIndex: colIndex + 1},
		{rowIndex: rowIndex, colIndex: colIndex + 2},
		{rowIndex: rowIndex, colIndex: colIndex + 3},
	}
	buffer := getBuffer(grid, indexes)
	match = buffer == word || buffer == wordInReverse
	if match {
		utils.Debugf("Left right match at position: %d,%d\n", rowIndex, colIndex)
	}
	return
}

func checkTopLeftDiagonal(grid [][]string, rowIndex, colIndex int, word, wordInReverse string) (match bool, indexes []Indexes) {
	indexes = []Indexes{
		{rowIndex: rowIndex, colIndex: colIndex},
		{rowIndex: rowIndex + 1, colIndex: colIndex + 1},
		{rowIndex: rowIndex + 2, colIndex: colIndex + 2},
		{rowIndex: rowIndex + 3, colIndex: colIndex + 3},
	}
	buffer := getBuffer(grid, indexes)
	match = buffer == word || buffer == wordInReverse
	if match {
		utils.Debugf("Top left match at position: %d,%d\n", rowIndex, colIndex)
	}
	return
}

func checkBottomLeftDiagonal(grid [][]string, rowIndex, colIndex int, word, wordInReverse string) (match bool, indexes []Indexes) {
	indexes = []Indexes{
		{rowIndex: rowIndex, colIndex: colIndex},
		{rowIndex: rowIndex - 1, colIndex: colIndex + 1},
		{rowIndex: rowIndex - 2, colIndex: colIndex + 2},
		{rowIndex: rowIndex - 3, colIndex: colIndex + 3},
	}
	buffer := getBuffer(grid, indexes)
	match = buffer == word || buffer == wordInReverse
	if match {
		utils.Debugf("Bottom left match at position: %d,%d\n", rowIndex, colIndex)
	}
	return
}

func getBuffer(grid [][]string, indexes []Indexes) (buffer string) {
	lastIndexes := indexes[len(indexes)-1]
	if lastIndexes.rowIndex < 0 || lastIndexes.rowIndex >= len(grid) || lastIndexes.colIndex < 0 || lastIndexes.colIndex >= len(grid[0]) {
		return
	}

	for _, i := range indexes {
		buffer += grid[i.rowIndex][i.colIndex]
	}
	return
}

var checkerFunctions = []CheckerFunction{
	checkTopDown,
	checkLeftRight,
	checkTopLeftDiagonal,
	checkBottomLeftDiagonal,
}

func Part1(inputPath string) {
	grid := [][]string{}
	for _, line := range utils.GetInputLines(inputPath) {
		grid = append(grid, strings.Split(line, ""))
	}

	word := "XMAS"
	wordInReverse := "SAMX"

	matchingIndexes := map[Indexes]struct{}{}
	totalMatchCount := 0
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			for _, checkerFunction := range checkerFunctions {
				match, indexes := checkerFunction(grid, rowIndex, colIndex, word, wordInReverse)
				if match {
					totalMatchCount++
					for _, i := range indexes {
						matchingIndexes[i] = struct{}{}
					}
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
