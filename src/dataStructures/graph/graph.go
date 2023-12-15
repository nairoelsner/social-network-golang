package graph

import (
	"errors"
	"fmt"

	"github.com/nairoelsner/socialNetworkGo/src/dataStructures/vertex"
)

type Graph struct {
	vertices    map[interface{}]vertex.Vertex
	verticesQty int
}

func NewGraph() *Graph {
	graph := &Graph{
		vertices:    make(map[interface{}]vertex.Vertex),
		verticesQty: 0,
	}

	return graph
}

func (g *Graph) String() string {
	return fmt.Sprintf("%v %v", g.vertices, g.verticesQty)
}

func (g *Graph) AddVertex(key interface{}, value interface{}, connectionTypes []string) error {
	if _, vertexExists := g.vertices[key]; vertexExists {
		return errors.New("Vertex already exist!")
	}

	vertex := vertex.NewVertex(key, value, connectionTypes)
	g.vertices[key] = *vertex
	g.verticesQty++

	return nil
}

func (g *Graph) GetVertex(key interface{}) (vertex.Vertex, bool) {
	vertex, vertexExists := g.vertices[key]
	return vertex, vertexExists
}

func (g *Graph) GetVertices() []vertex.Vertex {
	vertices := []vertex.Vertex{}
	for _, vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}

	return vertices
}

func (g *Graph) GetVerticesKeys() []interface{} {
	vertices := []interface{}{}
	for key := range g.vertices {
		vertices = append(vertices, key)
	}

	return vertices

}

func (g *Graph) AddUnidirectionalEdge(verticesKeys [2]interface{}, connType string, weight int) error {
	v1, v1Exists := g.GetVertex(verticesKeys[0])
	_, v2Exists := g.GetVertex(verticesKeys[1])

	if !v1Exists || !v2Exists {
		return errors.New("One or two vertices does not exist!")
	}

	response := v1.AddConnection(verticesKeys[1], connType, weight)
	if !response {
		return errors.New("Couldn't add connection!")
	}

	return nil
}

func (g *Graph) AddBidirectionalEdge(verticesKeys [2]interface{}, connType1 string, connType2 string, weight1 int, weight2 int) error {
	v1, v1Exists := g.GetVertex(verticesKeys[0])
	v2, v2Exists := g.GetVertex(verticesKeys[1])

	if !v1Exists || !v2Exists {
		return errors.New("One or two vertices does not exist!")
	}

	response1 := v1.AddConnection(verticesKeys[1], connType1, weight1)
	response2 := v2.AddConnection(verticesKeys[0], connType2, weight2)
	if !response1 || !response2 {
		return errors.New("Couldn't add connection!")
	}

	return nil
}
