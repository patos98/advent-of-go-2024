package day17

import (
	"advent-of-go-2024/utils"
	"fmt"
	"slices"

	"github.com/tunabay/go-bitarray"
)

func Part2() {
	expectedOutput := []int{2, 4, 1, 1, 7, 5, 0, 3, 1, 4, 4, 0, 5, 5, 3}
	slices.Reverse(expectedOutput)
	possibleRegisterABits := []string{
		"0000011011",
		"0000111010",
		"0000000101", // ez lehet, nem fullosan jó és innen lefelé
		"0000100100",
		"0010000111",
		"0110000110",
	}

	for _, registerABits := range possibleRegisterABits {
		result := computeBits(registerABits, expectedOutput, 0)
		if result != "" {
			fmt.Println("Result: ", bitarray.MustParse(result).ToInt().Int64())
		}
	}
}

func computeBits(registerABits string, expectedOutput []int, outputIndex int) string {
	if outputIndex >= len(expectedOutput) {
		return registerABits
	}

	output := expectedOutput[outputIndex]
	ba := bitarray.MustParse(registerABits)

	utils.Debugln("Expected output: ", output)

	for lowerBits := 0; lowerBits < 8; lowerBits++ {
		lowerBitsString := fmt.Sprintf("%03b", byte(lowerBits))
		concatenated := ba.String() + lowerBitsString
		shift := lowerBits ^ 1
		end := len(concatenated) - shift
		expectedUpperBits := fmt.Sprintf("%03b", byte((output^101^lowerBits)%8))

		utils.Debugln("concat: ", concatenated)
		utils.Debugln("shift: ", shift)
		utils.Debugln("end: ", end)
		utils.Debugln("lowerbits: ", lowerBitsString)
		utils.Debugln("upperbits: ", concatenated[end-3:end])
		utils.Debugln("expected upperbits: ", expectedUpperBits)
		utils.Debugln()

		if concatenated[end-3:end] == expectedUpperBits {
			registerABits = computeBits(concatenated, expectedOutput, outputIndex+1)
			if registerABits != "" {
				return registerABits
			}
		}
	}

	return ""
}
