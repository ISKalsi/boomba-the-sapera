package hazard_related

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHazardFromRow(t *testing.T) {
	expected := []models.Coord{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
	}

	actual := createHazardCoordsFromRows(4, 0, 2)

	assert.Equal(t, expected, actual)
}

func TestHazardFromColumn(t *testing.T) {
	expected := []models.Coord{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 0},
		{2, 1},
		{2, 2},
		{2, 3},
	}

	actual := createHazardCoordsFromColumns(4, 0, 2)

	assert.Equal(t, expected, actual)
}
