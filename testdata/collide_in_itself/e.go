package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var _ = GameRequest{
	Game: Game{
		ID:      "e19c2921-a4c8-4a77-b58a-d6a7de6ce63b",
		Timeout: 500,
	},
	Turn: 10,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 0, Y: 1},
			{X: 2, Y: 6},
			{X: 4, Y: 8},
			{X: 7, Y: 7},
			{X: 8, Y: 0},
		},
		Snakes: []Battlesnake{
			{
				ID:     "1",
				Name:   "Boomba",
				Health: 90,
				Body: []Coord{
					{X: 9, Y: 1},
					{X: 9, Y: 2},
					{X: 9, Y: 3},
				},
				Head:   Coord{X: 9, Y: 1},
				Length: 3,
				Shout:  "",
			},
			{
				ID:     "2",
				Name:   "Bunny",
				Health: 92,
				Body: []Coord{
					{X: 1, Y: 9},
					{X: 2, Y: 9},
					{X: 2, Y: 10},
				},
				Head:   Coord{X: 1, Y: 9},
				Length: 3,
				Shout:  "",
			},
			{
				ID:     "3",
				Name:   "Engine Chill",
				Health: 90,
				Body: []Coord{
					{X: 9, Y: 7},
					{X: 9, Y: 6},
					{X: 9, Y: 5},
					{X: 10, Y: 5},
				},
				Head:   Coord{X: 9, Y: 7},
				Length: 4,
				Shout:  "",
			},
			{
				ID:     "4",
				Name:   "SociableSnake",
				Health: 98,
				Body: []Coord{
					{X: 5, Y: 7},
					{X: 5, Y: 6},
					{X: 5, Y: 5},
					{X: 5, Y: 4},
					{X: 6, Y: 4},
				},
				Head:   Coord{X: 5, Y: 7},
				Length: 5,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "1",
		Name:   "Boomba",
		Health: 90,
		Body: []Coord{
			{X: 9, Y: 1},
			{X: 9, Y: 2},
			{X: 9, Y: 3},
		},
		Head:   Coord{X: 9, Y: 1},
		Length: 3,
		Shout:  "",
	},
}
