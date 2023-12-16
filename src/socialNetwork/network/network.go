package network

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nairoelsner/socialNetworkGo/src/dataStructures/graph"
	"github.com/nairoelsner/socialNetworkGo/src/socialNetwork/user"
)

type Network struct {
	graph.Graph
	usersQty int
}

func NewNetwork() *Network {
	network := &Network{
		Graph:    *graph.NewGraph(),
		usersQty: 0,
	}

	return network
}

func (n *Network) String() string {
	return fmt.Sprintf("Quantidade de usuários: %v", n.usersQty)
}

func (n *Network) GetAllUsernames() []string {
	keys := n.Graph.GetVerticesKeys()
	usernames := make([]string, len(keys))

	for i, key := range keys {
		usernames[i] = key.(string)
	}

	return usernames
}

func (n *Network) GetUser(username string) (map[string]interface{}, bool) {
	userVertex, userExists := n.Graph.GetVertex(username)

	if !userExists {
		return nil, false
	}

	user, ok := userVertex.GetValue().(*user.User)
	if !ok {
		return nil, false
	}

	connections := make(map[string][]string)
	for connType, connMap := range userVertex.GetConnections() {
		for username := range connMap {
			connections[connType] = append(connections[connType], username.(string))
		}
	}

	result := map[string]interface{}{
		"connections": connections,
		"info":        user.GetInfo(),
	}

	return result, true
}

func (n *Network) UserExists(username string) bool {
	_, vertexExists := n.Graph.GetVertex(username)
	return vertexExists
}

func (n *Network) AddUser(username string, password string, name string) error {
	if userExists := n.UserExists(username); userExists {
		return errors.New("Username already exists!")
	}

	user := user.NewUser(username, password, name)
	err := n.AddVertex(username, user, []string{"following", "followers"})
	if err != nil {
		return errors.New("Couldn't add user!")
	}

	return nil
}

func (n *Network) UpdateUser(username string, info map[string]string) error {
	userVertex, userExists := n.Graph.GetVertex(username)
	if !userExists {
		return errors.New("User doesn't exist!")
	}

	user := userVertex.GetValue().(*user.User)
	err := user.UpdateInfo(info)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (n *Network) Login(username string, password string) error {
	userVertex, userExists := n.Graph.GetVertex(username)
	if !userExists {
		return errors.New("User doesn't exist!")
	}

	user := userVertex.GetValue().(*user.User)
	userInfo := user.GetAuthInfo()

	if userInfo["username"] != username || userInfo["password"] != password {
		return errors.New("Wrong password!")
	}
	return nil
}

func (n *Network) AddFollower(followerUsername string, followingUsername string) error {
	if !n.UserExists(followerUsername) || !n.UserExists(followingUsername) {
		return errors.New("User or users doesn't exist!")
	}

	if followerUsername == followingUsername {
		return errors.New("User can't follow himself!")
	}

	err := n.Graph.AddBidirectionalEdge([2]interface{}{followerUsername, followingUsername}, "following", "followers", 0, 0)
	if err != nil {
		return errors.New("Couldn't follow user!")
	}

	return nil
}

func (n *Network) CreatePost(username1 string, username2 string, text string) error {
	if !n.UserExists(username1) || !n.UserExists(username2) {
		return errors.New("User doesn't exist!")
	}

	userVertex, _ := n.Graph.GetVertex(username2)
	user := userVertex.GetValue().(*user.User)
	user.CreatePost(username1, text)
	return nil
}

func (n *Network) Search(username string, searchTerm string) ([]string, error) {
	_, userExists := n.Graph.GetVertex(username)
	if !userExists {
		return nil, errors.New("User doesn't exist!")
	}

	visited := map[string]bool{}
	for _, user := range n.GetAllUsernames() {
		visited[user] = false
	}
	visited[username] = true

	foundUsers := []string{}
	queue := []string{username}
	for len(queue) > 0 {
		currentUsername := queue[0]
		queue = queue[1:]

		userVertex, _ := n.Graph.GetVertex(currentUsername)
		neighborhood := userVertex.GetConnectedKeys()
		for _, neighbor := range neighborhood {
			neighborStr, _ := neighbor.(string)
			if !visited[neighborStr] {
				visited[neighborStr] = true
				queue = append(queue, neighborStr)

				neighborVertex, _ := n.Graph.GetVertex(neighborStr)
				neighborUser := neighborVertex.GetValue().(*user.User)

				neighborSearchableInfo := neighborUser.GetSearchableInfo()
				if strings.Contains(neighborSearchableInfo, searchTerm) {
					foundUsers = append(foundUsers, neighborStr)
				}
			}
		}
	}

	for username, haveBeenVisited := range visited {
		if !haveBeenVisited {
			userVertex, _ := n.Graph.GetVertex(username)
			user := userVertex.GetValue().(*user.User)

			userSearchableInfo := user.GetSearchableInfo()
			if strings.Contains(userSearchableInfo, searchTerm) {
				foundUsers = append(foundUsers, username)
			}

		}
	}

	return foundUsers, nil
}

func (n *Network) GetUserCenteredGraph(username string) (map[string]interface{}, error) {
	data, err := n.Graph.BreadthFirstSearch(username, 3, "following")
	if err != nil {
		return nil, errors.New("User doesn't exist!")
	}

	connectionsInterface := data["connections"].(map[interface{}][]interface{})
	distancesInterface := data["distances"].(map[interface{}]int)

	connections := make(map[string][]string)
	for key, value := range connectionsInterface {
		strKey := key.(string)
		strValue := make([]string, len(value))
		for i, v := range value {
			strValue[i] = v.(string)
		}
		connections[strKey] = strValue
	}

	distances := make(map[string]int)
	for key, value := range distancesInterface {
		strKey := key.(string)
		distances[strKey] = value
	}

	nodes := make([]string, 0, len(distances))
	for key := range distances {
		nodes = append(nodes, key)
	}

	edges := []map[string]string{}
	for user, connectedUsers := range connections {
		for _, connectedUser := range connectedUsers {
			edges = append(edges, map[string]string{"source": user, "target": connectedUser})

		}
	}
	userCenteredGraph := map[string]interface{}{"nodes": nodes, "distances": distances, "edges": edges}

	return userCenteredGraph, nil
}

func (n *Network) GetGraph() (map[string]interface{}, error) {
	nodes := n.GetAllUsernames()
	connections := []map[string]string{}
	followersQty := map[string]int{}

	for _, username := range nodes {
		userVertex, _ := n.Graph.GetVertex(username)
		userFollowing := userVertex.GetConnection("following")
		followersQty[username] = len(userVertex.GetConnection("followers"))

		for followedByUser := range userFollowing {
			connections = append(connections, map[string]string{"source": username, "target": followedByUser.(string)})
		}
	}
	networkGraph := map[string]interface{}{"nodes": nodes, "connections": connections, "followersQty": followersQty}

	return networkGraph, nil
}
