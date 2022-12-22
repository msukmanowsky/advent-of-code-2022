package day6

import (
	"bufio"
	"strings"
	"testing"
)

func Test_consumeUntilUniqueWindowFound(t *testing.T) {
	windowSize := 4
	reader := bufio.NewReader(strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	result, err := consumeUntilUniqueWindowFound(reader, windowSize)
	if err != nil {
		t.Fatalf("error: %+v", err)
	}
	if result.matchFound != true {
		t.Fatalf("match not found")
	}
	if result.charsRead != 7 {
		t.Fatalf("should've read 7 runes, but read %d", result.charsRead)
	}
	if string(result.window) != "jpqm" {
		t.Fatalf("didn't produce the right window: %s", string(result.window))
	}

}
