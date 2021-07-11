package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/ds/stack"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Algorithm struct {
	board      models.Board
	start      models.Coord
	end        models.Coord
	solvedPath *stack.Stack
	isSolving  bool
	head       models.Coord
}

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

func Init(b models.Board) *Algorithm {
	s := b.Snakes[0]
	return &Algorithm{
		board:      b,
		start:      s.Head,
		head:       s.Head,
		end:        b.Food[0],
		solvedPath: stack.New(),
		isSolving:  false,
	}
}

func (a *Algorithm) reset(b models.Board) {
	s := b.Snakes[0]
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.end = b.Food[0]
	a.solvedPath.Clear()
	a.isSolving = false
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	lengthChanged := false
	for i, snake := range gr.Board.Snakes {
		if a.board.Snakes[i].Length != snake.Length {
			lengthChanged = true
			break
		}
	}

	if lengthChanged {
		a.reset(gr.Board)
		if a.aStarSearch() {
			return a.nextMove()
		} else {
			return getRandomMove("random move: ")
		}
	} else if a.solvedPath.Len() != 0 {
		return a.nextMove()
	} else if a.isSolving {
		return getRandomMove("solving...: ")
	} else {
		a.reset(gr.Board)
		if a.aStarSearch() {
			return a.nextMove()
		} else {
			return getRandomMove("random move: ")
		}
	}
}
