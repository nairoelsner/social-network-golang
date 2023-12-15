package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nairoelsner/socialNetworkGo/src/socialNetwork/network"
)

func main() {
	network := network.NewNetwork()
	network.AddUser("n_elsner", "senha123", "Nairo Elsner")
	network.AddUser("clarossa", "senha123", "Clarisse Estima")

	network.AddFollower("n_elsner", "clarossa")
	network.AddFollower("clarossa", "n_elsner")

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 1 day
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Status": "Live"})
	})

	r.POST("/login", func(c *gin.Context) {
		var user map[string]string
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "JSON invalid format!"})
			return
		}
		username := user["username"]
		password := user["password"]

		err := network.Login(username, password)
		if err != nil {
			c.JSON(400, gin.H{"error": "Credenciais inválidas"})
		} else {
			c.JSON(200, gin.H{"username": username})
		}
	})

	r.POST("/register", func(c *gin.Context) {
		var user map[string]string
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "JSON invalid format!"})
			return
		}

		username := user["username"]
		password := user["password"]
		name := user["name"]

		fmt.Println(username, password, name)

		err := network.AddUser(username, password, name)
		if err != nil {
			fmt.Println(err)
			c.Status(400)
		} else {
			c.Status(200)
		}
	})

	r.GET("/users", func(c *gin.Context) {
		users := network.GetAllUsernames()

		data := map[string]interface{}{
			"usernames": users,
		}

		c.JSON(200, data)
	})

	r.GET("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, userExists := network.GetUser(username)
		if !userExists {
			c.JSON(400, gin.H{"error": "User doesn't exist!"})
			return
		}

		c.JSON(200, user)
	})

	r.PUT("/users/:username", func(c *gin.Context) {
		username := c.Param("username")

		var info map[string]string
		if err := c.BindJSON(&info); err != nil {
			c.JSON(400, gin.H{"error": "JSON invalid format!"})
			return
		}

		err := network.UpdateUser(username, info)
		if err != nil {
			c.JSON(400, gin.H{"error": "Couldn't update user!"})
		} else {
			c.Status(200)
		}
	})

	r.POST("/follow", func(c *gin.Context) {
		var users map[string]string
		if err := c.BindJSON(&users); err != nil {
			c.JSON(400, gin.H{"error": "Formato JSON inválido"})
			return
		}

		username1 := users["username1"]
		username2 := users["username2"]

		if err := network.AddFollower(username1, username2); err != nil {
			c.JSON(400, gin.H{"error": "Couldn't follow user"})
		} else {
			c.Status(200)
		}
	})

	r.POST("/create-post", func(c *gin.Context) {
		var data map[string]string
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "JSON invalid format!"})
		}

		username1 := data["username1"]
		username2 := data["username2"]
		text := data["text"]

		if err := network.CreatePost(username1, username2, text); err != nil {
			c.JSON(400, gin.H{"error": "Couldn't publish on mural"})
		} else {
			c.Status(200)
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
