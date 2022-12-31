package day10

import (
	"bufio"
	"strings"
	"testing"
)

var sampleInput = strings.TrimSpace(`
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`)

// cycle num: registerX value
var sampleInputCycles = map[int]int{
	20:  21,
	60:  19,
	100: 18,
	140: 21,
	180: 16,
	220: 18,
}

func Test_executeCommandNoop(t *testing.T) {
	cpu := CPU{}
	cpu.executeCommand("noop", nil, nil)
	if cpu.Cycle != 1 {
		t.Fatalf("expected noop to take 1 cycle, took %d", cpu.Cycle)
	}
}

func Test_executeCommandAddX(t *testing.T) {
	cpu := CPU{}
	cpu.executeCommand("addx 10", nil, nil)
	if cpu.Cycle != 2 {
		t.Fatalf("expected addx to take 2 cycles, took %d", cpu.Cycle)
	}
	if cpu.RegisterX != 10 {
		t.Fatalf("expected registerX to be 10, got: %d", cpu.RegisterX)
	}
}

func Test_sample(t *testing.T) {
	cpu := NewCPU()
	reader := bufio.NewReader(strings.NewReader(sampleInput))
	scanner := bufio.NewScanner(reader)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		cpu.executeCommand(line, func(cpu CPU) {
			if cpu.Cycle == 20 && cpu.RegisterX != 21 {
				t.Fatalf("expected registerX to be 21 on cycle 20, got: %d", cpu.RegisterX)
				return
			}

			wantRegisterX, cycleExists := sampleInputCycles[cpu.Cycle]
			if !cycleExists {
				return
			}
			if wantRegisterX != cpu.RegisterX {
				t.Logf("lineNo: %d, cpu: %+v", i, cpu)
				t.Fatalf("want registerX to be %d on cycle %d, got %d", wantRegisterX, cpu.Cycle, cpu.RegisterX)
			}
		}, nil)
		i += 1
	}
}
