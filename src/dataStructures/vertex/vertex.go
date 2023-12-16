package vertex

import "fmt"

type Vertex struct {
	key         interface{}
	value       interface{}
	connections map[string]map[interface{}]int
}

func NewVertex(key interface{}, value interface{}, connectionTypes []string) *Vertex {
	vertex := &Vertex{
		key:         key,
		value:       value,
		connections: make(map[string]map[interface{}]int),
	}

	for _, connType := range connectionTypes {
		vertex.AddConnectionType(connType)
	}

	return vertex
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%v %v", v.value, v.connections)
}

func (v *Vertex) GetKey() interface{} {
	return v.key
}

func (v *Vertex) GetValue() interface{} {
	return v.value
}

func (v *Vertex) GetConnections() map[string]map[interface{}]int {
	return v.connections
}

func (v *Vertex) GetConnection(connectionType string) map[interface{}]int {
	return v.connections[connectionType]
}

func (v *Vertex) GetWeight(adjKey interface{}, connectionType string) int {
	return v.connections[connectionType][adjKey]
}

func (v *Vertex) AddConnection(adjKey interface{}, connectionType string, weight int) bool {
	if _, ok := v.connections[connectionType]; !ok {
		return false
	}

	v.connections[connectionType][adjKey] = weight
	return true
}

func (v *Vertex) AddConnectionType(connectionType string) {
	v.connections[connectionType] = make(map[interface{}]int)
}

func (v *Vertex) GetConnectedKeys() []interface{} {
	connectedKeys := make(map[interface{}]struct{})
	for _, connections := range v.connections {
		for key := range connections {
			connectedKeys[key] = struct{}{}
		}
	}
	keysList := make([]interface{}, 0, len(connectedKeys))
	for key := range connectedKeys {
		keysList = append(keysList, key)
	}
	return keysList
}
