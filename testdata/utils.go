package testdata

import "github.com/ISKalsi/boomba-the-sapera/models"

func CreateHazardCoordsFromRows(width int, rows ...int) []models.Coord {
	hazards := make([]models.Coord, len(rows)*width)
	for i, row := range rows {
		for j := 0; j < width; j++ {
			c := models.Coord{X: j, Y: row}
			hazards[i*width+j] = c
		}
	}
	return hazards
}

func CreateHazardCoordsFromColumns(height int, cols ...int) []models.Coord {
	hazards := make([]models.Coord, len(cols)*height)
	for i, col := range cols {
		for j := 0; j < height; j++ {
			c := models.Coord{X: col, Y: j}
			hazards[i*height+j] = c
		}
	}
	return hazards
}
