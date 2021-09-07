package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/services"
	"kasznar.hu/hello/app/util"
)

func TeamRoutes(router *gin.RouterGroup) {
	router.GET("/teams", listAllTeams)
	router.DELETE("/teams/:id", deleteTeam)
	router.GET("/team/:id/users", listUserByTeamId)
	router.POST("/teams/create", createTeam)
	router.POST("/teams/:teamId/user/:userId", addUserToTeam)
}

func listAllTeams(c *gin.Context) {
	c.JSON(200, services.ListAllTeams())
}

func createTeam(c *gin.Context) {
	var request dto.CreateTeam
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateTeam(request)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"userId": id,
	})
}

func deleteTeam(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.DeleteTeam(uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func listUserByTeamId(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, _ := services.ListUsersInTeam(id)

	c.JSON(200, response)
}

func addUserToTeam(c *gin.Context) {
	teamId, err := util.StringToUint(c.Param("teamId"))
	userId, err := util.StringToUint(c.Param("userId"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = services.AddUserToTeam(teamId, userId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
