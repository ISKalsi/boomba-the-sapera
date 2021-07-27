package models

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

func (b Battlesnake) GetBlockedCoords() []Coord {
	if b.Health == 100 {
		return b.Body
	} else {
		return b.Body[:b.Length-1]
	}
}
