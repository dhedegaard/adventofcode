package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	opcode   string
	register string
	offset   int
}

// Executes a given instruction, returns the new program counter.
func executeInst(memoryPtr *map[string]uint, instructions *[]instruction, pc int) int {
	// Fetch memory and current instruction.
	memory := *memoryPtr
	inst := (*instructions)[pc]

	// Execute a given opcode, with the optional register/offset provided.
	switch inst.opcode {
	case "hlf":
		memory[inst.register] /= 2
		break
	case "tpl":
		memory[inst.register] *= 3
		break
	case "inc":
		memory[inst.register]++
		break
	case "jmp":
		pc += inst.offset - 1
		break
	case "jie":
		if memory[inst.register]%2 == 0 {
			pc += inst.offset - 1
		}
		break
	case "jio":
		if memory[inst.register] == 1 {
			pc += inst.offset - 1
		}
		break
	default:
		panic(fmt.Sprint("Don't know how to execute:", inst))
	}

	// Always increment the program counter, therefore the jumps subtracts
	// the expected position by 1 to negate this.
	return pc + 1
}

// string -> int, or return default.
func strToInt(input string, def int) int {
	// If input is empty, the regex match failed and defaulted to the empty
	// string.
	if len(input) == 0 {
		return def
	}
	// If the input starts with +, remove it since atoi() parses a positive
	// number anyways.
	if input[0] == '+' {
		input = input[1:]
	}
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Parse program into slice of structs.
	before := time.Now()
	reParse := regexp.MustCompile(`(.{3}) ([ab])?,? ?([+-]\d+)?`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	var instructions []instruction
	for scanner.Scan() {
		// Match line against regexp.
		rc := reParse.FindStringSubmatch(scanner.Text())
		// Append an instruction to the instructionset.
		instructions = append(instructions, instruction{
			opcode:   rc[1],
			register: rc[2],
			offset:   strToInt(rc[3], 0),
		})
	}
	fmt.Println("Parsed", len(instructions), "instructions in",
		time.Now().Sub(before))

	// Initialize state.
	pc := 0
	memory := map[string]uint{
		"a": 0,
		"b": 0,
	}
	before = time.Now()
	// Execute instructions.
	for pc >= 0 && pc < len(instructions) {
		pc = executeInst(&memory, &instructions, pc)
	}
	fmt.Println("END part1! memory:", memory, "pc:", pc, "took:", time.Now().Sub(before))

	// Now run again, where a starts as 1.
	pc = 0
	memory = map[string]uint{
		"a": 1,
		"b": 0,
	}
	before = time.Now()
	// Execute instructions.
	for pc >= 0 && pc < len(instructions) {
		pc = executeInst(&memory, &instructions, pc)
	}
	fmt.Println("END part2! memory:", memory, "pc:", pc, "took:", time.Now().Sub(before))
}

const input = `jio a, +16
inc a
inc a
tpl a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
tpl a
inc a
jmp +23
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`
