package food_related

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest1 = GameRequest{
	Game: Game{
		ID:      "f9a93cc3-7657-42d7-90d2-a4c7136b6180",
		Timeout: 500,
	},
	Turn: 46,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 1, Y: 1},
			{X: 3, Y: 7},
			{X: 2, Y: 0},
			{X: 9, Y: 4},
			{X: 0, Y: 9},
			{X: 0, Y: 10},
			{X: 5, Y: 10},
			{X: 8, Y: 9},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 49,
				Body: []Coord{
					{X: 1, Y: 5},
					{X: 1, Y: 6},
					{X: 0, Y: 6},
					{X: 0, Y: 7},
					{X: 0, Y: 8},
					{X: 1, Y: 8},
					{X: 1, Y: 7},
					{X: 2, Y: 7},
				},
				Head:   Coord{X: 1, Y: 5},
				Length: 8,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "SociableSnake",
				Health: 93,
				Body: []Coord{
					{X: 5, Y: 1},
					{X: 5, Y: 2},
					{X: 5, Y: 3},
					{X: 5, Y: 4},
					{X: 5, Y: 5},
					{X: 5, Y: 6},
					{X: 5, Y: 7},
					{X: 5, Y: 8},
				},
				Head:   Coord{X: 5, Y: 1},
				Length: 8,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 49,
		Body: []Coord{
			{X: 1, Y: 5},
			{X: 1, Y: 6},
			{X: 0, Y: 6},
			{X: 0, Y: 7},
			{X: 0, Y: 8},
			{X: 1, Y: 8},
			{X: 1, Y: 7},
			{X: 2, Y: 7},
		},
		Head:   Coord{X: 1, Y: 5},
		Length: 8,
		Shout:  "",
	},
}
