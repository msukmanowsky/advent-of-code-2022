package day2

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

type move struct {
	Name   string
	Points int
}

var rock = move{
	Name:   "rock",
	Points: 1,
}
var paper = move{
	Name:   "paper",
	Points: 2,
}
var scissors = move{
	Name:   "scissors",
	Points: 3,
}

type roundResult string

const (
	win        roundResult = "win"
	loss       roundResult = "loss"
	draw       roundResult = "draw"
	winPoints              = 6
	lossPoints             = 0
	drawPoints             = 3
)

type roundSummary struct {
	Result       roundResult
	PlayerMove   move
	OpponentMove move
	PlayerPoints int
}

// Day 2, Part 1 https://adventofcode.com/2022/day/2
func Day2_1() {
	moveStrToMove := map[string]move{
		// Opponent moves
		"A": rock,
		"B": paper,
		"C": scissors,
		// Player moves
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	parseRound := func(input string) roundSummary {
		moves := strings.SplitN(input, " ", 2)
		if len(moves) != 2 {
			log.Fatalln("expected 2 moves, found ", len(moves))
		}
		opponentMove, ok := moveStrToMove[moves[0]]
		if !ok {
			log.Fatalln("invalid opponent move found: ", input)
		}
		playerMove, ok := moveStrToMove[moves[1]]
		if !ok {
			log.Fatalln("invalid player move found: ", input)
		}

		if opponentMove == playerMove {
			// Draw
			return roundSummary{
				Result:       draw,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: drawPoints + playerMove.Points,
			}
		}

		// Moves are different
		switch true {
		case (playerMove == rock && opponentMove == scissors):
			return roundSummary{
				Result:       win,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: winPoints + playerMove.Points,
			}
		case (playerMove == paper && opponentMove == rock):
			return roundSummary{
				Result:       win,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: winPoints + playerMove.Points,
			}
		case (playerMove == scissors && opponentMove == paper):
			return roundSummary{
				Result:       win,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: winPoints + playerMove.Points,
			}
		case (playerMove == rock && opponentMove == paper):
			return roundSummary{
				Result:       loss,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: lossPoints + playerMove.Points,
			}
		case (playerMove == paper && opponentMove == scissors):
			return roundSummary{
				Result:       loss,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: lossPoints + playerMove.Points,
			}
		case (playerMove == scissors && opponentMove == rock):
			return roundSummary{
				Result:       loss,
				PlayerMove:   playerMove,
				OpponentMove: opponentMove,
				PlayerPoints: lossPoints + playerMove.Points,
			}
		default:
			log.Fatalln("unhandled move: ", input)
			return roundSummary{}
		}
	}

	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		roundSummary := parseRound(text)
		// log.Printf("%s", text)
		// logRoundSummary(roundSummary)
		totalScore += roundSummary.PlayerPoints
	}
	log.Printf("Total player score is: %d", totalScore)
}

// Day 2, Part 2 https://adventofcode.com/2022/day/2
func Day2_2() {
	rawOpponentMoveToMove := map[string]move{
		"A": rock,
		"B": paper,
		"C": scissors,
	}
	rawRoundResultToRoundResult := map[string]roundResult{
		"X": loss,
		"Y": draw,
		"Z": win,
	}

	parseRound := func(input string) roundSummary {
		moves := strings.SplitN(input, " ", 2)
		if len(moves) != 2 {
			log.Fatalln("expected 2 moves, found ", len(moves))
		}
		opponentMove, ok := rawOpponentMoveToMove[moves[0]]
		if !ok {
			log.Fatalln("invalid opponent move found: ", input)
		}
		roundResult, ok := rawRoundResultToRoundResult[moves[1]]
		if !ok {
			log.Fatalln("invalid round result found: ", input)
		}

		if roundResult == draw {
			return roundSummary{
				Result:       draw,
				OpponentMove: opponentMove,
				PlayerMove:   opponentMove,
				PlayerPoints: drawPoints + opponentMove.Points,
			}
		}

		switch true {
		// Losses
		case (roundResult == loss && opponentMove == rock):
			return roundSummary{
				Result:       loss,
				OpponentMove: opponentMove,
				PlayerMove:   scissors,
				PlayerPoints: lossPoints + scissors.Points,
			}
		case (roundResult == loss && opponentMove == paper):
			return roundSummary{
				Result:       loss,
				OpponentMove: opponentMove,
				PlayerMove:   rock,
				PlayerPoints: lossPoints + rock.Points,
			}
		case (roundResult == loss && opponentMove == scissors):
			return roundSummary{
				Result:       loss,
				OpponentMove: opponentMove,
				PlayerMove:   paper,
				PlayerPoints: lossPoints + paper.Points,
			}

		case (roundResult == win && opponentMove == rock):
			return roundSummary{
				Result:       win,
				OpponentMove: opponentMove,
				PlayerMove:   paper,
				PlayerPoints: winPoints + paper.Points,
			}
		case (roundResult == win && opponentMove == paper):
			return roundSummary{
				Result:       win,
				OpponentMove: opponentMove,
				PlayerMove:   scissors,
				PlayerPoints: winPoints + scissors.Points,
			}
		case (roundResult == win && opponentMove == scissors):
			return roundSummary{
				Result:       win,
				OpponentMove: opponentMove,
				PlayerMove:   rock,
				PlayerPoints: winPoints + rock.Points,
			}
		default:
			log.Fatalln("unhandled case: ", input)
			return roundSummary{}
		}
	}

	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		roundSummary := parseRound(text)
		totalScore += roundSummary.PlayerPoints
	}
	log.Printf("Total player score is: %d", totalScore)
}
