package testdata

import . "github.com/ISKalsi/boomba-the-sapera/models"

var starterSnake = Battlesnake{
	ID:     "5615499c-48cf-4228-8fe0-ae1433968b67",
	Name:   "bunty",
	Health: 100,
	Body: []Coord{
		{X: 5, Y: 9},
		{X: 5, Y: 9},
		{X: 5, Y: 9},
	},
	Head:   Coord{X: 5, Y: 9},
	Length: 3,
	Shout:  "",
}

var threeLengthSnake = Battlesnake{
	ID:     "3415499c-48cf-4228-8fe0-ae1433968b67",
	Name:   "pavandubey",
	Health: 100,
	Body: []Coord{
		{X: 5, Y: 9},
		{X: 4, Y: 9},
		{X: 4, Y: 10},
	},
	Head:   Coord{X: 5, Y: 9},
	Length: 3,
	Shout:  "",
}
