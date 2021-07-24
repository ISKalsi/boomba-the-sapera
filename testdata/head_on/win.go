package head_on

import . "github.com/ISKalsi/boomba-the-sapera/models"

var WinCollidingSnakesRequest = GameRequest{
	Game: Game{
		ID:      "12c1ff79-5b45-4ed1-b3f8-7ae236dc02c0",
		Timeout: 500,
	},
	Turn: 0,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 5},
		},
		Snakes: []Battlesnake{
			pavandubey,
			chintukumar,
		},
	},
	You: chintukumar,
}
