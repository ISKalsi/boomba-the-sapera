package hazard_related

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest1 = GameRequest{
	Game: Game{
		ID:      "6f38284e-4e7f-44e0-b2d6-79ae12a140fd",
		Timeout: 500,
	},
	Turn: 32,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 2, Y: 1},
		},
		Snakes: []Battlesnake{
			{
				ID:     "gs_C44XpmH6fjYM7kqyh4xSVMSc",
				Name:   "Boomba",
				Health: 89,
				Body: []Coord{
					{X: 3, Y: 1},
					{X: 3, Y: 2},
					{X: 3, Y: 3},
					{X: 3, Y: 4},
					{X: 3, Y: 5},
				},
				Head:   Coord{X: 3, Y: 1},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "gs_Dbvjv66B7hdgTKBFpqCG73wR",
				Name:   "Bunny",
				Health: 83,
				Body: []Coord{
					{X: 1, Y: 1},
					{X: 0, Y: 1},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
					{X: 0, Y: 4},
					{X: 1, Y: 4},
				},
				Head:   Coord{X: 1, Y: 1},
				Length: 6,
				Shout:  "",
			},
			{
				ID:     "abcd",
				Name:   "Imaginary",
				Health: 100,
				Body: []Coord{
					{X: 4, Y: 5},
					{X: 4, Y: 6},
					{X: 3, Y: 6},
					{X: 2, Y: 6},
					{X: 2, Y: 5},
				},
				Head:   Coord{X: 4, Y: 5},
				Length: 5,
				Shout:  "",
			},
		},
		Hazards: testdata.CreateHazardCoordsFromRows(11, 0),
	},
	You: Battlesnake{
		ID:     "gs_C44XpmH6fjYM7kqyh4xSVMSc",
		Name:   "Boomba",
		Health: 89,
		Body: []Coord{
			{X: 3, Y: 1},
			{X: 3, Y: 2},
			{X: 3, Y: 3},
			{X: 3, Y: 4},
			{X: 3, Y: 5},
		},
		Head:   Coord{X: 3, Y: 1},
		Length: 5,
		Shout:  "",
	},
}
