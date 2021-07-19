package testdata

import . "github.com/ISKalsi/boomba-the-sapera/models"

var OutOfBoundsEdgeCaseRequest = GameRequest{
	Game: Game{
		ID:      "77f14cc5-9ada-413d-b42f-c67d3a2c4234",
		Timeout: 500,
	},
	Turn: 63,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 0, Y: 0},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 76,
				Body: []Coord{
					{X: 3, Y: 0},
					{X: 3, Y: 1},
					{X: 3, Y: 2},
					{X: 3, Y: 3},
					{X: 3, Y: 4},
					{X: 4, Y: 4},
					{X: 4, Y: 3},
					{X: 4, Y: 2},
					{X: 4, Y: 1},
				},
				Head:   Coord{X: 3, Y: 0},
				Length: 9,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Bunny",
				Health: 80,
				Body: []Coord{
					{X: 1, Y: 0},
					{X: 1, Y: 1},
					{X: 1, Y: 2},
					{X: 1, Y: 3},
					{X: 1, Y: 4},
					{X: 1, Y: 5},
					{X: 2, Y: 5},
					{X: 3, Y: 5},
					{X: 4, Y: 5},
					{X: 4, Y: 6},
				},
				Head:   Coord{X: 1, Y: 0},
				Length: 10,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 76,
		Body: []Coord{
			{X: 3, Y: 0},
			{X: 3, Y: 1},
			{X: 3, Y: 2},
			{X: 3, Y: 3},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 4, Y: 3},
			{X: 4, Y: 2},
			{X: 4, Y: 1},
		},
		Head:   Coord{X: 3, Y: 0},
		Length: 9,
		Shout:  "",
	},
}
