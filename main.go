package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nairoelsner/socialNetworkGo/src/socialNetwork/network"
)

func main() {
	network := network.NewNetwork()
	network.AddUser("clarossa", "senha123", "Clarisse Estima")
	network.AddUser("endriys", "senha123", "Gabriel Endres")
	network.AddUser("n_elsner", "senha123", "Nairo Elsner")
	network.AddUser("seven_renato", "senha123", "Paulo Renato")

	network.AddFollower("clarossa", "endriys")
	network.AddFollower("clarossa", "n_elsner")
	network.AddFollower("clarossa", "seven_renato")
	network.AddFollower("endriys", "clarossa")
	network.AddFollower("endriys", "n_elsner")
	network.AddFollower("endriys", "seven_renato")
	network.AddFollower("n_elsner", "clarossa")
	network.AddFollower("n_elsner", "endriys")
	network.AddFollower("n_elsner", "seven_renato")
	network.AddFollower("seven_renato", "clarossa")
	network.AddFollower("seven_renato", "endriys")
	network.AddFollower("seven_renato", "n_elsner")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

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

	r.POST("/search", func(c *gin.Context) {
		var data map[string]string
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Formato JSON inválido"})
			return
		}

		username := data["username"]
		searchTerm := data["searchTerm"]

		results, err := network.Search(username, searchTerm)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"results": results})
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
