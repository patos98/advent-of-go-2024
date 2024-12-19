package day17

import (
	"advent-of-go-2024/utils"
	"log"
	"math"
)

type Computer struct {
	registerA          int
	registerB          int
	registerC          int
	instructionPointer int
	output             chan int
}

func (c *Computer) executeProgram(program []int) {
	instructions := c.getInstructionsByOpCodes()
	programLength := len(program)
	utils.Debugf("Program length: %d\n", programLength)

	for {
		utils.Debugf("Instruction pointer: %d\n", c.instructionPointer)
		if c.instructionPointer >= programLength-1 {
			return
		}

		opCode := program[c.instructionPointer]
		operand := program[c.instructionPointer+1]
		utils.Debugf("Op code: %d\n", opCode)
		utils.Debugf("Operand: %d\n", operand)

		instructionPointerMoved := instructions[opCode](operand)
		utils.Debugf("Instruction pointer moved: %t\n", instructionPointerMoved)
		if !instructionPointerMoved {
			c.instructionPointer += 2
		}
	}
}

func (c *Computer) getInstructionsByOpCodes() map[int]func(int) bool {
	return map[int]func(int) bool{
		0: c.adv,
		1: c.bxl,
		2: c.bst,
		3: c.jnz,
		4: c.bxc,
		5: c.out,
		6: c.bdv,
		7: c.cdv,
	}
}

func (c *Computer) adv(operand int) (instructionPointerMoved bool) {
	denominator := math.Pow(2, float64(c.getComboOperand(operand)))
	ratio := float64(c.registerA) / denominator
	c.registerA = int(math.Floor(ratio))
	return
}

func (c *Computer) bxl(operand int) (instructionPointerMoved bool) {
	c.registerB = c.registerB ^ operand
	return
}

func (c *Computer) bst(operand int) (instructionPointerMoved bool) {
	c.registerB = c.getComboOperand(operand) % 8
	return
}

func (c *Computer) jnz(operand int) (instructionPointerMoved bool) {
	if c.registerA == 0 {
		return
	}

	c.instructionPointer = operand
	return true
}

func (c *Computer) bxc(operand int) (instructionPointerMoved bool) {
	c.registerB = c.registerB ^ c.registerC
	return
}

func (c *Computer) out(operand int) (instructionPointerMoved bool) {
	c.output <- c.getComboOperand(operand) % 8
	return
}

func (c *Computer) bdv(operand int) (instructionPointerMoved bool) {
	denominator := math.Pow(2, float64(c.getComboOperand(operand)))
	ratio := float64(c.registerA) / denominator
	c.registerB = int(math.Floor(ratio))
	return
}

func (c *Computer) cdv(operand int) (instructionPointerMoved bool) {
	denominator := math.Pow(2, float64(c.getComboOperand(operand)))
	ratio := float64(c.registerA) / denominator
	c.registerC = int(math.Floor(ratio))
	return
}

func (c *Computer) getComboOperand(comboOperand int) int {
	if comboOperand < 4 {
		return comboOperand
	}

	if comboOperand == 4 {
		return c.registerA
	}

	if comboOperand == 5 {
		return c.registerB
	}

	if comboOperand == 6 {
		return c.registerC
	}

	log.Panicf("Invalid combo operand: %d", comboOperand)
	return -1
}
