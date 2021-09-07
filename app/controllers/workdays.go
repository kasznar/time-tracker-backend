package controllers

import (
	"github.com/gin-gonic/gin"
	"kasznar.hu/hello/app/dto"
	"kasznar.hu/hello/app/services"
	"kasznar.hu/hello/app/util"
	"net/http"
	"strconv"
	"time"
)

func WorkdayRoutes(router *gin.RouterGroup) {
	router.POST("/workday", listAllWorkdays)
	router.POST("/workday/create", createWorkday)
	router.POST("/workday/update", updateWorkday)
	router.GET("/workday/user/:id", listWorkdayByUser)
	router.GET("/workday/team/:id/month/:year/:month", listWorkdaysByTeamAndMonth)
}

func listAllWorkdays(c *gin.Context) {
	var request dto.ListWorkDaysRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.UserId == nil && request.Year == nil && request.Month == nil {
		c.JSON(200, services.ListAllWorkDays())
	}

	if request.UserId != nil && request.Year == nil && request.Month == nil {
		c.JSON(200, services.ListWorkDaysByUser(*request.UserId))
	}

	if request.UserId != nil && request.Year != nil && request.Month != nil {
		month := time.Month(*request.Month)
		c.JSON(200, services.ListWorkDaysByUserAndMonth(*request.UserId, *request.Year, month))
	}
}

func listWorkdaysByTeamAndMonth(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	month, err := strconv.Atoi(c.Param("month"))
	teamId, err := util.StringToUint(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, services.ListWorkDaysByTeamAndMonth(teamId, year, time.Month(month)))
}

func createWorkday(c *gin.Context) {
	// Validate request
	var request dto.CrateWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateWorkday(request)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"workdayId": id,
	})
}

func updateWorkday(c *gin.Context) {
	// Validate request
	var request dto.UpdateWorkdayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.UpdateWorkday(request)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"workdayId": id,
	})
}

func listWorkdayByUser(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := services.ListWorkDaysByUser(id)

	c.JSON(200, response)
}
