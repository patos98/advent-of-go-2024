package day17

import "fmt"

func Part1(input Input) {
	computer := Computer{
		registerA: input.registerA,
		registerB: input.registerB,
		registerC: input.registerC,
		output:    make(chan int),
	}

	output := []int{}
	go func() {
		for o := range computer.output {
			output = append(output, o)
		}
	}()

	computer.executeProgram(input.program)
	fmt.Println("Output: ", output)
}
