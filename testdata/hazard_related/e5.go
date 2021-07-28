package hazard_related

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest5 = GameRequest{
	Game: Game{
		ID:      "30f24d94-9ebe-4798-a75b-dd450af6d6ed",
		Timeout: 500,
	},
	Turn: 106,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 8, Y: 1},
			{X: 10, Y: 10},
		},
		Hazards: append(
			testdata.CreateHazardCoordsFromColumns(11, 10),
			testdata.CreateHazardCoordsFromRows(11, 8, 9, 10)...,
		),
		Snakes: []Battlesnake{
			{
				ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
				Name:   "Boomba",
				Health: 53,
				Body: []Coord{
					{X: 9, Y: 9},
					{X: 10, Y: 9},
					{X: 10, Y: 10},
					{X: 9, Y: 10},
					{X: 8, Y: 10},
					{X: 7, Y: 10},
					{X: 6, Y: 10},
					{X: 6, Y: 9},
					{X: 6, Y: 8},
					{X: 5, Y: 8},
					{X: 5, Y: 7},
					{X: 6, Y: 7},
					{X: 7, Y: 7},
				},
				Head:   Coord{X: 9, Y: 9},
				Length: 13,
				Shout:  "",
			},
			{
				ID:     "gs_HkjRFJFBvWRW6RyShgdXR4kC",
				Name:   "SociableSnake",
				Health: 90,
				Body: []Coord{
					{X: 5, Y: 3},
					{X: 5, Y: 4},
					{X: 5, Y: 5},
					{X: 6, Y: 5},
					{X: 7, Y: 5},
					{X: 8, Y: 5},
					{X: 8, Y: 6},
					{X: 9, Y: 6},
					{X: 9, Y: 5},
					{X: 9, Y: 4},
					{X: 9, Y: 3},
					{X: 8, Y: 3},
				},
				Head:   Coord{X: 5, Y: 3},
				Length: 12,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
		Name:   "Boomba",
		Health: 53,
		Body: []Coord{
			{X: 9, Y: 9},
			{X: 10, Y: 9},
			{X: 10, Y: 10},
			{X: 9, Y: 10},
			{X: 8, Y: 10},
			{X: 7, Y: 10},
			{X: 6, Y: 10},
			{X: 6, Y: 9},
			{X: 6, Y: 8},
			{X: 5, Y: 8},
			{X: 5, Y: 7},
			{X: 6, Y: 7},
			{X: 7, Y: 7},
		},
		Head:   Coord{X: 9, Y: 9},
		Length: 13,
		Shout:  "",
	},
}
