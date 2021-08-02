package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var profiles []Profile

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
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}

func main() {
	profiles = append(profiles, Profile{Id: GetNextId(), Name: "CodeHouse", University: "7/31/2021"})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./profile-vue/dist", false)))
	r.GET("/api/profiles", GetProfiles)
	r.POST("/api/profiles", SignUpProfile)
	r.DELETE("/api/profiles/:id", DeleteProfile)
	r.Run(":8091")
}
