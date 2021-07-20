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

func (a *Algorithm) aStarSearch() bool {
	cells := a.initGrid()
	if !a.isGoingToTail && !a.isOk(cells[a.destination]) {
		return false
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
					a.tracePath(cells)
					return true
				} else if a.isOkAndNotVisited(neighborCell) {
					g := currentCell.G + 1
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

	return false
}
