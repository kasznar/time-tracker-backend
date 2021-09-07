package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"kasznar.hu/hello/app/controllers"
	"kasznar.hu/hello/app/repository"
)

func main() {
	repository.Connect()
	repository.MigrateOrm()

	router := gin.Default()

	base := router.Group("/api")

	controllers.UserRoutes(base)
	controllers.CalendarDayRoutes(base)
	controllers.WorkdayRoutes(base)
	controllers.TeamRoutes(base)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
