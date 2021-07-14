package testdata

import . "github.com/ISKalsi/boomba-the-sapera/models"

var snakes = []Battlesnake{
	{
		ID:     "9515499c-48cf-4228-8fe0-ae1433968b67",
		Name:   "pd",
		Health: 100,
		Body: []Coord{
			{X: 5, Y: 9},
			{X: 5, Y: 9},
			{X: 5, Y: 9},
		},
		Head:   Coord{X: 5, Y: 9},
		Length: 3,
		Shout:  "",
	},
	{
		ID:     "9515499c-48cf-4228-8fe0-ae1433968b67",
		Name:   "pd",
		Health: 100,
		Body: []Coord{
			{X: 5, Y: 9},
			{X: 4, Y: 9},
			{X: 4, Y: 10},
		},
		Head:   Coord{X: 5, Y: 9},
		Length: 3,
		Shout:  "",
	},
}

var StartGameRequest = GameRequest{
	Game: Game{
		ID:      "73b1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 6, Y: 10},
			{X: 5, Y: 5},
		},
		Snakes: snakes[:1],
	},
	You: snakes[0],
}

var ThreeLengthSnakeRequest = GameRequest{
	Game: Game{
		ID:      "73b1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 6, Y: 10},
			{X: 5, Y: 5},
		},
		Snakes: snakes[:1],
	},
	You: snakes[0],
}
