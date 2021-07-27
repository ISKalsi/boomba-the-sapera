package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest6 = GameRequest{
	Game: Game{
		ID:      "fae12cd9-f24a-4b8f-b7e0-df08ddec713f",
		Timeout: 500,
	},
	Turn: 29,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 1, Y: 3},
		},
		Snakes: []Battlesnake{
			{
				ID:     "1",
				Name:   "snake30series",
				Health: 91,
				Body: []Coord{
					{X: 2, Y: 5},
					{X: 3, Y: 5},
					{X: 4, Y: 5},
					{X: 5, Y: 5},
					{X: 6, Y: 5},
					{X: 7, Y: 5},
				},
				Head:   Coord{X: 2, Y: 5},
				Length: 6,
				Shout:  "",
			},
			{
				ID:     "2",
				Name:   "drebin",
				Health: 87,
				Body: []Coord{
					{X: 2, Y: 1},
					{X: 2, Y: 2},
					{X: 2, Y: 3},
					{X: 3, Y: 3},
					{X: 3, Y: 4},
					{X: 4, Y: 4},
				},
				Head:   Coord{X: 2, Y: 1},
				Length: 6,
				Shout:  "",
			},
			{
				ID:     "3",
				Name:   "Boomba",
				Health: 73,
				Body: []Coord{
					{X: 0, Y: 3},
					{X: 0, Y: 4},
					{X: 1, Y: 4},
					{X: 1, Y: 5},
				},
				Head:   Coord{X: 0, Y: 3},
				Length: 4,
				Shout:  "",
			},
			{
				ID:     "4",
				Name:   "BobFrit Jr",
				Health: 71,
				Body: []Coord{
					{X: 3, Y: 0},
					{X: 4, Y: 0},
					{X: 4, Y: 1},
				},
				Head:   Coord{X: 3, Y: 0},
				Length: 3,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3",
		Name:   "Boomba",
		Health: 73,
		Body: []Coord{
			{X: 0, Y: 3},
			{X: 0, Y: 4},
			{X: 1, Y: 4},
			{X: 1, Y: 5},
		},
		Head:   Coord{X: 0, Y: 3},
		Length: 4,
		Shout:  "",
	},
}
