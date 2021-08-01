package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var users []User

func GetNextId() int {
	value := nextId
	nextId++
	return value
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": users})
}

func SignUpUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = GetNextId()
	users = append(users, user)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(user.Id))
}

func DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	if id, err := strconv.Atoi(idString); err == nil {
		for index := range users {
			if users[index].Id == id {
				users = append(users[:index], users[index+1:]...)
				c.Writer.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}

func main() {
	users = append(users, User{Id: GetNextId(), Username: "CodeHouse", Password: "7/31/2021"})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./vue-project/dist", false)))
	r.GET("/api/users", GetUsers)
	r.POST("/api/users", SignUpUser)
	r.DELETE("/api/users/:id", DeleteUser)
	r.Run(":8090")
}
