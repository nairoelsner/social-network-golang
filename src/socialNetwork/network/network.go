package network

import (
	"errors"
	"fmt"

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
	err := n.AddVertex(username, user, []string{"follows", "followers"})
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

	err := n.Graph.AddBidirectionalEdge([2]interface{}{followerUsername, followingUsername}, "follows", "followers", 0, 0)
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
