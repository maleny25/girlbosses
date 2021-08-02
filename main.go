package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var users []User
var profiles []Profile //array of type profiles

// map of username: password
var userMap = make(map[string]string)
var loggedinUser int

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

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": users})
}

func UserLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user.Username)
	if pass, work := userMap[user.Username]; work {
		if pass == user.Password {
			for index := range users {
				if users[index].Username == user.Username {
					loggedinUser = index
					c.String(http.StatusOK, c.FullPath()+"/"+strconv.Itoa(index))
					return
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username", "user": user.Username})
}

func SignUpUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = GetNextId()
	users = append(users, user)
	userMap[user.Username] = user.Password
	loggedinUser = user.Id

	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(user.Id))
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
	users = append(users, User{Id: GetNextId(), Username: "CodeHouse", Password: "7/31/2021"})
	userMap["CodeHouse"] = "7/31/2021"

	profiles = append(profiles, Profile{Id: GetNextId(), Name: "Jane Doe", Age: "22", Gender: "Female", Race: "White",
		University: "College University", Major: "Buisness", Minor: "", Gradyear: "2021",
		Description: "Graduating senior looking for a consulting company."})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./profile-vue/dist", false)))
	r.GET("/api/profiles", GetProfiles)
	r.POST("/api/profiles", SignUpProfile)
	r.GET("/api/profiles/:id", GetProfile)

	users = append(users, User{Id: GetNextId(), Username: "CodeHouse", Password: "7/31/2021"})
	profiles = append(profiles, Profile{Id: 0, Name: "name", Age: "0", University: "university"})

	r.Use(static.Serve("/", static.LocalFile("./vue-project/dist", false)))
	r.GET("/api/users", GetUsers)
	r.POST("/api/users", SignUpUser)
	r.POST("/api/users/login", UserLogin)
	r.Run(":8090")

}
