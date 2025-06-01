package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"slices"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var users = []User{
	{
		ID:1,
		Username: "andreacanton@duck.com",
		Password: "ficcante",
		Firstname: "Andrea",
		Lastname: "Canton",
	},
}

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users);
}

func login(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBindJSON(&request);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	idx := slices.IndexFunc(users, func(c User) bool {
		return c.Username == request.Username && c.Password == request.Password
	})

	if idx == -1 {
		c.JSON(http.StatusUnauthorized, "Wrong username or password")
		return
	}

	c.JSON(http.StatusOK, "Success")
}

func main() {
	router := gin.Default();

	router.GET("/users", getUsers)
	router.POST("/login", login)
	
	router.Run(":3000")
}
