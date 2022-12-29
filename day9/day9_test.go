package day9

import (
	"reflect"
	"testing"
)

type executeCommandTest struct {
	name              string
	head              Position
	tail              Position
	command           Command
	wantHead          Position
	wantTail          Position
	wantTailPositions map[Position]bool
}

var commandTests = []executeCommandTest{
	// When H and T overlap
	{
		name:     "overlap",
		head:     Position{X: 0, Y: 0},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "U", Magnitude: 1},
		wantHead: Position{X: 0, Y: 1},
		wantTail: Position{X: 0, Y: 0},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}: true,
		},
	},
	// Simple moves
	{
		name:     "left - 1",
		head:     Position{X: -1, Y: 0},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "L", Magnitude: 1},
		wantHead: Position{X: -2, Y: 0},
		wantTail: Position{X: -1, Y: 0},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}:  true,
			{X: -1, Y: 0}: true,
		},
	},
	{
		name:     "right - 1",
		head:     Position{X: 1, Y: 0},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "R", Magnitude: 1},
		wantHead: Position{X: 2, Y: 0},
		wantTail: Position{X: 1, Y: 0},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
		},
	},
	{
		name:     "up - 1",
		head:     Position{X: 0, Y: 1},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "U", Magnitude: 1},
		wantHead: Position{X: 0, Y: 2},
		wantTail: Position{X: 0, Y: 1},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}: true,
			{X: 0, Y: 1}: true,
		},
	},
	{
		name:     "down - 1",
		head:     Position{X: 0, Y: -1},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "D", Magnitude: 1},
		wantHead: Position{X: 0, Y: -2},
		wantTail: Position{X: 0, Y: -1},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}:  true,
			{X: 0, Y: -1}: true,
		},
	},
	{
		name:     "example 1",
		head:     Position{X: 0, Y: 0},
		tail:     Position{X: 0, Y: 0},
		command:  Command{Direction: "R", Magnitude: 4},
		wantHead: Position{X: 4, Y: 0},
		wantTail: Position{X: 3, Y: 0},
		wantTailPositions: map[Position]bool{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
		},
	},
	{
		name:     "example 2",
		head:     Position{X: 4, Y: 0},
		tail:     Position{X: 3, Y: 0},
		command:  Command{Direction: "U", Magnitude: 4},
		wantHead: Position{X: 4, Y: 4},
		wantTail: Position{X: 4, Y: 3},
		wantTailPositions: map[Position]bool{
			{X: 3, Y: 0}: true,
			{X: 4, Y: 1}: true,
			{X: 4, Y: 2}: true,
			{X: 4, Y: 3}: true,
		},
	},
	{
		name:     "example 3",
		head:     Position{X: 4, Y: 4},
		tail:     Position{X: 4, Y: 3},
		command:  Command{Direction: "L", Magnitude: 3},
		wantHead: Position{X: 1, Y: 4},
		wantTail: Position{X: 2, Y: 4},
		wantTailPositions: map[Position]bool{
			{X: 4, Y: 3}: true,
			{X: 3, Y: 4}: true,
			{X: 2, Y: 4}: true,
		},
	},
}

func Test_executeCommand(t *testing.T) {
	for _, test := range commandTests {
		t.Logf("test: %s", test.name)
		gotHead, gotTail, gotTailPositions := executeCommand(test.head, test.tail, test.command)
		if gotHead != test.wantHead {
			t.Fatalf("wanted head %+v, got %+v", test.wantHead, gotHead)
		}
		if gotTail != test.wantTail {
			t.Fatalf("wanted tail %+v, got %+v", test.wantTail, gotTail)
		}
		if !reflect.DeepEqual(gotTailPositions, test.wantTailPositions) {
			t.Fatalf("wanted tailPositions %+v, got %+v", test.wantTailPositions, gotTailPositions)
		}
	}
}
