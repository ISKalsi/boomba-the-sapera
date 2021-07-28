package algo

import (
	"container/heap"
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/grid"
	"github.com/ISKalsi/boomba-the-sapera/ds/pque"
)

func (a *Algorithm) initCellsToVisitList(cells grid.Grid) *pque.PriorityQueue {
	start := cells[a.start]
	start.F, start.G, start.H = 0, 0, 0
	start.Coord = a.start
	start.ParentCoord = a.start

	cellsToVisit := &pque.PriorityQueue{}
	heap.Push(cellsToVisit, start)

	return cellsToVisit
}

func (a *Algorithm) tracePath(g grid.Grid) {
	a.clearSolvedPath()

	current := a.destination
	parent := g[current].ParentCoord

	for current != parent {
		a.solvedPath = append(a.solvedPath, current)
		current = parent
		parent = g[current].ParentCoord
	}

	reverse(a.solvedPath)
}

func (a *Algorithm) resetBlockTailFlag() {
	if a.dontBlockTailOrHead {
		a.dontBlockTailOrHead = false
	}
}

func (a *Algorithm) aStarSearch() (bool, float64) {
	defer a.resetBlockTailFlag()

	if a.start == a.destination {
		return false, -1
	}

	cells := a.initGrid()
	if !a.dontBlockTailOrHead && !cells[a.destination].IsOk() {
		return false, -1
	}

	cellsToVisit := a.initCellsToVisitList(cells)
	for cellsToVisit.Len() != 0 {
		currentCell := heap.Pop(cellsToVisit).(*cell.Cell)
		currentCell.IsVisited = true

		for dir := range directionToIndex {
			neighborCoord := currentCell.Coord.Sum(dir)

			if !neighborCoord.IsOutside(a.board.Width, a.board.Height) {
				neighborCell := cells[neighborCoord]

				if neighborCoord == a.destination {
					neighborCell.ParentCoord = currentCell.Coord
					g := currentCell.G + neighborCell.Weight

					if g < a.health {
						a.tracePath(cells)
						return true, g
					} else {
						return false, -1
					}
				} else if neighborCell.IsOkAndNotVisited() {
					g := currentCell.G + neighborCell.Weight
					h := neighborCoord.CalculateHeuristics(a.destination)
					f := g + h

					if neighborCell.F == -1 || neighborCell.F > f {
						neighborCell.F, neighborCell.G, neighborCell.H = f, g, h
						neighborCell.ParentCoord = currentCell.Coord
						heap.Push(cellsToVisit, neighborCell)
					}
				}
			}
		}
	}

	return false, -1
}
