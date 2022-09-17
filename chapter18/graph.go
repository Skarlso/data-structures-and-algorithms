package chapter18

// Vertex represents a single node of a graph.
type Vertex[T comparable] struct {
	Value     T
	Neighbors map[T]*Vertex[T]
}

func (v *Vertex[T]) AddNeighbor(vertex *Vertex[T]) {
	if _, ok := v.Neighbors[vertex.Value]; ok {
		return
	}
	v.Neighbors[vertex.Value] = vertex
	vertex.AddNeighbor(v)
}

func (v *Vertex[T]) HasNeighbor(vertex *Vertex[T]) bool {
	for _, n := range v.Neighbors {
		if n == vertex {
			return true
		}
	}
	return false
}

func DFS[T comparable](current, goal *Vertex[T], visited map[T]struct{}) *Vertex[T] {
	if current.Value == goal.Value {
		return current
	}

	visited[current.Value] = struct{}{}
	for _, n := range current.Neighbors {
		if _, ok := visited[n.Value]; ok {
			continue
		}
		if n.Value == goal.Value {
			return n
		}

		found := DFS(n, goal, visited)
		if found != nil {
			return found
		}
	}

	return nil
}

func ShortestPath[T comparable](start, goal *Vertex[T]) []*Vertex[T] {
	queue := make([]*Vertex[T], 0)
	cameFrom := make(map[*Vertex[T]]*Vertex[T])
	cameFrom[start] = nil

	// BFS
	queue = append(queue, start)
	var current *Vertex[T]
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current == goal {
			break
		}
		for _, n := range current.Neighbors {
			if _, ok := cameFrom[n]; !ok {
				queue = append(queue, n)
				cameFrom[n] = current
			}
		}
	}

	// Path traversal backwards.
	current = goal
	var path []*Vertex[T]
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	return path
}

func BFS[T comparable](start, goal *Vertex[T]) *Vertex[T] {
	queue := make([]*Vertex[T], 0)
	visited := make(map[*Vertex[T]]struct{})

	queue = append(queue, start)
	var current *Vertex[T]
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current == goal {
			return current
		}
		for _, n := range current.Neighbors {
			if _, ok := visited[current]; !ok {
				queue = append(queue, n)
				visited[n] = struct{}{}
			}
		}
	}
	return nil
}

// NewVertex creates a new vertex with a given value and initializes the neighbor slice.
func NewVertex[T comparable](value T) *Vertex[T] {
	return &Vertex[T]{
		Value:     value,
		Neighbors: make(map[T]*Vertex[T]),
	}
}
