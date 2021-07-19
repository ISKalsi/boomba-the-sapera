package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

type Algorithm struct {
	board          models.Board
	start          models.Coord
	destination    models.Coord
	solvedPath     []models.Coord
	head           models.Coord
	headCollisions possibleHeadCollisions
}

type possibleHeadCollisions struct {
	coords []models.Coord
}

func (phc possibleHeadCollisions) GetBlockedCoords() []models.Coord {
	return phc.coords
}

func Init(b models.Board, s models.Battlesnake) *Algorithm {
	return &Algorithm{
		board:          b,
		start:          s.Head,
		head:           s.Head,
		destination:    s.Head.FindNearest(b.Food),
		solvedPath:     make([]models.Coord, 0),
		headCollisions: possibleHeadCollisions{coords: []models.Coord{}},
	}
}

func (a *Algorithm) clearSolvedPath() {
	a.solvedPath = a.solvedPath[:0]
}

func (a *Algorithm) SetNewStart(c models.Coord) {
	a.start = c
}

func (a *Algorithm) SetNewDestination(c models.Coord) {
	a.destination = c
}

func (a *Algorithm) findPossibleLosingHeadCollisions(ourSnake models.Battlesnake) {
	var dangerCoords []models.Coord
	for _, opponent := range a.board.Snakes {
		if opponent.Length >= ourSnake.Length {
			h := ourSnake.Head.CalculateHeuristics(opponent.Head)
			if h == 2 {
				for dir := range directionToIndex {
					c := opponent.Head.Sum(dir)
					if !c.IsOutside(a.board.Width, a.board.Height) {
						dangerCoords = append(dangerCoords, c)
					}
				}
			}
		}
	}

	a.headCollisions.coords = dangerCoords
}

func (a *Algorithm) reset(b models.Board, s models.Battlesnake) {
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.destination = s.Head.FindNearest(b.Food)
	a.clearSolvedPath()
	a.findPossibleLosingHeadCollisions(s)
}

func (a *Algorithm) getDirection(next models.Coord) string {
	dir := next.Diff(a.head)
	a.head = next
	return parseMoveDirectionToString(directionToIndex[dir])
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	a.reset(gr.Board, gr.You)

	if a.aStarSearch() {
		shortestPathNextCoord := a.solvedPath[0]
		g := a.initGrid()
		virtualSnake := g.MoveVirtualSnakeAlongPath(gr.You.Body, a.solvedPath)

		a.SetNewStart(virtualSnake[0])
		a.SetNewDestination(virtualSnake[len(virtualSnake)-1])

		if a.longestPath() {
			return a.getDirection(shortestPathNextCoord)
		}
	}

	a.SetNewStart(gr.You.Head)
	a.SetNewDestination(gr.You.Body[len(gr.You.Body)-1])

	if a.longestPath() {
		return a.getDirection(a.solvedPath[0])
	} else {
		g := a.initGrid()
		maxD := -1
		var maxDir models.Coord

		for dir := range directionToIndex {
			test := gr.You.Head.Sum(dir)
			if !test.IsOutside(a.board.Width, a.board.Height) && !g[test].IsBlocked {
				d := int(a.start.CalculateHeuristics(a.board.Food[0]))
				if d > maxD {
					maxD = d
					maxDir = dir
				}
			}
		}

		return parseMoveDirectionToString(directionToIndex[maxDir])
	}
}
