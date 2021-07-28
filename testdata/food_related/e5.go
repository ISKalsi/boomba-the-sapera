package food_related

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest5 = GameRequest{
	Game: Game{
		ID:      "67e27771-a3ce-4c79-989a-97c3d3826e7f",
		Timeout: 500,
	},
	Turn: 76,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 5, Y: 10},
			{X: 9, Y: 9},
			{X: 10, Y: 9},
			{X: 0, Y: 2},
		},
		Hazards: testdata.CreateHazardCoordsFromRows(11, 8, 9, 10),
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Ava",
				Health: 98,
				Body: []Coord{
					{X: 6, Y: 2},
					{X: 5, Y: 2},
					{X: 5, Y: 3},
					{X: 4, Y: 3},
					{X: 3, Y: 3},
					{X: 2, Y: 3},
					{X: 2, Y: 4},
					{X: 1, Y: 4},
					{X: 0, Y: 4},
					{X: 0, Y: 3},
					{X: 0, Y: 2},
					{X: 0, Y: 1},
				},
				Head:   Coord{X: 6, Y: 2},
				Length: 12,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Boomba",
				Health: 16,
				Body: []Coord{
					{X: 7, Y: 9},
					{X: 8, Y: 9},
					{X: 9, Y: 9},
					{X: 9, Y: 8},
					{X: 8, Y: 8},
					{X: 7, Y: 8},
					{X: 6, Y: 8},
					{X: 5, Y: 8},
					{X: 5, Y: 9},
					{X: 5, Y: 10},
				},
				Head:   Coord{X: 7, Y: 9},
				Length: 10,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
		Name:   "Boomba",
		Health: 16,
		Body: []Coord{
			{X: 7, Y: 9},
			{X: 8, Y: 9},
			{X: 9, Y: 9},
			{X: 9, Y: 8},
			{X: 8, Y: 8},
			{X: 7, Y: 8},
			{X: 6, Y: 8},
			{X: 5, Y: 8},
			{X: 5, Y: 9},
			{X: 5, Y: 10},
		},
		Head:   Coord{X: 7, Y: 9},
		Length: 10,
		Shout:  "",
	},
}
