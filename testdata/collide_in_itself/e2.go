package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var _ = GameRequest{
	Game: Game{
		ID:      "77f14cc5-9ada-413d-b42f-c67d3a2c4234",
		Timeout: 500,
	},
	Turn: 63,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 2},
			{X: 6, Y: 5},
			{X: 7, Y: 8},
			{X: 9, Y: 1},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 87,
				Body: []Coord{
					{X: 5, Y: 9},
					{X: 5, Y: 10},
					{X: 6, Y: 10},
					{X: 6, Y: 9},
					{X: 6, Y: 8},
					{X: 6, Y: 7},
					{X: 6, Y: 6},
				},
				Head:   Coord{X: 5, Y: 9},
				Length: 7,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Bunny",
				Health: 96,
				Body: []Coord{
					{X: 1, Y: 7},
					{X: 1, Y: 8},
					{X: 1, Y: 9},
					{X: 1, Y: 10},
					{X: 0, Y: 10},
					{X: 0, Y: 9},
					{X: 0, Y: 8},
					{X: 0, Y: 7},
					{X: 0, Y: 6},
					{X: 0, Y: 5},
				},
				Head:   Coord{X: 1, Y: 7},
				Length: 10,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 87,
		Body: []Coord{
			{X: 5, Y: 9},
			{X: 5, Y: 10},
			{X: 6, Y: 10},
			{X: 6, Y: 9},
			{X: 6, Y: 8},
			{X: 6, Y: 7},
			{X: 6, Y: 6},
		},
		Head:   Coord{X: 5, Y: 9},
		Length: 7,
		Shout:  "",
	},
}
