package day6

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

type consumeResult struct {
	charsRead  int
	window     []rune
	matchFound bool
}

func consumeUntilUniqueWindowFound(reader *bufio.Reader, windowSize int) (*consumeResult, error) {
	result := consumeResult{
		matchFound: false,
	}
	window := make([]rune, windowSize)
	charsRead := 0

	// initialize the window
	for i := 0; i < windowSize; i++ {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("not enough runes in input to initialize window of size: %d", windowSize)
			} else {
				return nil, err
			}
		}
		window[i] = r
		charsRead += 1
	}

	windowSet := windowToSet(window)
	if len(windowSet) == windowSize {
		result.charsRead = charsRead
		result.window = window
		result.matchFound = true
		return &result, nil
	}

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		charsRead += 1
		// shift left by 1
		window = append(window[1:], r)
		windowSet = windowToSet(window)
		if len(windowSet) == windowSize {
			result.charsRead = charsRead
			result.window = window
			result.matchFound = true
			break
		}
	}

	return &result, nil
}

func windowToSet(window []rune) map[rune]bool {
	m := map[rune]bool{}
	for _, r := range window {
		m[r] = true
	}
	return m
}

func Day6_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	windowSize := 4
	reader := bufio.NewReader(file)

	result, err := consumeUntilUniqueWindowFound(reader, windowSize)
	if err != nil {
		log.Fatal(err)
	}

	if !result.matchFound {
		fmt.Printf("no unique %d-length sequence found in input", windowSize)
		return
	}

	fmt.Printf("consumed %d chars to find first non-repeating %d-length sequence %+v\n", result.charsRead, windowSize, string(result.window))
}

func Day6_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	windowSize := 14
	reader := bufio.NewReader(file)

	result, err := consumeUntilUniqueWindowFound(reader, windowSize)
	if err != nil {
		log.Fatal(err)
	}

	if !result.matchFound {
		fmt.Printf("no unique %d-length sequence found in input", windowSize)
		return
	}

	fmt.Printf("consumed %d chars to find first non-repeating %d-length sequence %+v\n", result.charsRead, windowSize, string(result.window))
}
