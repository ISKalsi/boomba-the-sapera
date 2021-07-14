package algo

import (
	"container/heap"
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
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
	current := a.destination
	parent := g[current].ParentCoord

	for current != parent {
		a.solvedPath.Push(current)
		current = parent
		parent = g[current].ParentCoord
	}
}

func (a *Algorithm) aStarSearch() bool {
	a.isSolving = true

	cells := a.initGrid()
	cellsToVisit := a.initCellsToVisitList(cells)

	for cellsToVisit.Len() != 0 {
		currentCell := heap.Pop(cellsToVisit).(*cell.Cell)
		currentCell.IsVisited = true

		for dir := range directionToIndex {
			neighborCoord := coord.Sum(currentCell, &dir)

			if !coord.IsOutside(&neighborCoord, a.board.Width, a.board.Height) {
				neighborCell := cells[neighborCoord]

				if neighborCoord == a.destination {
					neighborCell.ParentCoord = currentCell.Coord
					a.tracePath(cells)
					println("Path Found!")
					a.isSolving = false
					return true
				} else if !neighborCell.IsVisited && !neighborCell.IsBlocked {
					g := currentCell.G + 1
					h := coord.CalculateHeuristics(&neighborCoord, &a.destination)
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

	a.isSolving = false
	return false
}
