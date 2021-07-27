package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/grid"
	"github.com/ISKalsi/boomba-the-sapera/models"
	"math"
	"sort"
)

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

type Algorithm struct {
	board          models.Board
	start          models.Coord
	destination    models.Coord
	isGoingToTail  bool
	solvedPath     []models.Coord
	head           models.Coord
	health         float64
	headCollisions possibleHeadCollisions
}

func Init(b models.Board, s models.Battlesnake) *Algorithm {
	return &Algorithm{
		board:          b,
		start:          s.Head,
		head:           s.Head,
		destination:    s.Head.FindNearestCoordsFrom(b.Food)[0],
		health:         float64(s.Health),
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

func (a *Algorithm) sortFoodByHeuristicFrom(head models.Coord) {
	sort.Slice(a.board.Food, func(i, j int) bool {
		h1 := a.board.Food[i].CalculateHeuristics(head)
		h2 := a.board.Food[j].CalculateHeuristics(head)
		return h1 < h2
	})
}

func (a *Algorithm) findNearestPlausibleFood() (bool, models.Coord) {
	a.sortFoodByHeuristicFrom(a.head)

	heads := make([]models.Coord, len(a.board.Snakes))
	for i, snake := range a.board.Snakes {
		heads[i] = snake.Head
	}

	for _, foodCoord := range a.board.Food {
		nearestHeads := foodCoord.FindNearestCoordsFrom(heads)
		if len(nearestHeads) == 1 && nearestHeads[0] == a.head {
			return true, foodCoord
		}
	}

	return false, a.board.Food[0]
}

func (a *Algorithm) reset(b models.Board, s models.Battlesnake) {
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.health = float64(s.Health)
	a.isGoingToTail = false
	a.clearSolvedPath()
}

func (a *Algorithm) getDirection(next models.Coord) string {
	dir := next.Diff(a.head)
	a.head = next
	return parseMoveDirectionToString(directionToIndex[dir])
}

func (a *Algorithm) initGrid() grid.Grid {
	obstacles := make([]grid.ObstacleProvider, len(a.board.Snakes))
	for i, snake := range a.board.Snakes {
		obstacles[i] = snake
	}

	maybeObstacles := make([]grid.PotentialObstacleProvider, 1)
	maybeObstacles[0] = a.headCollisions

	g := grid.WithObstacles(a.board.Width, a.board.Height, obstacles, maybeObstacles)

	for _, hazardCoord := range a.board.Hazards {
		g[hazardCoord].Weight = cell.WeightHazard
	}

	return g
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	a.reset(gr.Board, gr.You)
	a.findPossibleLosingHeadCollisions(gr.You)

	pathFound := false

	nearestFoodFound, foodCoord := a.findNearestPlausibleFood()
	a.destination = foodCoord

	if nearestFoodFound {
		pathFound, _ = a.aStarSearch()
	}

	if pathFound {
		shortestPathNextCoord := a.solvedPath[0]
		g := a.initGrid()
		virtualSnake := g.MoveVirtualSnakeAlongPath(gr.You.Body, a.solvedPath)

		ourSnakeIndex := 0
		originalSnakeBody := gr.You.Body[:]
		originalSnakeHealth := a.health

		for i := range a.board.Snakes {
			if a.board.Snakes[i].ID == gr.You.ID {
				ourSnakeIndex = i
				a.board.Snakes[i].Body = virtualSnake
				a.board.Snakes[i].Head = virtualSnake[0]
				if g[foodCoord].Weight == cell.WeightHazard {
					a.health -= 100 - cell.WeightHazard
				}
				break
			}
		}

		a.SetNewStart(virtualSnake[0])
		a.SetNewDestination(virtualSnake[len(virtualSnake)-1])
		a.isGoingToTail = true

		willEatFoodInNextTurn := false
		if len(a.solvedPath) == 1 {
			willEatFoodInNextTurn = true
		}

		if found, _ := a.aStarSearch(); found {
			if len(a.solvedPath) != 1 && !willEatFoodInNextTurn {
				return a.getDirection(shortestPathNextCoord)
			}
		}

		a.board.Snakes[ourSnakeIndex].Body = originalSnakeBody
		a.board.Snakes[ourSnakeIndex].Head = originalSnakeBody[0]
		a.health = originalSnakeHealth
	}

	a.SetNewStart(gr.You.Head)
	a.SetNewDestination(gr.You.Body[len(gr.You.Body)-1])
	a.isGoingToTail = true

	if a.longestPath() {
		return a.getDirection(a.solvedPath[0])
	} else {
		g := a.initGrid()
		minF := math.Inf(1)
		var maxDir models.Coord

		for dir := range directionToIndex {
			test := gr.You.Head.Sum(dir)

			isOwnBody := gr.You.Body[1] == test
			if isOwnBody || test.IsOutside(a.board.Width, a.board.Height) {
				continue
			}

			if minF == math.Inf(1) && !g[test].IsBlocked {
				minF = -gr.You.Head.CalculateHeuristics(foodCoord) + g[test].Weight
				maxDir = dir
				continue
			} else if !g[test].IsOk() {
				continue
			}

			H := gr.You.Head.CalculateHeuristics(foodCoord)
			G := g[test].Weight
			F := G - H
			if F <= minF {
				minF = F
				maxDir = dir
			}
		}

		return parseMoveDirectionToString(directionToIndex[maxDir])
	}
}
