package collide_in_itself

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest3 = GameRequest{
	Game: Game{
		ID:      "a8672453-eb7f-4157-ad76-3ea87b2a68a2",
		Timeout: 500,
	},
	Turn: 59,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 4, Y: 7},
		},
		Hazards: append(
			testdata.CreateHazardCoordsFromColumns(11, 10),
			testdata.CreateHazardCoordsFromRows(11, 0, 0)...,
		),
		Snakes: []Battlesnake{
			{
				ID:     "gs_HkjRFJFBvWRW6RyShgdXR4kC",
				Name:   "Ava",
				Health: 37,
				Body: []Coord{
					{X: 8, Y: 1},
					{X: 8, Y: 2},
					{X: 8, Y: 3},
					{X: 8, Y: 4},
					{X: 8, Y: 5},
					{X: 8, Y: 6},
				},
				Head:   Coord{X: 8, Y: 1},
				Length: 6,
				Shout:  "",
			},
			{
				ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
				Name:   "Boomba",
				Health: 85,
				Body: []Coord{
					{X: 1, Y: 0},
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 1, Y: 1},
					{X: 1, Y: 1},
				},
				Head:   Coord{X: 1, Y: 0},
				Length: 5,
				Shout:  "",
			},
			{
				ID:     "gs_HkjRFadsfBvWRW6RyShgdXR4kC",
				Name:   "King Gobbla",
				Health: 98,
				Body: []Coord{
					{X: 1, Y: 2},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
					{X: 1, Y: 3},
					{X: 2, Y: 3},
					{X: 2, Y: 4},
					{X: 2, Y: 5},
					{X: 1, Y: 5},
					{X: 1, Y: 6},
					{X: 1, Y: 7},
					{X: 1, Y: 8},
				},
				Head:   Coord{X: 1, Y: 2},
				Length: 11,
				Shout:  "",
			},
			{
				ID:     "gs_ghjRFadsfBvWRW6hdShgdXR4kC",
				Name:   "SociableSnake",
				Health: 53,
				Body: []Coord{
					{X: 10, Y: 1},
					{X: 10, Y: 2},
					{X: 10, Y: 3},
					{X: 9, Y: 3},
					{X: 9, Y: 4},
					{X: 9, Y: 5},
				},
				Head:   Coord{X: 10, Y: 1},
				Length: 6,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
		Name:   "Boomba",
		Health: 85,
		Body: []Coord{
			{X: 1, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 1, Y: 1},
		},
		Head:   Coord{X: 1, Y: 0},
		Length: 5,
		Shout:  "",
	},
}
