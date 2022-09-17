package chapter18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	alice := NewVertex("alice")
	bob := NewVertex("bob")
	cynthia := NewVertex("cynthia")

	alice.AddNeighbor(bob)
	alice.AddNeighbor(cynthia)
	bob.AddNeighbor(cynthia)
	cynthia.AddNeighbor(bob)
}

func TestDFS(t *testing.T) {
	alice := NewVertex("alice")
	bob := NewVertex("bob")
	cynthia := NewVertex("cynthia")
	kitti := NewVertex("kitti")
	rob := NewVertex("rob")

	alice.AddNeighbor(cynthia)
	bob.AddNeighbor(cynthia)
	cynthia.AddNeighbor(bob)
	alice.AddNeighbor(kitti)
	kitti.AddNeighbor(bob)

	found := DFS(alice, kitti, map[string]struct{}{})
	assert.Equal(t, kitti, found)
	found = DFS(alice, rob, map[string]struct{}{})
	assert.Nil(t, found)
}

func TestBFS(t *testing.T) {
	alice := NewVertex("alice")
	bob := NewVertex("bob")
	cynthia := NewVertex("cynthia")
	kitti := NewVertex("kitti")
	rob := NewVertex("rob")

	alice.AddNeighbor(cynthia)
	bob.AddNeighbor(cynthia)
	cynthia.AddNeighbor(bob)
	alice.AddNeighbor(kitti)
	kitti.AddNeighbor(bob)

	found := BFS(alice, kitti)
	assert.Equal(t, kitti, found)
	found = BFS(alice, rob)
	assert.Nil(t, found)
}

func TestShortestPath(t *testing.T) {
	idris := NewVertex("idris")
	kamil := NewVertex("kamil")
	lina := NewVertex("lina")
	sasha := NewVertex("sasha")
	marco := NewVertex("marco")
	ken := NewVertex("ken")
	talia := NewVertex("talia")

	idris.AddNeighbor(kamil)
	idris.AddNeighbor(talia)
	kamil.AddNeighbor(lina)
	lina.AddNeighbor(sasha)
	sasha.AddNeighbor(marco)
	marco.AddNeighbor(ken)
	ken.AddNeighbor(talia)
	talia.AddNeighbor(idris)

	path := ShortestPath(idris, lina)
	assert.Equal(t, []*Vertex[string]{
		lina,
		kamil,
		idris,
	}, path)
}
