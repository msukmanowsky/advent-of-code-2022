package day5

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

var initalStackInput string = `
                        [Z] [W] [Z]
        [D] [M]         [L] [P] [G]
    [S] [N] [R]         [S] [F] [N]
    [N] [J] [W]     [J] [F] [D] [F]
[N] [H] [G] [J]     [H] [Q] [H] [P]
[V] [J] [T] [F] [H] [Z] [R] [L] [M]
[C] [M] [C] [D] [F] [T] [P] [S] [S]
[S] [Z] [M] [T] [P] [C] [D] [C] [D]
 1   2   3   4   5   6   7   8   9 

`

func Test_readInitialStacks(t *testing.T) {
	reader := strings.NewReader(initalStackInput)
	scanner := bufio.NewScanner(reader)
	scanner.Scan() // initial newline
	stacks := readInitialStacks(scanner)
	if numStacks := len(stacks); numStacks != 9 {
		t.Errorf("wanted 9 stacks, got: %d", numStacks)
	}
	stack := stacks[0]
	// pop
	var top string
	top, stack = stack[len(stack)-1], stack[:len(stack)-1]
	if top != "N" {
		t.Errorf("wanted top of stack 0 to be N, got: %+v", top)
	}
}

type parseInstructionTest struct {
	line string
	want operatorInstuction
}

var parseInstructionTests = []parseInstructionTest{
	{
		line: "move 3 from 9 to 6",
		want: operatorInstuction{
			FromStackIdx: 8,
			ToStackIdx:   5,
			NumElements:  3,
		},
	},
	{
		line: "move 11 from 1 to 8",
		want: operatorInstuction{
			FromStackIdx: 0,
			ToStackIdx:   7,
			NumElements:  11,
		},
	},
	{
		line: "move 17 from 13 to 15",
		want: operatorInstuction{
			FromStackIdx: 12,
			ToStackIdx:   14,
			NumElements:  17,
		},
	},
}

func Test_parseInstruction(t *testing.T) {
	for _, test := range parseInstructionTests {
		instruction := parseInstruction(test.line)
		if !reflect.DeepEqual(instruction, test.want) {
			t.Errorf("failed to parse operator instruction '%s', wanted %+v, got %+v", test.line, test.want, instruction)
		}
	}
}

type executeInstructionTest struct {
	instruction   operatorInstuction
	initialStacks [][]string
	want          [][]string
}

func Test_executePart1Instruction(t *testing.T) {
	test := executeInstructionTest{
		instruction: operatorInstuction{
			FromStackIdx: 1,
			ToStackIdx:   0,
			NumElements:  1,
		},
		initialStacks: [][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		},
		want: [][]string{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
	}
	stacks := test.initialStacks
	executePart1Instruction(test.instruction, stacks)
	if !reflect.DeepEqual(stacks, test.want) {
		t.Errorf("final stack does not match: %+v != %+v", stacks, test.want)
	}
}

func Test_executePart2Instruction(t *testing.T) {
	test := executeInstructionTest{
		instruction: operatorInstuction{
			FromStackIdx: 0,
			ToStackIdx:   2,
			NumElements:  3,
		},
		initialStacks: [][]string{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
		want: [][]string{
			{},
			{"M", "C"},
			{"P", "Z", "N", "D"},
		},
	}
	stacks := test.initialStacks
	executePart2Instruction(test.instruction, stacks)
	if !reflect.DeepEqual(stacks, test.want) {
		t.Errorf("final stack does not match: %+v != %+v", stacks, test.want)
	}
}
