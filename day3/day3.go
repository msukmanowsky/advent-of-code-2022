package day3

import (
	"bufio"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

func getPriorities() map[rune]int {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorities := make(map[rune]int)
	for i, r := range letters {
		priorities[r] = i + 1
	}
	return priorities
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

func getSetElements(m map[rune]bool) []rune {
	keys := make([]rune, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i += 1
	}
	return keys
}

type result struct {
	Line                    string
	Compartment1            string
	Compartment2            string
	SharedTypes             []string
	TotalSharedTypePriority int
}

func parseLine(line string, priorities map[rune]int) result {
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
		priority, ok := priorities[k]
		if !ok {
			log.Fatalln("unhandled priority: ", k)
		}
		sharedTypes = append(sharedTypes, string(k))
		totalSharedTypePriority += priority
	}

	return result{
		Line:                    line,
		Compartment1:            compartment1,
		Compartment2:            compartment2,
		SharedTypes:             sharedTypes,
		TotalSharedTypePriority: totalSharedTypePriority,
	}
}

func Day3_1() {
	priorities := getPriorities()

	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSharedTypePriority := 0
	for scanner.Scan() {
		result := parseLine(scanner.Text(), priorities)
		totalSharedTypePriority += result.TotalSharedTypePriority
	}
	log.Printf("Total Shared Type Priority: %d", totalSharedTypePriority)
}

func Day3_2() {
	priorities := getPriorities()
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	group := []string{}
	totalSharedTypePriority := 0
	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, line)
		if len(group) == 3 {
			s1 := setFromString(group[0])
			s2 := setFromString(group[1])
			s3 := setFromString(group[2])
			intersection := setIntersection(s1, s2)
			intersection = setIntersection(intersection, s3)
			if len(intersection) != 1 {
				log.Fatalln("badge parsing failed for group: ", group)
			}
			badge := getSetElements(intersection)[0]
			priority := priorities[badge]
			totalSharedTypePriority += priority
			group = nil
		}
	}
	if len(group) != 0 {
		log.Fatalln("unresolved group", group)
	}

	log.Printf("Total Shared Type Priority: %d", totalSharedTypePriority)
}
