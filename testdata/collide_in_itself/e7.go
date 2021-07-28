package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest7 = GameRequest{
	Game: Game{
		ID:      "7a002609-452d-4883-99b3-d5cdbd9479cb",
		Timeout: 500,
	},
	Turn: 53,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 10},
			{X: 7, Y: 8},
		},
		Snakes: []Battlesnake{
			{
				ID:     "1",
				Name:   "Ava",
				Health: 87,
				Body: []Coord{
					{X: 6, Y: 5},
					{X: 5, Y: 5},
					{X: 5, Y: 4},
					{X: 5, Y: 3},
					{X: 4, Y: 3},
					{X: 4, Y: 2},
					{X: 5, Y: 2},
				},
				Head:   Coord{X: 5, Y: 5},
				Length: 7,
				Shout:  "",
			},
			{
				ID:     "3",
				Name:   "Boomba",
				Health: 95,
				Body: []Coord{
					{X: 3, Y: 6},
					{X: 4, Y: 6},
					{X: 4, Y: 7},
					{X: 4, Y: 8},
					{X: 3, Y: 8},
					{X: 3, Y: 9},
					{X: 4, Y: 9},
					{X: 5, Y: 9},
					{X: 5, Y: 8},
				},
				Head:   Coord{X: 3, Y: 6},
				Length: 9,
				Shout:  "",
			},
			{
				ID:     "4",
				Name:   "Evader",
				Health: 80,
				Body: []Coord{
					{X: 2, Y: 7},
					{X: 2, Y: 6},
					{X: 2, Y: 5},
					{X: 2, Y: 4},
					{X: 1, Y: 4},
					{X: 0, Y: 4},
					{X: 0, Y: 3},
				},
				Head:   Coord{X: 2, Y: 7},
				Length: 7,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3",
		Name:   "Boomba",
		Health: 95,
		Body: []Coord{
			{X: 3, Y: 6},
			{X: 4, Y: 6},
			{X: 4, Y: 7},
			{X: 4, Y: 8},
			{X: 3, Y: 8},
			{X: 3, Y: 9},
			{X: 4, Y: 9},
			{X: 5, Y: 9},
			{X: 5, Y: 8},
		},
		Head:   Coord{X: 3, Y: 6},
		Length: 9,
		Shout:  "",
	},
}
