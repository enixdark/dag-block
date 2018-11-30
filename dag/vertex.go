package dag

import (
	"fmt"

	"github.com/enixdark/dag-block/utils"
)

// Vertex type implements a vertex of a Directed Acyclic graph or DAG.
type Vertex struct {
	ID       string
	Parents  *utils.OrderedSet
	Children *utils.OrderedSet
}

// NewVertex creates a new vertex.
func NewVertex(id string, value interface{}) *Vertex {
	v := &Vertex{
		ID:       id,
		Parents:  utils.NewOrderedSet(),
		Children: utils.NewOrderedSet(),
	}

	return v
}

func (v* Vertex) Childrens() *utils.OrderedSet {
	return v.Children
}

// Degree return the number of parents and children of the vertex
func (v *Vertex) Degree() int {
	return v.Parents.Size() + v.Children.Size()
}

// InDegree return the number of parents of the vertex or the number of edges
// entering on it.
func (v *Vertex) InDegree() int {
	return v.Parents.Size()
}

// OutDegree return the number of children of the vertex or the number of edges
// leaving it.
func (v *Vertex) OutDegree() int {
	return v.Children.Size()
}

// String implements stringer interface and prints an string representation
// of this instance.
func (v *Vertex) String() string {
	result := fmt.Sprintf("ID: %s - Parents: %d - Children: %d\n", v.ID, v.Parents.Size(), v.Children.Size())

	return result
}