package chapter18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	atlanta := NewCity("Atlanta")
	boston := NewCity("Boston")
	chicago := NewCity("Chicago")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston, 100)
	atlanta.AddRoute(denver, 160)
	boston.AddRoute(chicago, 120)
	boston.AddRoute(denver, 180)
	chicago.AddRoute(elPaso, 60)
	denver.AddRoute(chicago, 40)
	denver.AddRoute(elPaso, 140)

	path := DijkstraShortestPath(atlanta, elPaso)
	assert.Equal(t, []*City{
		elPaso,
		chicago,
		boston,
		atlanta,
	}, path)
}
