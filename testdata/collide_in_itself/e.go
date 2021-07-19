package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest1 = GameRequest{
	Game: Game{
		ID:      "3e42a310-b43e-4ccd-81c6-af965c94c031",
		Timeout: 500,
	},
	Turn: 14,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 2, Y: 10},
			{X: 7, Y: 1},
		},
		Snakes: []Battlesnake{
			{
				ID:     "1",
				Name:   "Boomba",
				Health: 90,
				Body: []Coord{
					{X: 4, Y: 2},
					{X: 4, Y: 3},
					{X: 3, Y: 3},
					{X: 3, Y: 2},
					{X: 3, Y: 1},
				},
				Head:   Coord{X: 4, Y: 2},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "2",
				Name:   "Bunny",
				Health: 92,
				Body: []Coord{
					{X: 5, Y: 3},
					{X: 5, Y: 4},
					{X: 5, Y: 5},
					{X: 6, Y: 5},
					{X: 7, Y: 5},
				},
				Head:   Coord{X: 5, Y: 3},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "3",
				Name:   "amazinglySnakey",
				Health: 90,
				Body: []Coord{
					{X: 4, Y: 0},
					{X: 5, Y: 0},
					{X: 6, Y: 0},
					{X: 7, Y: 0},
					{X: 8, Y: 0},
				},
				Head:   Coord{X: 4, Y: 0},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "4",
				Name:   "The young food finder",
				Health: 98,
				Body: []Coord{
					{X: 3, Y: 5},
					{X: 2, Y: 5},
					{X: 2, Y: 6},
				},
				Head:   Coord{X: 3, Y: 5},
				Length: 3,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "1",
		Name:   "Boomba",
		Health: 90,
		Body: []Coord{
			{X: 4, Y: 2},
			{X: 4, Y: 3},
			{X: 3, Y: 3},
			{X: 3, Y: 2},
			{X: 3, Y: 1},
		},
		Head:   Coord{X: 4, Y: 2},
		Length: 5,
		Shout:  "",
	},
}
