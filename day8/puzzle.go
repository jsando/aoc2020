package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inst struct {
	opcode  string
	operand int
}

func main() {
	program := make([]inst, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		opcode := line[0:3]
		operand, err := strconv.Atoi(line[4:])
		if err != nil {
			panic(err)
		}
		program = append(program, inst{opcode: opcode, operand: operand})
		fmt.Printf("%s %d\n", opcode, operand)
	}
	fmt.Printf("parse program, %d instructions\n", len(program))
	fmt.Printf("Part 1: accumulator value before re-executing inst: %d\n", part1(program))
	fmt.Printf("Part 2: accumulator value when program fixed: %d\n", part2(program))
}

func part1(program []inst) int {
	// track which lines were executed
	covered := make([]bool, len(program))
	pc := 0
	acc := 0
	for {
		if covered[pc] {
			break
		}
		covered[pc] = true
		inst := program[pc]
		fmt.Printf("%4d: %s %d (%d)\n", pc, inst.opcode, inst.operand, acc)
		switch inst.opcode {
		case "acc":
			acc += inst.operand
			pc++
		case "jmp":
			pc += inst.operand
		case "nop":
			pc++
		}
	}
	return acc
}

// exchange nop<->jmp until program terminates successfully
func part2(program []inst) int {
	for i := 0; i < len(program); i++ {
		oldOpcode := program[i].opcode
		switch program[i].opcode {
		case "nop":
			program[i].opcode = "jmp"
		case "jmp":
			program[i].opcode = "nop"
		default:
			continue
		}
		success, acc := execute(program)
		if success {
			return acc
		}
		program[i].opcode = oldOpcode
	}
	panic("didn't find successfull program execution!")
}

// run the program, return true and acc value if ran to completion, else false and acc value
func execute(program []inst) (bool, int) {
	// track which lines were executed
	covered := make([]bool, len(program))
	pc := 0
	acc := 0
	for {
		if pc >= len(program) {
			return true, acc
		}
		if covered[pc] {
			return false, acc
		}
		covered[pc] = true
		inst := program[pc]
		fmt.Printf("%4d: %s %d (%d)\n", pc, inst.opcode, inst.operand, acc)
		switch inst.opcode {
		case "acc":
			acc += inst.operand
			pc++
		case "jmp":
			pc += inst.operand
		case "nop":
			pc++
		}
	}
}
