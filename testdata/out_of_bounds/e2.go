package out_of_bounds

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest2 = GameRequest{
	Game: Game{
		ID:      "4dfe30fd-f7df-461c-b1e1-c6afcf6dc239",
		Timeout: 500,
	},
	Turn: 36,
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
				Health: 93,
				Body: []Coord{
					{X: 4, Y: 10},
					{X: 5, Y: 10},
					{X: 6, Y: 10},
					{X: 7, Y: 10},
					{X: 8, Y: 10},
				},
				Head:   Coord{X: 4, Y: 10},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Boomba",
				Health: 82,
				Body: []Coord{
					{X: 4, Y: 8},
					{X: 5, Y: 8},
					{X: 6, Y: 8},
					{X: 7, Y: 8},
					{X: 7, Y: 7},
					{X: 8, Y: 7},
					{X: 9, Y: 7},
					{X: 10, Y: 7},
				},
				Head:   Coord{X: 4, Y: 8},
				Length: 8,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 93,
		Body: []Coord{
			{X: 4, Y: 10},
			{X: 5, Y: 10},
			{X: 6, Y: 10},
			{X: 7, Y: 10},
			{X: 8, Y: 10},
		},
		Head:   Coord{X: 4, Y: 10},
		Length: 5,
		Shout:  "",
	},
}
