package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/services"
	"kasznar.hu/hello/app/util"
	"net/http"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/users", listAllUsers)
	router.GET("/users/:id", getUserById)
	router.DELETE("/users/:id", deleteUser)
	router.POST("/users/create", createUser)
}

func listAllUsers(c *gin.Context) {
	c.JSON(200, services.ListAllUsers())
}

func getUserById(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := services.FindUserById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(200, response)
}

func createUser(c *gin.Context) {
	// Validate request
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateUser(request)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"userId": id,
	})
}

func deleteUser(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.DeleteUser(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
