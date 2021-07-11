package algo

import (
	"container/heap"
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/grid"
	"github.com/ISKalsi/boomba-the-sapera/ds/pque"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

func (a *Algorithm) initOpenList(cells grid.Grid) *pque.PriorityQueue {
	start := cells[a.start]
	start.F, start.G, start.H = 0, 0, 0
	start.Coord = a.start
	start.Parent = a.start

	openList := &pque.PriorityQueue{}
	heap.Push(openList, start)

	return openList
}

func (a *Algorithm) initGrid() grid.Grid {
	g := grid.New(a.board.Width, a.board.Height)

	for i := 0; i < a.board.Width; i++ {
		for j := 0; j < a.board.Height; j++ {
			c := models.Coord{X: j, Y: i}
			g[c].Coord = c
			g[c].IsBlocked = false
		}
	}

	for _, snake := range a.board.Snakes {
		for _, coord := range snake.Body {
			g[coord].IsBlocked = true
		}
	}
	return g
}

func (a *Algorithm) tracePath(g grid.Grid) {
	current := a.end
	parent := g[current].Parent

	for current != parent {
		a.solvedPath.Push(current)
		current = parent
		parent = g[current].Parent
	}
}

func (a *Algorithm) nextMove() string {
	next := a.solvedPath.Pop().(models.Coord)
	dir := grid.Diff(&next, &a.head)
	a.head = next
	return parseMoveIndexToString(directions[dir])
}

func (a *Algorithm) aStarSearch() bool {
	a.isSolving = true

	cells := a.initGrid()
	visited := map[models.Coord]bool{}
	openList := a.initOpenList(cells)

	for openList.Len() != 0 {
		current := heap.Pop(openList).(*cell.Cell)
		visited[current.Coord] = true

		for dir := range directions {
			coord := grid.Sum(current, &dir)

			if !grid.IsOutOfGrid(&coord, a.board.Width, a.board.Height) {
				neighbor := cells[coord]

				if coord == a.end {
					neighbor.Parent = current.Coord
					a.tracePath(cells)
					println("Path Found!")
					a.isSolving = false
					return true
				} else if !visited[coord] && !cells.IsBlocked(&coord) {
					g := current.G + 1
					h := grid.CalculateHeuristics(&coord, &a.end)
					f := g + h

					if neighbor.F == -1 || neighbor.F > f {
						neighbor.F, neighbor.G, neighbor.H = f, g, h
						neighbor.Parent = current.Coord
						heap.Push(openList, neighbor)
					}
				}
			}
		}
	}

	a.isSolving = false
	return false
}
