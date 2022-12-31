package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

type CPU struct {
	InternalCommand string
	Command         string
	CommandArgs     []string
	Cycle           int
	RegisterX       int
}

func NewCPU() *CPU {
	cpu := CPU{Cycle: 1, RegisterX: 1}
	return &cpu
}

func (cpu *CPU) executeCommand(rawCmd string, onCycleStart func(cpu CPU), onCycleEnd func(cpu CPU)) {
	parts := strings.Split(rawCmd, " ")
	cpu.Command, cpu.CommandArgs = parts[0], parts[1:]
	var internalCommands []string
	switch cpu.Command {
	case "noop":
		internalCommands = []string{"noop"}
	case "addx":
		internalCommands = []string{"addx_1", "addx_2"}
	}

	for _, cpu.InternalCommand = range internalCommands {
		if onCycleStart != nil {
			onCycleStart(*cpu)
		}

		cpu.Cycle += 1
		switch cpu.InternalCommand {
		case "addx_2":
			v, err := strconv.ParseInt(cpu.CommandArgs[0], 10, 64)
			if err != nil {
				log.Fatalf("bad command: %s", rawCmd)
			}
			cpu.RegisterX += int(v)
		}

		if onCycleEnd != nil {
			onCycleEnd(*cpu)
		}
	}
	cpu.Command = ""
	cpu.CommandArgs = nil
	cpu.InternalCommand = ""
}

func Day10_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cpu := NewCPU()
	signalStrengths := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		cpu.executeCommand(line, func(cpu CPU) {
			switch cpu.Cycle {
			case 20, 60, 100, 140, 180, 220:
				signalStrengths = append(signalStrengths, cpu.Cycle*cpu.RegisterX)
			}
		}, nil)
	}

	sum := 0
	for _, signalStength := range signalStrengths {
		sum += signalStength
	}
	log.Printf("sum of signal strengths: %d", sum)
}

const SpriteWidth = 3
const ScreenWidth = 40
const ScreenHeight = 6
const NumPixels = ScreenWidth * ScreenHeight

func SPrintScreen(screen string) string {
	sb := strings.Builder{}
	for i, c := range screen {
		sb.WriteRune(c)
		if i > 0 && (i+1)%ScreenWidth == 0 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func Day10_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Register X controls mid horizontal position of the sprite.
	// This CRT screen draws the top row of pixels left-to-right, then the row
	// below that, and so on. The left-most pixel in each row is in position 0,
	// and the right-most pixel in each row is in position 39.
	// If the sprite is positioned such that one of its three pixels is the pixel
	// currently being drawn, the screen produces a lit pixel (#); otherwise, the
	// screen leaves the pixel dark (.).
	cpu := NewCPU()
	screen := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		cpu.executeCommand(line, func(cpu CPU) {

			// RegisterX is 0-indexed, but cycles are 1-indexed, so we'll make a
			// correction
			cycle := (cpu.Cycle - 1) % ScreenWidth

			// Determine if the current cycle is within the bounds of the sprite
			if cycle >= cpu.RegisterX-1 && cycle <= cpu.RegisterX+1 {
				screen.WriteString("#")
			} else {
				screen.WriteString(".")
			}
		}, nil)
	}
	log.Printf("screen\n%s", SPrintScreen(screen.String()))
}
