package algo

import "github.com/ISKalsi/boomba-the-sapera/algo/cell"

func (a *Algorithm) longestPath() (bool, float64) {
	pathFound := false
	pathCost := 0.0

	if pathFound, pathCost = a.aStarSearch(); !pathFound {
		return false, -1
	}

	cells := a.initGrid()

	for _, c := range a.solvedPath {
		cells[c].IsVisited = true
	}

	index, current := 0, a.head
	w := a.board.Width
	h := a.board.Height
	for pathCost+2 < a.health {
		dirToNext := a.solvedPath[index].Diff(current)
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
			currentTestCoord := current.Sum(testDir)
			nextTestCoord := next.Sum(testDir)

			if !currentTestCoord.IsOutside(w, h) && !nextTestCoord.IsOutside(w, h) {
				currentTestCell := cells[currentTestCoord]
				nextTestCell := cells[nextTestCoord]

				if currentTestCell.IsOkAndNotVisited() && nextTestCell.IsOkAndNotVisited() {
					if currentTestCell.Weight != cell.WeightHazard && nextTestCell.Weight != cell.WeightHazard {
						currentTestCell.IsVisited = true
						nextTestCell.IsVisited = true

						a.solvedPath = insert(a.solvedPath, index, currentTestCoord)
						a.solvedPath = insert(a.solvedPath, index+1, nextTestCoord)

						pathCost += currentTestCell.Weight
						pathCost += nextTestCell.Weight

						extended = true
						break
					}
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

	return true, pathCost
}
