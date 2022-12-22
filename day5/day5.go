package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

const stackWidth = 4

type operatorInstuction struct {
	FromStackIdx int
	ToStackIdx   int
	NumElements  int
}

func readInitialStacks(s *bufio.Scanner) [][]string {
	stacks := [][]string{}
	for {
		hasMore := s.Scan()
		if !hasMore {
			break
		}
		line := s.Text()
		if len(line) == 0 {
			break
		}
		if !strings.Contains(line, "[") {
			// We've reached the "stack index" line
			continue
		}

		// Each stack is fixed-width, 4 characters
		numStacks := (len(line) + 1) / 4
		if len(stacks) == 0 {
			stacks = make([][]string, numStacks)
		}
		for i := 0; i < numStacks; i++ {
			offset := i * stackWidth
			crate := line[offset : offset+stackWidth-1]
			crate = crate[1:2]
			if crate == " " {
				continue
			}
			stacks[i] = append(stacks[i], crate)
		}
	}

	// we read the stacks from top to bottom so we need to reverse their order
	for _, stack := range stacks {
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}
	}
	return stacks
}

func parseInstruction(line string) operatorInstuction {
	parts := strings.Split(line, " ")
	if len(parts) != 6 {
		log.Fatalln("bad instruction: ", line)
	}
	_numElements := parts[1]
	_fromStackIdx := parts[3]
	_toStackIdx := parts[5]
	numElements, err := strconv.Atoi(_numElements)
	if err != nil {
		log.Fatalln(err)
	}
	fromStackIdx, err := strconv.Atoi(_fromStackIdx)
	if err != nil {
		log.Fatalln(err)
	}
	toStackIdx, err := strconv.Atoi(_toStackIdx)
	if err != nil {
		log.Fatalln(err)
	}
	return operatorInstuction{
		FromStackIdx: fromStackIdx - 1,
		ToStackIdx:   toStackIdx - 1,
		NumElements:  numElements,
	}
}

func executePart1Instruction(instruction operatorInstuction, stacks [][]string) {
	fromStack := stacks[instruction.FromStackIdx]
	toStack := stacks[instruction.ToStackIdx]
	if len(fromStack) < instruction.NumElements {
		log.Fatalf("fromStack only has %d crates, cannot pop %d", len(fromStack), instruction.NumElements)
	}
	for i := 0; i < instruction.NumElements; i++ {
		// pop from
		var crate string
		crate, fromStack = fromStack[len(fromStack)-1], fromStack[:len(fromStack)-1]
		// push to
		toStack = append(toStack, crate)
	}
	stacks[instruction.FromStackIdx] = fromStack
	stacks[instruction.ToStackIdx] = toStack
}

func Day5_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := readInitialStacks(scanner)
	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		executePart1Instruction(instruction, stacks)
	}
	var sb strings.Builder
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		sb.WriteString(stack[len(stack)-1])
	}
	log.Printf("final crates at the top of the stacks: %s", sb.String())
}

// instead of doing a pop, push per element, we move all at the same time
func executePart2Instruction(instruction operatorInstuction, stacks [][]string) {
	fromStack := stacks[instruction.FromStackIdx]
	toStack := stacks[instruction.ToStackIdx]
	if len(fromStack) < instruction.NumElements {
		log.Fatalf("fromStack only has %d crates, cannot move %d", len(fromStack), instruction.NumElements)
	}
	// this is an array "cut"
	// grab the top N elements
	startIdx := len(fromStack) - instruction.NumElements
	crates := make([]string, instruction.NumElements)
	copy(crates, fromStack[startIdx:])

	// Remove them from the old stack
	fromStack = fromStack[:startIdx]

	// Add them to the new stack
	toStack = append(toStack, crates...)

	stacks[instruction.FromStackIdx] = fromStack
	stacks[instruction.ToStackIdx] = toStack
}

func Day5_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := readInitialStacks(scanner)
	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		executePart2Instruction(instruction, stacks)
	}
	var sb strings.Builder
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		sb.WriteString(stack[len(stack)-1])
	}
	log.Printf("final crates at the top of the stacks: %s", sb.String())
}
