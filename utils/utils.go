package utils

import (
	"fmt"
	"os"
	"strings"
)

var DebugMode bool = false

func GetInputLines(inputPath string) []string {
	byteContent, _ := os.ReadFile(inputPath)
	content := string(byteContent)
	lines := strings.Split(content, "\r\n")

	return lines
}

func Debugf(format string, a ...any) {
	if DebugMode {
		fmt.Printf(format, a...)
	}
}

func Debugln(a ...any) {
	if DebugMode {
		fmt.Println()
	}
}
