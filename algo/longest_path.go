package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

func (a *Algorithm) longestPath() []*cell.Cell {
	if pathFound := a.aStarSearch(); !pathFound {
		a.solvedPath.Clear()
		return make([]*cell.Cell, 0)
	}

	cells := a.initGrid()
	longestPath := make([]*cell.Cell, a.solvedPath.Len())

	for i := 0; a.solvedPath.Len() != 0; i++ {
		c := a.solvedPath.Pop().(models.Coord)
		longestPath[i] = cells[c]
		cells[c].IsVisited = true
	}

	index, current := 0, a.head
	w := a.board.Width
	h := a.board.Height
	for {
		dirToNext := coord.Diff(longestPath[index], &current)
		next := longestPath[index].Coord

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
			currentTest := coord.Sum(&current, &testDir)
			nextTest := coord.Sum(&next, &testDir)

			if !coord.IsOutside(&currentTest, w, h) && !coord.IsOutside(&nextTest, w, h) {
				currentCell := cells[currentTest]
				nextCell := cells[nextTest]

				if a.isOkAndNotVisited(currentCell) && a.isOkAndNotVisited(nextCell) {
					currentCell.IsVisited = true
					nextCell.IsVisited = true

					longestPath = insert(longestPath, index, currentCell)
					longestPath = insert(longestPath, index+1, nextCell)

					extended = true
					break
				}
			}
		}

		if !extended {
			current = next
			index++
			if index >= len(longestPath) {
				break
			}
		}
	}

	return longestPath
}
