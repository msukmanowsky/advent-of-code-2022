package day8

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/utils"
)

func readInput(r io.Reader) (elements [][]int) {
	elements = [][]int{}

	scanner := bufio.NewScanner(r)
	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for colIdx, rune_ := range line {
			cell := int(rune_ - '0')
			row[colIdx] = cell
		}
		elements = append(elements, row)
		rowIdx += 1
	}

	return elements
}

// How many trees are visible from outside the grid?
func Day8_1() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	elements := readInput(file)
	numRows := len(elements)
	numCols := len(elements[0])
	numVisible := 0
	for rowIdx, row := range elements {
		if rowIdx == 0 || rowIdx == numRows-1 {
			// Shortcut, we're on the top or bottom side of the perimiter so we are
			// visible by default
			numVisible += numCols
			continue
		}
		for colIdx, cell := range row {
			if colIdx == 0 || colIdx == numCols-1 {
				// Shortcut, we're on the left or right side of the perimiter so we are
				// visible by default
				numVisible += 1
				continue
			}

			// For each cell we "look" top, left, bottom and right which isn't the
			// most efficient way to do this. One thought for an improvement, as you
			// scan successive ranges, you can keep track of the max value in that
			// range. Benefit would be fewer scans and comparisons needed as you go
			// along.

			// Look up first. We should do this in reverse order, but since we need
			// to scan all rows, it doesn't matter
			isVisible := true
			for _, topRow := range elements[:rowIdx] {
				topCell := topRow[colIdx]
				if cell <= topCell {
					// Not taller, bail out
					isVisible = false
					break
				}
			}
			if isVisible {
				// Taller than all in top, move on
				numVisible += 1
				continue
			}

			// Not taller from top check the left
			isVisible = true
			for _, leftCell := range elements[rowIdx][:colIdx] {
				if cell <= leftCell {
					// Not taller, bail out
					isVisible = false
					break
				}
			}
			if isVisible {
				// Taller than all in right, move on
				numVisible += 1
				continue
			}

			// Not taller from top or right, check the bottom
			isVisible = true
			for _, bottomRow := range elements[rowIdx+1:] {
				bottomCell := bottomRow[colIdx]
				if cell <= bottomCell {
					isVisible = false
					break
				}
			}
			if isVisible {
				// Taller than all in bottom, move on
				numVisible += 1
				continue
			}

			// Not taller from the top, left or bottom, check the right. We should
			// do this in reverse order, but since we need to scan the full row, it
			// doesn't matter
			isVisible = true
			for _, rightCell := range elements[rowIdx][colIdx+1:] {
				if cell <= rightCell {
					// Not taller, bail out
					isVisible = false
					break
				}
			}
			if isVisible {
				// Taller than all in right, move on
				numVisible += 1
				continue
			}
		}
	}
	log.Printf("number of visible trees: %d\n", numVisible)
}

func Day8_2() {
	file, err := os.Open(utils.ModRelativeFilePath("input.txt"))
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	elements := readInput(file)
	maxSenicScore := -1

	for rowIdx, row := range elements {
		for colIdx, treeHeight := range row {
			// Look up, in reverse order
			topVisible := 0
			topRows := elements[:rowIdx]
			for i := len(topRows) - 1; i >= 0; i-- {
				topVisible += 1
				topRow := topRows[i]
				otherTreeHeight := topRow[colIdx]
				if otherTreeHeight >= treeHeight {
					break
				}
			}

			// Look down
			bottomVisible := 0
			bottomRows := elements[rowIdx+1:]
			for _, bottomRow := range bottomRows {
				bottomVisible += 1
				otherTreeHeight := bottomRow[colIdx]
				if otherTreeHeight >= treeHeight {
					break
				}
			}

			// Look left, again need to do this in reverse order
			left := elements[rowIdx][:colIdx]
			leftVisible := 0
			for i := len(left) - 1; i >= 0; i-- {
				leftVisible += 1
				otherTreeHeight := left[i]
				if otherTreeHeight >= treeHeight {
					break
				}
			}

			// Look right
			rightVisible := 0
			right := elements[rowIdx][colIdx+1:]
			for _, otherTreeHeight := range right {
				rightVisible += 1
				if otherTreeHeight >= treeHeight {
					break
				}
			}

			senicScore := topVisible * bottomVisible * rightVisible * leftVisible
			if senicScore > maxSenicScore {
				maxSenicScore = senicScore
			}

		}
	}
	log.Printf("max senic score is: %d\n", maxSenicScore)
}
