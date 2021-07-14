package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
)

func (a *Algorithm) longestPath() bool {
	if pathFound := a.aStarSearch(); !pathFound {
		return false
	}

	cells := a.initGrid()

	for _, c := range a.solvedPath {
		cells[c].IsVisited = true
	}

	index, current := 0, a.head
	w := a.board.Width
	h := a.board.Height
	for {
		dirToNext := coord.Diff(&a.solvedPath[index], &current)
		next := a.solvedPath[index]

		d := directionToIndex[dirToNext]
		var tests [2]int
		if d == LEFT || d == RIGHT {
			tests = [2]int{UP, DOWN}
		} else if d == UP || d == DOWN {
			tests = [2]int{LEFT, RIGHT}
		}

		extended := false
		for _, test := range tests {
			testDir := indexToDirection[test]
			currentTestCoord := coord.Sum(&current, &testDir)
			nextTestCoord := coord.Sum(&next, &testDir)

			if !coord.IsOutside(&currentTestCoord, w, h) && !coord.IsOutside(&nextTestCoord, w, h) {
				currentCell := cells[currentTestCoord]
				nextCell := cells[nextTestCoord]

				if a.isOkAndNotVisited(currentCell) && a.isOkAndNotVisited(nextCell) {
					currentCell.IsVisited = true
					nextCell.IsVisited = true

					a.solvedPath = insert(a.solvedPath, index, currentTestCoord)
					a.solvedPath = insert(a.solvedPath, index+1, nextTestCoord)

					extended = true
					break
				}
			}
		}

		if !extended {
			current = next
			index++
			if index >= len(a.solvedPath) {
				break
			}
		}
	}

	return true
}
