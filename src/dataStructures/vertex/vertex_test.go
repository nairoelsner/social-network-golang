// vertex_test.go
package vertex

import (
	"testing"
)

func TestVertexMethods(t *testing.T) {
	vertex := NewVertex("n_elsner", "Nairo Elsner", []string{"followers", "following"})

	expected := "Nairo Elsner map[followers:map[] following:map[]]"
	result := vertex.String()
	t.Log("vertex.String --->", result)
	if result != expected {
		t.Errorf("String() retornou %s, esperado %s", result, expected)
	}

	vertex2 := NewVertex("clarossa", "Clarisse Estima", []string{"followers", "following"})

	vertex.AddConnection(vertex2.GetKey(), "followers", 0)
	expected = "Nairo Elsner map[followers:map[clarossa:0] following:map[]]"
	result = vertex.String()
	t.Log("vertex.String --->", result)
	if result != expected {
		t.Errorf("String() retornou %s, esperado %s", result, expected)
	}

	//go test -v para mostrar os prints
}
