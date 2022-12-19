package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

type assignment struct {
	Min uint16
	Max uint16
}

// Does a contain other?
func (a assignment) contains(other assignment) bool {
	return a.Min <= other.Min && a.Max >= other.Max
}

// Does a overlap at all with other?
func (a assignment) overlaps(other assignment) bool {
	var min, max assignment
	if a.Min < other.Min {
		min = a
		max = other
	} else {
		min = other
		max = a
	}
	if max.Min <= min.Max {
		return true
	}
	return false
}

func parseLine(line string) []assignment {
	assignments := make([]assignment, 2)
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		log.Fatalln("invalid line: ", line)
	}

	for i, part := range parts {
		assignmentParts := strings.Split(part, "-")
		if len(assignmentParts) != 2 {
			log.Fatalln("invalid assignment in line: ", line)
		}
		min, err := strconv.ParseUint(assignmentParts[0], 10, 16)
		if err != nil {
			log.Fatalln("invalid assignment in line: ", line)
		}
		max, err := strconv.ParseUint(assignmentParts[1], 10, 16)
		if err != nil {
			log.Fatalln("invalid assignment in line: ", line)
		}
		assignments[i] = assignment{
			Min: uint16(min),
			Max: uint16(max),
		}
	}

	return assignments
}

func Day4_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := parseLine(line)
		a := assignments[0]
		b := assignments[1]
		if a.contains(b) || b.contains(a) {
			count += 1
		}
	}
	log.Printf("Total number of pairs with fully contained segments: %d", count)
}

func Day4_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		assignments := parseLine(line)
		a := assignments[0]
		b := assignments[1]
		if a.overlaps(b) {
			count += 1
		}
	}
	log.Printf("Total number of pairs with overlapping segments: %d", count)
}
