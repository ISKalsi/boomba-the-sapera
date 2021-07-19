package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var _ = GameRequest{
	Game: Game{
		ID:      "b20ef6ad-8a9d-4afd-8d29-977aa7b41006",
		Timeout: 500,
	},
	Turn: 5,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 10, Y: 2},
			{X: 8, Y: 1},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 95,
				Body: []Coord{
					{X: 9, Y: 2},
					{X: 9, Y: 1},
					{X: 10, Y: 1},
				},
				Head:   Coord{X: 9, Y: 2},
				Length: 3,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 95,
		Body: []Coord{
			{X: 9, Y: 2},
			{X: 9, Y: 1},
			{X: 10, Y: 1},
		},
		Head:   Coord{X: 9, Y: 2},
		Length: 3,
		Shout:  "",
	},
}
