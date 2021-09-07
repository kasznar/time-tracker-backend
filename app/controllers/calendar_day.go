package controllers

import (
	"github.com/gin-gonic/gin"
	"kasznar.hu/hello/app/services"
	"net/http"
	"strconv"
	"time"
)

func CalendarDayRoutes(router *gin.RouterGroup) {
	router.GET("/calendar-days", listAllCalendarDays)
	router.GET("/calendar-days/year/:year/month/:month", listCalendarDaysByMonth)
}

func listAllCalendarDays(c *gin.Context) {
	c.JSON(200, services.ListAllCalendarDays())
}

func listCalendarDaysByMonth(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	month, err := strconv.Atoi(c.Param("month"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, services.ListDaysInMonth(year, time.Month(month)))
}
