package chapter18

type WeightedGraphVertex struct {
	Value     int
	Neighbors map[*WeightedGraphVertex]int
}

func (w *WeightedGraphVertex) AddNeighbor(vertex *WeightedGraphVertex, weight int) {
	w.Neighbors[vertex] = weight
}

func NewWeightedGraphVertex(val int) *WeightedGraphVertex {
	return &WeightedGraphVertex{
		Value:     val,
		Neighbors: make(map[*WeightedGraphVertex]int),
	}
}
