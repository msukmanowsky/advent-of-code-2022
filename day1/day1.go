package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

func Day1() {
	// https://adventofcode.com/2022/day/1
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	caloriesByElf := make([]int, 1)
	elfIdx := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		switch text {
		case "":
			elfIdx += 1
			caloriesByElf = append(caloriesByElf, 0)
		default:
			calories, _ := strconv.Atoi(text)
			caloriesByElf[elfIdx] += calories
		}
	}

	maxValue := -1
	maxIdx := -1
	for idx, val := range caloriesByElf {
		if val > maxValue {
			maxValue = val
			maxIdx = idx
		}
	}

	log.Printf("Elf at index %d has the most calories with %d calories", maxIdx, maxValue)
}
