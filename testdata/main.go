package testdata

import . "github.com/ISKalsi/boomba-the-sapera/models"

var StartGameRequest = GameRequest{
	Game: Game{
		ID:      "43b1ff79-5b45-4ed1-b3f8-7ae236dc02sdf0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 9, Y: 5},
			{X: 7, Y: 10},
		},
		Snakes: []Battlesnake{
			starterSnake,
		},
	},
	You: starterSnake,
}

var ThreeLengthSnakeRequest = GameRequest{
	Game: Game{
		ID:      "83b1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 1, Y: 0},
			{X: 3, Y: 2},
			{X: 6, Y: 10},
		},
		Snakes: []Battlesnake{
			threeLengthSnake,
		},
	},
	You: threeLengthSnake,
}

var LoseCollidingSnakesRequest = GameRequest{
	Game: Game{
		ID:      "11b1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 5},
		},
		Snakes: []Battlesnake{
			pavandubey,
			chintukumar,
		},
	},
	You: pavandubey,
}

var WinCollidingSnakesRequest = GameRequest{
	Game: Game{
		ID:      "12c1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 5},
		},
		Snakes: []Battlesnake{
			pavandubey,
			chintukumar,
		},
	},
	You: chintukumar,
}

var EqualLengthCollidingSnakesRequest = GameRequest{
	Game: Game{
		ID:      "90c2ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 5},
		},
		Snakes: []Battlesnake{
			pavandubey,
			pintusharma,
		},
	},
	You: pavandubey,
}
