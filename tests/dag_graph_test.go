package tests

import (
	"testing"

	"github.com/enixdark/dag-block/lib/dag"
)

func TestDAG(t *testing.T) {
	d := dag.NewDAG()

	if d.Order() != 0 {
		t.Fatalf("DAG number of vertices expected to be 0 but got %d", d.Order())
	}
}

func TestDAG_AddVertex(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	if dag1.Order() != 1 {
		t.Fatalf("DAG number of vertices expected to be 1 but got %d", dag1.Order())
	}
}

func TestDAG_DeleteVertex(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}

	if dag1.Order() != 1 {
		t.Fatalf("DAG number of vertices expected to be 1 but got %d", dag1.Order())
	}

	err = dag1.DeleteVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't delete vertex from DAG: %s", err)
	}

	if dag1.Order() != 0 {
		t.Fatalf("DAG number of vertices expected to be 0 but got %d", dag1.Order())
	}

	err = dag1.DeleteVertex(vertex1)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't: %s", err)
	}
}

func TestDAG_AddEdge(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", "two")

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}
}

func TestDAG_AddEdge_FailsVertextDontExist(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex3, vertex2)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't")
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't")
	}
}

func TestDAG_AddEdge_FailsAlreadyExists(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err == nil {
		t.Fatalf("Edge already exists, AddEdge should fail but it doesn't")
	}
}

func TestDAG_DeleteEdge(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG")
	}

	size := dag1.Size()
	if size != 1 {
		t.Fatalf("Dag expected to have 1 edge but got %d", size)
	}

	err = dag1.DeleteEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't delete edge from DAG")
	}

	size = dag1.Size()
	if size != 0 {
		t.Fatalf("Dag expected to have 0 edges but got %d", size)
	}
}


func TestDAG_Order(t *testing.T) {
	dag1 := dag.NewDAG()

	expected_order := 0
	order := dag1.Order()
	if order != expected_order {
		t.Fatalf("Expected order to be %d but got %d", expected_order, order)
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	expected_order = 3
	order = dag1.Order()
	if order != expected_order {
		t.Fatalf("Expected order to be %d but got %d", expected_order, order)
	}
}

func TestDAG_Size(t *testing.T) {
	dag1 := dag.NewDAG()

	expected_size := 0
	size := dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)
	vertex4 := dag.NewVertex("4", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex4)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	expected_size = 0
	size = dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex2, vertex4)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	expected_size = 3
	size = dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}
}

func TestDAG_SinkVertices(t *testing.T) {
	dag1 := dag.NewDAG()

	sinkVertices := dag1.SinkVertices()
	if len(sinkVertices) != 0 {
		t.Fatalf("Expected to have 0 Sink vertices but got %d", len(sinkVertices))
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	sinkVertices = dag1.SinkVertices()
	if len(sinkVertices) != 2 {
		t.Fatalf("Expected to have 2 Sink vertices but got %d", len(sinkVertices))
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	sinkVertices = dag1.SinkVertices()
	if len(sinkVertices) != 1 {
		t.Fatalf("Expected to have 1 Sink vertex but got %d", len(sinkVertices))
	}
}

func TestDAG_SourceVertices(t *testing.T) {
	dag1 := dag.NewDAG()

	sourceVertices := dag1.SourceVertices()
	if len(sourceVertices) != 0 {
		t.Fatalf("Expected to have 0 Source vertices but got %d", len(sourceVertices))
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	sourceVertices = dag1.SourceVertices()
	if len(sourceVertices) != 2 {
		t.Fatalf("Expected to have 2 Source vertices but got %d", len(sourceVertices))
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	sourceVertices = dag1.SourceVertices()
	if len(sourceVertices) != 1 {
		t.Fatalf("Expected to have 1 Source vertex but got %d", len(sourceVertices))
	}

}

func TestDAG_Successors(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	successors, err := dag1.Successors(vertex1)
	if err != nil {
		t.Fatalf("Can't get %s successors: %s", vertex1, err)
	}
	if len(successors) != 1 {
		t.Fatalf("Expected to have 1 successor but got %d", len(successors))
	}
	if successors[0].ID != "2" {
		t.Fatalf("Successor vertex expected to be '2' but got %q", successors[0].ID)
	}
}

func TestDAG_Successors_VertexNotFound(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	successors, err := dag1.Successors(vertex3)
	if err == nil {
		t.Fatalf("Got %d successors for vertex %s, but expected to fail", len(successors), vertex1.ID)
	}
}

func TestDAG_Predecessors(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	predecessors, err := dag1.Predecessors(vertex2)
	if err != nil {
		t.Fatalf("Can't get %s predecessors: %s", vertex1, err)
	}
	if len(predecessors) != 1 {
		t.Fatalf("Expected to have 1 predecessor but got %d", len(predecessors))
	}
	if predecessors[0].ID != "1" {
		t.Fatalf("Predecessor vertex expected to be '1' but got %q", predecessors[0].ID)
	}
}

func TestDAG_Predecessors_VertexNotFound(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	predecessors, err := dag1.Predecessors(vertex3)
	if err == nil {
		t.Fatalf("Got %d predecessors for vertex %s, but expected to fail", len(predecessors), vertex3.ID)
	}
}