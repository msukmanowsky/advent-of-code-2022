package day3

import (
	"bufio"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

// Used python to generate this ;)
var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func setFromString(s string) map[rune]bool {
	m := map[rune]bool{}
	for _, r := range s {
		m[r] = true
	}
	return m
}

func setIntersection(m1, m2 map[rune]bool) map[rune]bool {
	sIntersection := map[rune]bool{}
	if len(m1) > len(m2) {
		m1, m2 = m2, m1 // better to iterate over a shorter set
	}
	for k, _ := range m1 {
		if m2[k] {
			sIntersection[k] = true
		}
	}
	return sIntersection
}

type result struct {
	Compartment1            string
	Compartment2            string
	SharedTypes             []string
	TotalSharedTypePriority int
}

func Day3() {
	parseLine := func(line string) result {
		lineLen := len(line)
		splitPoint := lineLen / 2
		compartment1 := line[:splitPoint]
		compartment1Set := setFromString(compartment1)
		compartment2 := line[splitPoint:]
		compartment2Set := setFromString(compartment2)
		intersection := setIntersection(compartment1Set, compartment2Set)
		sharedTypes := []string{}
		totalSharedTypePriority := 0
		for k, _ := range intersection {
			ks := string(k)
			priority, ok := priorities[ks]
			if !ok {
				log.Fatalln("unhandled priority: ", k)
			}
			sharedTypes = append(sharedTypes, ks)
			totalSharedTypePriority += priority
		}

		return result{
			Compartment1:            compartment1,
			Compartment2:            compartment2,
			SharedTypes:             sharedTypes,
			TotalSharedTypePriority: totalSharedTypePriority,
		}
	}

	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSharedTypePriority := 0
	for scanner.Scan() {
		result := parseLine(scanner.Text())
		totalSharedTypePriority += result.TotalSharedTypePriority
	}
	log.Printf("Total Shared Type Priority: %d", totalSharedTypePriority)
}
