package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var users []User
var profiles []Profile //array of type profiles

func GetNextId() int {
	value := nextId
	nextId++
	return value
}


func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": profiles})
}

func SignUpProfile(c *gin.Context) {
	var item Profile
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.Id = GetNextId()
	profiles = append(profiles, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(item.Id))
}

func DeleteProfile(c *gin.Context) {
	idString := c.Param("id")
	if id, err := strconv.Atoi(idString); err == nil {
		for index := range profiles {
			if profiles[index].Id == id {
				profiles = append(profiles[:index], profiles[index+1:]...)
				c.Writer.WriteHeader(http.StatusNoContent)

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

func GetProfile(c *gin.Context) {
	idString := c.Param("id")

	if id, err := strconv.Atoi(idString); err == nil {
		for index := range profiles {
			if profiles[index].Id == id {
				c.JSON(http.StatusOK, gin.H{"element": profiles[index]})
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}

func main() {
	profiles = append(profiles, Profile{Id: GetNextId(), Name: "Jane Doe", Age: "22", Gender: "Female", Race: "White",
		University: "College University", Major: "Buisness", Minor: "", Gradyear: "2021",
		Description: "Graduating senior looking for a consulting company."})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./profile-vue/dist", false)))
	r.GET("/api/profiles", GetProfiles)
	r.POST("/api/profiles", SignUpProfile)
	r.DELETE("/api/profiles/:id", DeleteProfile)

	users = append(users, User{Id: GetNextId(), Username: "CodeHouse", Password: "7/31/2021"})
	profiles = append(profiles, Profile{Id: 0, Name: "name", Age: 0, University: "university"})

	r.Use(static.Serve("/", static.LocalFile("./vue-project/dist", false)))
	r.GET("/api/users", GetUsers)
	r.POST("/api/users", SignUpUser)
	r.DELETE("/api/users/:id", DeleteUser)
	r.Run(":8090")

}
