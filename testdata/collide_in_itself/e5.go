package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest5 = GameRequest{
	Game: Game{
		ID:      "4dfe30fd-f7df-461c-b1e1-c6afcf6dc239",
		Timeout: 500,
	},
	Turn: 12,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 9, Y: 1},
			{X: 3, Y: 10},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 96,
				Body: []Coord{
					{X: 5, Y: 1},
					{X: 5, Y: 2},
					{X: 5, Y: 3},
					{X: 5, Y: 4},
					{X: 5, Y: 5},
				},
				Head:   Coord{X: 5, Y: 1},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Boomba",
				Health: 90,
				Body: []Coord{
					{X: 8, Y: 2},
					{X: 7, Y: 2},
					{X: 7, Y: 3},
					{X: 7, Y: 4},
				},
				Head:   Coord{X: 8, Y: 2},
				Length: 4,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 96,
		Body: []Coord{
			{X: 5, Y: 1},
			{X: 5, Y: 2},
			{X: 5, Y: 3},
			{X: 5, Y: 4},
			{X: 5, Y: 5},
		},
		Head:   Coord{X: 5, Y: 1},
		Length: 5,
		Shout:  "",
	},
}
