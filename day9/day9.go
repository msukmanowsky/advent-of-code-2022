package day9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	// x >= y
	return y
}

type Position struct {
	X, Y int
}

type Command struct {
	Direction string
	Magnitude int
}

func (pos *Position) step(cmd Command) {
	switch cmd.Direction {
	case "U":
		pos.Y += 1
	case "D":
		pos.Y -= 1
	case "R":
		pos.X += 1
	case "L":
		pos.X -= 1
	}
}

func parseCommand(line string) (*Command, error) {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid command: %s", line)
	}

	magnitude, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return nil, err
	}

	direction := parts[0]
	if direction != "U" && direction != "D" && direction != "L" && direction != "R" {
		return nil, fmt.Errorf("invalid command: %s", line)
	}

	return &Command{Direction: direction, Magnitude: int(magnitude)}, nil
}

func executeCommand(head Position, tail Position, cmd Command) (Position, Position, map[Position]bool) {
	tailPositions := map[Position]bool{
		tail: true,
	}
	for i := 0; i < cmd.Magnitude; i++ {
		// log.Printf("executing step %s (%d / %d)", cmd.Direction, i+1, cmd.Magnitude)
		// log.Printf("  head is at %+v, tail is at %+v", head, tail)
		head.step(cmd)
		// log.Printf("  moved head to: %+v", head)
		// If positive, head is RIGHT of tail
		// If negative, head is LEFT of tail
		xDist := head.X - tail.X
		// If positive, head is ABOVE tail
		// If negative, head is BELOW tail
		yDist := head.Y - tail.Y
		// Euclidean distance
		distance := math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
		isTouching := distance < 2.0
		if isTouching {
			// log.Printf("  head and tail are touching")
			continue
		}

		xMove := xDist
		if xMove == -2 {
			xMove = -1
		} else if xMove == 2 {
			xMove = 1
		}

		yMove := yDist
		if yMove == -2 {
			yMove = -1
		} else if yMove == 2 {
			yMove = 1
		}
		// log.Printf("  xDist=%d, yDist=%d moving tail x by %d and y by %d", xDist, yDist, xMove, yMove)
		tail.X += xMove
		tail.Y += yMove
		tailPositions[tail] = true
	}
	return head, tail, tailPositions
}

func executeCommandPart2(rope []Position, cmd Command) map[Position]bool {
	ropeLen := len(rope)
	head := &rope[ropeLen-1]
	tail := &rope[0]
	tailPositions := map[Position]bool{
		*tail: true,
	}

	// log.Printf("executing command: %+v", cmd)
	for i := 0; i < cmd.Magnitude; i++ {
		// first move the head
		head.step(cmd)
		// log.Printf("head moved %s1 to %+v", cmd.Direction, head)
		// now move the rest of the rope segments
		for j := ropeLen - 2; j >= 0; j-- {
			front := rope[j+1]
			back := &rope[j]
			xDist := front.X - back.X
			yDist := front.Y - back.Y
			distance := math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
			isTouching := distance < 2.0
			if isTouching {
				continue
			}

			xMove := xDist
			if xMove == -2 {
				xMove = -1
			} else if xMove == 2 {
				xMove = 1
			}

			yMove := yDist
			if yMove == -2 {
				yMove = -1
			} else if yMove == 2 {
				yMove = 1
			}
			back.X += xMove
			back.Y += yMove
		}
		// log.Printf("rope: %+v", rope)
		tailPositions[*tail] = true
	}
	return tailPositions
}

func Day9_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	head := Position{X: 0, Y: 0}
	tail := Position{X: 0, Y: 0}
	tailPositions := map[Position]bool{
		tail: true,
	}
	var newTailPositions map[Position]bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		headCommand, err := parseCommand(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		head, tail, newTailPositions = executeCommand(head, tail, *headCommand)
		for k, v := range newTailPositions {
			tailPositions[k] = v
		}
	}
	log.Printf("tail was moved to %d unique locations", len(tailPositions))
}

func Day9_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	numKnots := 10
	rope := make([]Position, numKnots)
	tailPositions := map[Position]bool{
		{X: 0, Y: 0}: true,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		headCommand, err := parseCommand(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		newTailPositions := executeCommandPart2(rope, *headCommand)
		for k, v := range newTailPositions {
			tailPositions[k] = v
		}
	}
	log.Printf("tail was moved to %d unique locations", len(tailPositions))
}
