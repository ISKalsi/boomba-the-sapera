package food_related

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest3 = GameRequest{
	Game: Game{
		ID:      "41d58616-4a9b-43a2-bd08-784f5fae03ef",
		Timeout: 500,
	},
	Turn: 26,
	Board: Board{
		Height:  11,
		Width:   11,
		Hazards: testdata.CreateHazardCoordsFromRows(11, 10),
		Food: []Coord{
			{X: 5, Y: 9},
			{X: 8, Y: 4},
			{X: 9, Y: 0},
			{X: 9, Y: 8},
		},
		Snakes: []Battlesnake{
			{
				ID:     "1",
				Name:   "Badly Coded Snake",
				Health: 98,
				Body: []Coord{
					{X: 2, Y: 0},
					{X: 1, Y: 0},
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 0, Y: 2},
					{X: 1, Y: 2},
					{X: 2, Y: 2},
					{X: 2, Y: 3},
				},
				Head:   Coord{X: 2, Y: 0},
				Length: 8,
				Shout:  "",
			},
			{
				ID:     "2",
				Name:   "Boomba",
				Health: 92,
				Body: []Coord{
					{X: 9, Y: 9},
					{X: 10, Y: 9},
					{X: 10, Y: 8},
					{X: 10, Y: 7},
					{X: 10, Y: 6},
				},
				Head:   Coord{X: 9, Y: 9},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "3",
				Name:   "SociableSnake",
				Health: 82,
				Body: []Coord{
					{X: 10, Y: 2},
					{X: 9, Y: 2},
					{X: 8, Y: 2},
					{X: 8, Y: 3},
					{X: 7, Y: 3},
				},
				Head:   Coord{X: 10, Y: 2},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "4",
				Name:   "Boomslang",
				Health: 72,
				Body: []Coord{
					{X: 10, Y: 0},
					{X: 10, Y: 1},
					{X: 9, Y: 1},
					{X: 8, Y: 1},
				},
				Head:   Coord{X: 10, Y: 0},
				Length: 4,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "2",
		Name:   "Boomba",
		Health: 92,
		Body: []Coord{
			{X: 9, Y: 9},
			{X: 10, Y: 9},
			{X: 10, Y: 8},
			{X: 10, Y: 7},
			{X: 10, Y: 6},
		},
		Head:   Coord{X: 9, Y: 9},
		Length: 5,
		Shout:  "",
	},
}
