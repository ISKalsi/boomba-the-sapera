package out_of_bounds

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest3 = GameRequest{
	Game: Game{
		ID:      "63445946-14ac-4abb-a077-843cf82d9bed",
		Timeout: 500,
	},
	Turn: 38,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 10, Y: 2},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 99,
				Body: []Coord{
					{X: 8, Y: 9},
					{X: 8, Y: 9},
					{X: 7, Y: 9},
					{X: 6, Y: 9},
					{X: 5, Y: 9},
					{X: 4, Y: 9},
					{X: 3, Y: 9},
					{X: 2, Y: 9},
				},
				Head:   Coord{X: 8, Y: 9},
				Length: 8,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Ava",
				Health: 92,
				Body: []Coord{
					{X: 7, Y: 5},
					{X: 7, Y: 6},
					{X: 7, Y: 7},
					{X: 6, Y: 7},
				},
				Head:   Coord{X: 7, Y: 5},
				Length: 4,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 99,
		Body: []Coord{
			{X: 8, Y: 9},
			{X: 8, Y: 9},
			{X: 7, Y: 9},
			{X: 6, Y: 9},
			{X: 5, Y: 9},
			{X: 4, Y: 9},
			{X: 3, Y: 9},
			{X: 2, Y: 9},
		},
		Head:   Coord{X: 8, Y: 9},
		Length: 8,
		Shout:  "",
	},
}
