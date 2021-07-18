package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Algorithm struct {
	board       models.Board
	start       models.Coord
	destination models.Coord
	solvedPath  []models.Coord
	isSolving   bool
	head        models.Coord
}

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

func Init(b models.Board, s models.Battlesnake) *Algorithm {
	return &Algorithm{
		board:       b,
		start:       s.Head,
		head:        s.Head,
		destination: b.Food[0],
		solvedPath:  make([]models.Coord, 0),
		isSolving:   false,
	}
}

func (a *Algorithm) reset(b models.Board, s models.Battlesnake) {
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.destination = b.Food[0]
	a.solvedPath = a.solvedPath[:0]
	a.isSolving = false
}

func (a *Algorithm) getDirection(next models.Coord) string {
	dir := coord.Diff(&next, &a.head)
	a.head = next
	return parseMoveDirectionToString(directionToIndex[dir])
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	a.reset(gr.Board, gr.You)

	if a.aStarSearch() {
		shortestPathNextCoord := a.solvedPath[0]
		g := a.initGrid()
		virtualSnake := g.MoveVirtualSnakeAlongPath(gr.You.Body, a.solvedPath)

		a.start = virtualSnake[0]
		a.destination = virtualSnake[len(virtualSnake)-1]

		if a.longestPath() {
			return a.getDirection(shortestPathNextCoord)
		}
	}

	a.start = gr.You.Head
	a.destination = gr.You.Body[len(gr.You.Body)-1]

	if a.longestPath() {
		return a.getDirection(a.solvedPath[0])
	} else {
		g := a.initGrid()
		maxD := -1
		var maxDir models.Coord

		for dir := range directionToIndex {
			test := coord.Sum(&gr.You.Head, &dir)
			if !g[test].IsBlocked {
				d := int(coord.CalculateHeuristics(&a.start, &a.board.Food[0]))
				if d > maxD {
					maxD = d
				}
			}
		}

		return parseMoveDirectionToString(directionToIndex[maxDir])
	}
}
